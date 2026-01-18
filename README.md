# resume-exporter

A CLI tool to convert [JSONResume](https://jsonresume.org/) to PDF using [Typst](https://typst.app/).

## Features

- **Single binary** — Typst compiler embedded, no external dependencies
- **Two templates** — Modern (two-column) and Classic (single-column)
- **System fonts** — Uses fonts installed on your system

## Install

```bash
# From source
cargo install --path .

# Or build manually
cargo build --release
./target/release/resume-exporter --help
```

## Usage

```bash
# Default: modern template, output to resume.pdf
resume-exporter resume.json

# Specify output file
resume-exporter resume.json -o my-resume.pdf

# Use classic template
resume-exporter resume.json -t classic

# Full options
resume-exporter resume.json -o output.pdf -t modern
```

## Development

```bash
# Build
just build

# Release build
just release

# Install locally
just install
```

## Thanks

- [JSONResume](https://jsonresume.org/) — Standard resume schema
- [Typst](https://typst.app/) — Modern typesetting system
