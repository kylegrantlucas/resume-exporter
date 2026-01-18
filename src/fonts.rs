//! System font discovery and loading

use std::path::PathBuf;
use typst::text::{Font, FontBook, FontInfo};

/// Get platform-specific system font directories
pub fn get_system_font_dirs() -> Vec<PathBuf> {
    let mut dirs = Vec::new();

    #[cfg(target_os = "macos")]
    {
        dirs.push("/System/Library/Fonts".into());
        dirs.push("/Library/Fonts".into());
        if let Some(home) = dirs::home_dir() {
            dirs.push(home.join("Library/Fonts"));
        }
        // Also check for fonts in common app locations
        dirs.push("/System/Library/Fonts/Supplemental".into());
    }

    #[cfg(target_os = "linux")]
    {
        dirs.push("/usr/share/fonts".into());
        dirs.push("/usr/local/share/fonts".into());
        if let Some(home) = dirs::home_dir() {
            dirs.push(home.join(".fonts"));
            dirs.push(home.join(".local/share/fonts"));
        }
    }

    #[cfg(target_os = "windows")]
    {
        if let Some(windir) = std::env::var_os("WINDIR") {
            dirs.push(PathBuf::from(windir).join("Fonts"));
        }
        if let Some(localappdata) = std::env::var_os("LOCALAPPDATA") {
            dirs.push(PathBuf::from(localappdata).join("Microsoft\\Windows\\Fonts"));
        }
    }

    dirs
}

/// Load all fonts from system directories
pub fn load_system_fonts() -> (FontBook, Vec<Font>) {
    let mut fonts = Vec::new();
    let mut infos = Vec::new();

    for dir in get_system_font_dirs() {
        load_fonts_from_dir(&dir, &mut fonts, &mut infos);
    }

    let book = FontBook::from_infos(infos);
    (book, fonts)
}

fn load_fonts_from_dir(dir: &PathBuf, fonts: &mut Vec<Font>, infos: &mut Vec<FontInfo>) {
    let Ok(entries) = std::fs::read_dir(dir) else {
        return;
    };

    for entry in entries.flatten() {
        let path = entry.path();

        if path.is_dir() {
            // Recurse into subdirectories
            load_fonts_from_dir(&path, fonts, infos);
            continue;
        }

        let Some(ext) = path.extension().and_then(|e| e.to_str()) else {
            continue;
        };

        // Only load font files
        if !matches!(ext.to_lowercase().as_str(), "ttf" | "otf" | "ttc" | "otc") {
            continue;
        }

        let Ok(data) = std::fs::read(&path) else {
            continue;
        };

        let buffer = typst::foundations::Bytes::new(data);

        // A font file can contain multiple fonts (especially .ttc files)
        for i in 0.. {
            let Some(font) = Font::new(buffer.clone(), i) else {
                break;
            };
            infos.push(font.info().clone());
            fonts.push(font);
        }
    }
}
