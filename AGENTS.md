# Infobits Intl - Agent Instructions

## Project Overview

Multi-framework internationalization package providing country, language, currency, and continent data. Uses a Go-based code generator to produce Dart, TypeScript, and Go packages from JSON source data.

## Repository Structure

```
infobits_intl/
├── data/                    # Source of truth (JSON + SVG assets)
│   ├── core/                # Core data files
│   │   ├── countries.json   # 248 countries
│   │   ├── languages.json   # 185 languages
│   │   ├── currencies.json  # 179 currencies
│   │   └── continents.json  # 7 continents
│   ├── i18n/                # Translations by entity type
│   │   ├── countries/
│   │   ├── languages/
│   │   ├── currencies/
│   │   ├── continents/
│   │   └── capitals/
│   └── assets/flags/        # SVG country flags (256 files)
├── generator/               # Go code generator
│   ├── cmd/                 # CLI commands (extract, generate, validate)
│   └── internal/            # Core logic (config, data, generator, template)
├── packages/
│   ├── dart/                # Flutter package (generated)
│   ├── typescript/          # npm package (generated)
│   └── go/                  # Go module (generated)
├── generator.yaml           # Generator configuration
└── Makefile                 # Build commands
```

## Key Commands

```bash
make generate              # Generate all packages (Dart, TypeScript, Go)
make generate-dart         # Generate Dart package only
make generate-typescript   # Generate TypeScript package only
make generate-go           # Generate Go package only
make extract               # Extract data from existing Dart files to JSON
```

## Workflow

1. **Data changes**: Edit files in `data/core/` or `data/i18n/`
2. **Regenerate**: Run `make generate`
3. **Generated files**: Output goes to `packages/{dart,typescript,go}/`

## Generated File Conventions

- **Dart**: Files use `.g.dart` suffix (e.g., `countries.g.dart`)
- **TypeScript**: Standard `.ts` extension
- **Go**: Standard `.go` extension

## Data Schema

### countries.json
```json
{
  "id": "unitedStates",        // Enum name
  "alpha2Code": "US",          // ISO 3166-1 alpha-2
  "alpha3Code": "USA",         // ISO 3166-1 alpha-3
  "numericCode": 840,          // ISO 3166-1 numeric
  "nativeName": "United States",
  "capital": "Washington, D.C.",
  "mainLanguage": "english",   // Reference to language id
  "languages": ["english"],
  "tld": ".us",
  "callingCode": 1,
  "continent": "northAmerica", // Reference to continent id
  "currency": "usd"            // Reference to currency id
}
```

### languages.json
```json
{
  "id": "english",
  "code": "en",                // ISO 639-1
  "nativeName": "English",
  "dialects": [],
  "defaultFlagCode": "gb"      // Optional, for flag display
}
```

### currencies.json
```json
{
  "id": "usd",
  "code": "USD",               // ISO 4217
  "nativeName": "US Dollar",
  "nativeNamePlural": "US Dollars",
  "symbol": "$"
}
```

### continents.json
```json
{
  "id": "northAmerica",
  "code": "NA",
  "name": "North America"
}
```

## Translation Files

Simple key-value JSON format:
```json
{
  "US": "United States",
  "GB": "United Kingdom"
}
```

Currencies use nested format:
```json
{
  "USD": {
    "name": "US Dollar",
    "pluralName": "US Dollars"
  }
}
```

## SVG Flags

- Source: `data/assets/flags/{alpha2}.svg` (lowercase)
- Minified and inlined into generated code
- Accessed via `countryFlags` map (Dart), `flags` object (TS), `Flags` map (Go)

## Adding New Data

1. Add entry to appropriate `data/core/*.json` file
2. Add translations to `data/i18n/{entity}/*.json` files
3. For countries: add SVG flag to `data/assets/flags/`
4. Run `make generate`

## Generator Architecture

The Go generator uses:
- `text/template` for code generation
- Embedded template strings in generator files
- Helper functions: `pascalCase`, `camelCase`, `escapeString`, etc.
- SVG minification (removes whitespace and comments)

## CI/CD Workflows

Located in `.github/workflows/`:

### lint.yml
Runs on push/PR to master. Jobs:
- `lint-generator`: golangci-lint on generator/
- `lint-go`: golangci-lint on packages/go/
- `lint-dart`: flutter analyze on packages/dart/
- `lint-typescript`: tsc --noEmit on packages/typescript/

### test.yml
Runs on push/PR to master. Jobs:
- `test-generator`: go test + data validation
- `test-go`: go test on packages/go/
- `test-dart`: flutter test on packages/dart/
- `test-typescript`: vitest + build on packages/typescript/

### deploy.yml
Manual workflow dispatch with inputs:
- `package`: typescript, dart, or both
- `version`: optional version override
- `dry_run`: test without publishing

**Required secrets:**
- `NPM_TOKEN`: npm access token for @infobits/intl
- `PUB_CREDENTIALS`: JSON from ~/.config/dart/pub-credentials.json

## Contributing Guidelines

### Branch Naming Convention

Use prefixes to categorize branches:

| Prefix | Purpose | Example |
|--------|---------|---------|
| `feature/` | New features | `feature/add-region-data` |
| `fix/` | Bug fixes | `fix/currency-symbol-encoding` |
| `docs/` | Documentation | `docs/api-examples` |
| `refactor/` | Code refactoring | `refactor/generator-templates` |
| `chore/` | Maintenance | `chore/update-dependencies` |

### Conventional Commits

All commits must follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>[optional scope]: <description>
```

**Types:** `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`

**Scopes:** `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Examples:**
```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Request Workflow

1. Create branch from `master` using naming convention above
2. Make changes and commit using conventional commits
3. Run `make test` and `make lint` before pushing
4. Open PR with conventional commit format title
5. Include summary, changes list, testing notes, and related issues in description

### Code Quality Checks

Before submitting:
```bash
make lint      # Run all linters
make test      # Run all tests
make validate  # Validate data files
```
