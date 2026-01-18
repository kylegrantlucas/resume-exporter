# resume-exporter justfile

# Default recipe - show available commands
default:
    @just --list

# Build debug binary
build:
    cargo build

# Build release binary
release:
    cargo build --release

# Run tests
test:
    cargo test

# Check for errors without building
check:
    cargo check

# Format code
fmt:
    cargo fmt

# Lint code
lint:
    cargo clippy

# Clean build artifacts
clean:
    cargo clean

# Build and install to ~/.cargo/bin
install:
    cargo install --path .

# Run with example (requires ../resume/resume.json)
example template="modern":
    cargo run --release -- ../resume/resume.json -o example-{{template}}.pdf -t {{template}}
    @echo "Generated example-{{template}}.pdf"
