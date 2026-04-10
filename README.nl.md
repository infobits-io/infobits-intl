# Infobits Intl

Multi-framework internationalisatiepakket met land-, taal-, valuta- en continentgegevens, inclusief vertalingen en SVG-vlaggen.

## Pakketten

| Pakket | Platform | Register | Documentatie |
|--------|----------|----------|--------------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Functies

- 248 landen met ISO-codes, hoofdsteden, landnummers, TLD's
- 185 talen met ISO 639-1 codes
- 179 valuta's met ISO 4217 codes en symbolen
- 7 continenten
- SVG-landvlaggen (inline ingebed)
- Meertalige vertalingen (Deens, Duits, Engels, Spaans, Frans, Italiaans, Chinees)

## Documentatie

Volledige API-documentatie met voorbeelden voor alle drie de programmeertalen is beschikbaar op [docs.infobits.io](https://docs.infobits.io) en in de [`docs/`](docs/) map:

- [Aan de slag](docs/getting-started.nl.md)
- [Landen](docs/countries.nl.md)
- [Talen](docs/languages.nl.md)
- [Valuta's](docs/currencies.nl.md)
- [Continenten](docs/continents.nl.md)
- [Vertalingen](docs/translations.nl.md)

## Projectstructuur

```
infobits_intl/
├── data/                    # Bron van waarheid (JSON + SVG-bestanden)
│   ├── core/                # Kerngegevensbestanden
│   ├── i18n/                # Vertalingen per entiteitstype
│   └── assets/flags/        # SVG-landvlaggen
├── generator/               # Go-codegenerator
├── packages/
│   ├── dart/                # Flutter-pakket
│   ├── typescript/          # npm-pakket
│   └── go/                  # Go-module
├── generator.yaml           # Generatorconfiguratie
└── Makefile                 # Build-commando's
```

## Commando's

```bash
# Alle pakketten genereren
make generate

# Specifiek pakket genereren
make generate-dart
make generate-typescript
make generate-go

# Gegevensbestanden valideren
make validate

# Tests uitvoeren
make test              # Alle pakketten
make test-dart
make test-typescript
make test-go

# Code controleren
make lint              # Generator + Go-pakket
make lint-generator
make lint-go
```

## CI/CD

GitHub Actions workflows zijn geconfigureerd voor:

- **Lint** (`lint.yml`): Wordt uitgevoerd bij push/PR naar master
  - Go-linting (generator en Go-pakket)
  - Dart-analyse
  - TypeScript-typecontrole

- **Test** (`test.yml`): Wordt uitgevoerd bij push/PR naar master
  - Generatortests en gegevensvalidatie
  - Tests voor Go-, Dart- en TypeScript-pakketten

- **Deploy** (`deploy.yml`): Handmatige activering
  - Publicatie naar npm en/of pub.dev
  - Ondersteunt dry-run modus
  - Optionele versie-override

## Ontwikkeling

### Nieuwe gegevens toevoegen

1. Voeg een item toe aan het juiste `data/core/*.json` bestand
2. Voeg vertalingen toe aan `data/i18n/{entity}/*.json` bestanden
3. Voor landen: voeg een SVG-vlag toe aan `data/assets/flags/`
4. Voer `make generate` uit

### Werkwijze

1. Bewerk bronbestanden in `data/`
2. Voer `make validate` uit om de gegevensintegriteit te controleren
3. Voer `make generate` uit om alle pakketten opnieuw te genereren
4. Voer `make test` uit om de gegenereerde code te verifieren
5. Commit de wijzigingen

## Bijdragen

### Branchnaamgeving

Gebruik beschrijvende branchnamen met voorvoegsels:

- `feature/` - Nieuwe functies (bijv. `feature/add-region-data`)
- `fix/` - Bugfixes (bijv. `fix/currency-symbol-encoding`)
- `docs/` - Documentatie-updates (bijv. `docs/api-examples`)
- `refactor/` - Code-refactoring (bijv. `refactor/generator-templates`)
- `chore/` - Onderhoudstaken (bijv. `chore/update-dependencies`)

### Conventional Commits

Dit project volgt de [Conventional Commits](https://www.conventionalcommits.org/) conventie. Formaat:

```
<type>[optionele scope]: <beschrijving>

[optionele body]

[optionele voettekst]
```

**Types:**

| Type | Beschrijving |
|------|--------------|
| `feat` | Nieuwe functie |
| `fix` | Bugfix |
| `docs` | Alleen documentatie |
| `style` | Codestijl (opmaak, puntkomma's, enz.) |
| `refactor` | Code-refactoring (geen functie/fix) |
| `perf` | Prestatieverbetering |
| `test` | Toevoegen of bijwerken van tests |
| `build` | Buildsysteem of afhankelijkheden |
| `ci` | CI/CD-configuratie |
| `chore` | Overige onderhoudstaken |

**Scopes** (optioneel): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Voorbeelden:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Maak een branch aan vanaf `master`
2. Breng wijzigingen aan volgens de bovenstaande conventies
3. Zorg ervoor dat alle tests slagen: `make test`
4. Zorg ervoor dat de linting slaagt: `make lint`
5. Open een PR met een duidelijke titel en beschrijving
6. Koppel gerelateerde issues

**Formaat van de PR-titel:** Gebruik het conventional commit formaat voor de PR-titel.

**PR-beschrijvingssjabloon:**

```markdown
## Samenvatting
Korte beschrijving van de wijzigingen.

## Wijzigingen
- Wijziging 1
- Wijziging 2

## Testen
Hoe de wijzigingen zijn getest.

## Gerelateerde issues
Fixes #123
```

## Licentie

MIT
