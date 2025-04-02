# n.2p5.xyz

A minimalist static site generator for n.2p5.xyz built with Go, templ, and TOML.

## Overview

This project generates a personal website from TOML data files using Go's templating capabilities. The site follows a clean, monospace aesthetic with a dark theme.

## Project Structure

```shell
.
├── data/               # Content as TOML files
│   ├── home.toml       # Home page content
│   └── media.toml      # Media/writing/speaking page content
├── dst/                # Generated HTML output
├── scripts/            # Deployment scripts
├── main.go             # Main application logic
├── site.templ          # Template definitions
├── site_templ.go       # Generated templ code
└── Makefile            # Build and deployment commands
```

## How It Works

1. Content is defined in TOML files in the data directory
2. Templates are defined in site.templ using the templ templating language
3. The Go application reads the TOML data and renders HTML using the templates
4. Output HTML files are written to the dst directory
5. Files are deployed to AWS S3 with CloudFront for distribution

## Development

Prerequisites

- Go 1.24+
- templ CLI
- entr (for file watching in dev mode)
- AWS CLI (for deployment)

## Commands

```shell
# Generate Go code from templ files
make generate

# Build the static site
make build

# Serve the site locally on http://localhost:8000
make serve

# Development mode: watch for changes and rebuild automatically
make dev

# Deploy to AWS S3
make deploy
```

## Deployment

The site is deployed to AWS S3 and distributed via CloudFront. The deployment script:

1. Syncs the dst directory to an S3 bucket
2. Creates a CloudFront invalidation to refresh the CDN cache

To deploy, set the environment variable N_2P5_XYZ_CDN_DIST_ID with your CloudFront distribution ID, then run:

```shell
make deploy
```

## License

This project is in the public domain under [The Unlicense](https://unlicense.org/).