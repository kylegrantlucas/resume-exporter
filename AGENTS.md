# resume-exporter

Rust CLI: JSONResume → PDF. Renders a Tera template (`src/templates/{modern,classic}.typ.tera`) to Typst source and compiles it with the embedded Typst 0.13 crates (no external `typst` binary needed at runtime).

## Layout
- `src/main.rs` CLI (clap), `src/resume.rs` JSONResume types, `src/render.rs` Tera → Typst world → PDF, `src/fonts.rs` system font discovery.
- `.github/workflows/ci.yml`: fmt/clippy/test on push to master + PRs. `release.yml`: on `v*` tags, builds linux-x86_64 / darwin-x86_64 / darwin-arm64 tarballs (the resume repo's CI downloads `resume-exporter-linux-x86_64.tar.gz` from the latest release).

## Commands (verified in `nix develop`)
- `cargo fmt --all --check`, `cargo clippy --all-targets -- -D warnings`, `cargo test` (no tests yet), `just build|release|install`.
- `just example modern` renders `../resume/resume.json`.
- `nix build` produces the binary (`packages.default`).

## Gotchas
- Fonts are only loaded from OS font dirs (`src/fonts.rs`); nothing is bundled. Templates want Avenir Next / Helvetica Neue (present on macOS; Linux CI downloads them into `~/.local/share/fonts`). `typst fonts` in the devShell lists what's visible.
- `.gitignore` ignores `*.json`, so sample resumes don't get committed.
- Default branch is `master`.
