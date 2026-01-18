use std::path::PathBuf;

use anyhow::{Context, Result};
use clap::Parser;

mod fonts;
mod render;
mod resume;

use render::{Renderer, Template};
use resume::Resume;

#[derive(Parser)]
#[command(name = "resume-exporter")]
#[command(about = "Convert JSONResume to PDF via Typst")]
#[command(version)]
struct Cli {
    /// Input JSON file (JSONResume v1.0.0 format)
    input: PathBuf,

    /// Output PDF file [default: <input>.pdf]
    #[arg(short, long)]
    output: Option<PathBuf>,

    /// Template to use: classic or modern
    #[arg(short, long, default_value = "modern")]
    template: Template,
}

fn main() -> Result<()> {
    let cli = Cli::parse();

    // Default output: input stem + .pdf
    let output = cli.output.unwrap_or_else(|| cli.input.with_extension("pdf"));

    // Load and parse resume
    let json = std::fs::read_to_string(&cli.input)
        .with_context(|| format!("Failed to read input file: {}", cli.input.display()))?;

    let resume: Resume = serde_json::from_str(&json)
        .with_context(|| format!("Failed to parse JSON from: {}", cli.input.display()))?;

    eprintln!("Loading fonts...");
    let renderer = Renderer::new().context("Failed to initialize renderer")?;

    eprintln!(
        "Rendering {} template for {}...",
        cli.template, resume.basics.name
    );
    let pdf_bytes = renderer
        .render(&resume, cli.template)
        .context("Failed to render resume")?;

    // Write output
    std::fs::write(&output, pdf_bytes)
        .with_context(|| format!("Failed to write output file: {}", output.display()))?;

    eprintln!("Wrote {}", output.display());
    Ok(())
}
