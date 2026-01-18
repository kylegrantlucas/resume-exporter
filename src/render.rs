//! Template rendering and Typst compilation

use std::collections::HashMap;

use anyhow::{Context, Result};
use chrono::NaiveDate;
use tera::{Tera, Value};
use typst::diag::{FileError, FileResult};
use typst::foundations::{Bytes, Datetime};
use typst::syntax::{FileId, Source};
use typst::text::{Font, FontBook};
use typst::utils::LazyHash;
use typst::{Library, World};

use crate::fonts::load_system_fonts;
use crate::resume::Resume;

/// Template names
#[derive(Debug, Clone, Copy, Default)]
pub enum Template {
    Classic,
    #[default]
    Modern,
}

impl std::str::FromStr for Template {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "classic" => Ok(Template::Classic),
            "modern" => Ok(Template::Modern),
            _ => Err(format!("Unknown template: {s}. Use 'classic' or 'modern'.")),
        }
    }
}

impl std::fmt::Display for Template {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Template::Classic => write!(f, "classic"),
            Template::Modern => write!(f, "modern"),
        }
    }
}

/// The main renderer that converts Resume → PDF
pub struct Renderer {
    tera: Tera,
    fonts: Vec<Font>,
    book: LazyHash<FontBook>,
}

impl Renderer {
    pub fn new() -> Result<Self> {
        let mut tera = Tera::default();

        // Load embedded templates
        tera.add_raw_template("classic", include_str!("templates/classic.typ.tera"))
            .context("Failed to load classic template")?;
        tera.add_raw_template("modern", include_str!("templates/modern.typ.tera"))
            .context("Failed to load modern template")?;

        // Register custom filters
        tera.register_filter("date", format_date_filter);
        tera.register_filter("escape_typst", escape_typst_filter);

        // Load system fonts
        let (book, fonts) = load_system_fonts();
        let book = LazyHash::new(book);

        Ok(Self { tera, fonts, book })
    }

    /// Render a resume to PDF bytes
    pub fn render(&self, resume: &Resume, template: Template) -> Result<Vec<u8>> {
        // Build Tera context
        let mut ctx = tera::Context::new();
        ctx.insert("basics", &resume.basics);
        ctx.insert("work", &resume.work);
        ctx.insert("education", &resume.education);
        ctx.insert("skills", &resume.skills);
        ctx.insert("projects", &resume.projects);
        ctx.insert("volunteer", &resume.volunteer);
        ctx.insert("awards", &resume.awards);
        ctx.insert("certificates", &resume.certificates);
        ctx.insert("publications", &resume.publications);
        ctx.insert("languages", &resume.languages);
        ctx.insert("interests", &resume.interests);
        ctx.insert("references", &resume.references);
        ctx.insert("grouped_work", &resume.grouped_work());

        // Render template to Typst source
        let template_name = template.to_string();
        let typst_source = self
            .tera
            .render(&template_name, &ctx)
            .with_context(|| format!("Failed to render {template_name} template"))?;

        // Compile Typst to PDF
        self.compile_to_pdf(&typst_source)
    }

    fn compile_to_pdf(&self, source: &str) -> Result<Vec<u8>> {
        let world = ResumeWorld::new(source.to_string(), &self.book, &self.fonts);

        // Compile to document
        let result = typst::compile(&world);

        let document = result.output.map_err(|errors| {
            let messages: Vec<_> = errors.iter().map(|e| e.message.to_string()).collect();
            anyhow::anyhow!("Typst compilation failed:\n{}", messages.join("\n"))
        })?;

        // Export to PDF
        let pdf_bytes = typst_pdf::pdf(&document, &typst_pdf::PdfOptions::default())
            .map_err(|errors| {
                let messages: Vec<_> = errors.iter().map(|e| format!("{:?}", e)).collect();
                anyhow::anyhow!("PDF export failed:\n{}", messages.join("\n"))
            })?;

        Ok(pdf_bytes)
    }
}

/// Custom Tera filter: format ISO date to "Jan 2022"
fn format_date_filter(value: &Value, _args: &HashMap<String, Value>) -> tera::Result<Value> {
    let date_str = value
        .as_str()
        .ok_or_else(|| tera::Error::msg("date filter expects a string"))?;

    if date_str.is_empty() {
        return Ok(Value::String(String::new()));
    }

    // Try parsing as YYYY-MM-DD
    if let Ok(date) = NaiveDate::parse_from_str(date_str, "%Y-%m-%d") {
        return Ok(Value::String(date.format("%b %Y").to_string()));
    }

    // Try parsing as YYYY-MM
    if date_str.len() >= 7 {
        if let Ok(date) = NaiveDate::parse_from_str(&format!("{}-01", &date_str[..7]), "%Y-%m-%d") {
            return Ok(Value::String(date.format("%b %Y").to_string()));
        }
    }

    // Try parsing as just YYYY
    if date_str.len() >= 4 {
        if let Ok(year) = date_str[..4].parse::<i32>() {
            return Ok(Value::String(year.to_string()));
        }
    }

    // Return as-is if we can't parse
    Ok(Value::String(date_str.to_string()))
}

/// Custom Tera filter: escape Typst special characters
fn escape_typst_filter(value: &Value, _args: &HashMap<String, Value>) -> tera::Result<Value> {
    let s = value
        .as_str()
        .ok_or_else(|| tera::Error::msg("escape_typst filter expects a string"))?;

    // Escape @ which is special in Typst (citations)
    let escaped = s.replace('@', "\\@");

    Ok(Value::String(escaped))
}

/// Minimal Typst World implementation for compilation
struct ResumeWorld<'a> {
    source: Source,
    library: LazyHash<Library>,
    book: &'a LazyHash<FontBook>,
    fonts: &'a [Font],
}

impl<'a> ResumeWorld<'a> {
    fn new(source_text: String, book: &'a LazyHash<FontBook>, fonts: &'a [Font]) -> Self {
        let source = Source::detached(source_text);
        let library = LazyHash::new(Library::default());

        Self {
            source,
            library,
            book,
            fonts,
        }
    }
}

impl World for ResumeWorld<'_> {
    fn library(&self) -> &LazyHash<Library> {
        &self.library
    }

    fn book(&self) -> &LazyHash<FontBook> {
        self.book
    }

    fn main(&self) -> FileId {
        self.source.id()
    }

    fn source(&self, id: FileId) -> FileResult<Source> {
        if id == self.source.id() {
            Ok(self.source.clone())
        } else {
            Err(FileError::NotFound(id.vpath().as_rootless_path().into()))
        }
    }

    fn file(&self, id: FileId) -> FileResult<Bytes> {
        Err(FileError::NotFound(id.vpath().as_rootless_path().into()))
    }

    fn font(&self, index: usize) -> Option<Font> {
        self.fonts.get(index).cloned()
    }

    fn today(&self, _offset: Option<i64>) -> Option<Datetime> {
        let now = chrono::Local::now();
        Datetime::from_ymd(
            now.format("%Y").to_string().parse().ok()?,
            now.format("%m").to_string().parse().ok()?,
            now.format("%d").to_string().parse().ok()?,
        )
    }
}
