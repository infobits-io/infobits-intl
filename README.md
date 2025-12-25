# Infobits Intl

Multi-framework internationalization package providing country, language, currency, and continent data with translations and SVG flags.

## Packages

| Package | Platform | Registry | Docs |
|---------|----------|----------|------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Features

- 248 countries with ISO codes, capitals, calling codes, TLDs
- 185 languages with ISO 639-1 codes
- 179 currencies with ISO 4217 codes and symbols
- 7 continents
- SVG country flags (inline embedded)
- Multi-language translations

## Project Structure

```
infobits_intl/
├── data/                    # Source of truth (JSON + SVG assets)
│   ├── core/                # Core data files
│   ├── i18n/                # Translations by entity type
│   └── assets/flags/        # SVG country flags
├── generator/               # Go code generator
├── packages/
│   ├── dart/                # Flutter package
│   ├── typescript/          # npm package
│   └── go/                  # Go module
├── generator.yaml           # Generator configuration
└── Makefile                 # Build commands
```

## Commands

```bash
# Generate all packages
make generate

# Generate specific package
make generate-dart
make generate-typescript
make generate-go

# Validate data files
make validate

# Run tests
make test              # All packages
make test-dart
make test-typescript
make test-go

# Lint code
make lint              # Generator + Go package
make lint-generator
make lint-go
```

## CI/CD

GitHub Actions workflows are configured for:

- **Lint** (`lint.yml`): Runs on push/PR to master
  - Go linting (generator and Go package)
  - Dart analysis
  - TypeScript type checking

- **Test** (`test.yml`): Runs on push/PR to master
  - Generator tests and data validation
  - Go, Dart, and TypeScript package tests

- **Deploy** (`deploy.yml`): Manual workflow dispatch
  - Publish to npm and/or pub.dev
  - Supports dry-run mode
  - Optional version override

## Development

### Adding New Data

1. Add entry to appropriate `data/core/*.json` file
2. Add translations to `data/i18n/{entity}/*.json` files
3. For countries: add SVG flag to `data/assets/flags/`
4. Run `make generate`

### Workflow

1. Edit source files in `data/`
2. Run `make validate` to check data integrity
3. Run `make generate` to regenerate all packages
4. Run `make test` to verify generated code
5. Commit changes

## Contributing

### Branch Naming

Use descriptive branch names with prefixes:

- `feature/` - New features (e.g., `feature/add-region-data`)
- `fix/` - Bug fixes (e.g., `fix/currency-symbol-encoding`)
- `docs/` - Documentation updates (e.g., `docs/api-examples`)
- `refactor/` - Code refactoring (e.g., `refactor/generator-templates`)
- `chore/` - Maintenance tasks (e.g., `chore/update-dependencies`)

### Conventional Commits

This project follows [Conventional Commits](https://www.conventionalcommits.org/). Format:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:**

| Type | Description |
|------|-------------|
| `feat` | New feature |
| `fix` | Bug fix |
| `docs` | Documentation only |
| `style` | Code style (formatting, semicolons, etc.) |
| `refactor` | Code refactoring (no feature/fix) |
| `perf` | Performance improvement |
| `test` | Adding or updating tests |
| `build` | Build system or dependencies |
| `ci` | CI/CD configuration |
| `chore` | Other maintenance tasks |

**Scopes** (optional): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Examples:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Create a feature branch from `master`
2. Make changes following the conventions above
3. Ensure all tests pass: `make test`
4. Ensure linting passes: `make lint`
5. Open a PR with a clear title and description
6. Link any related issues

**PR Title Format:** Use conventional commit format for the PR title.

**PR Description Template:**

```markdown
## Summary
Brief description of changes.

## Changes
- Change 1
- Change 2

## Testing
How the changes were tested.

## Related Issues
Fixes #123
```

## License

MIT
