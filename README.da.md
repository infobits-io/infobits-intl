# Infobits Intl

Internationaliseringspakke til flere frameworks, der leverer data om lande, sprog, valutaer og kontinenter med oversaettelser og SVG-flag.

## Pakker

| Pakke | Platform | Register | Dokumentation |
|-------|----------|----------|---------------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Funktioner

- 248 lande med ISO-koder, hovedstaeder, opkaldskoder, TLD'er
- 185 sprog med ISO 639-1-koder
- 179 valutaer med ISO 4217-koder og symboler
- 7 kontinenter
- SVG-landeflag (inline indlejret)
- Flersprogede oversaettelser (dansk, tysk, engelsk, spansk, fransk, italiensk, kinesisk)

## Dokumentation

Fuld API-dokumentation med eksempler for alle tre sprog er tilgaengelig pa [docs.infobits.io](https://docs.infobits.io) og i mappen [`docs/`](docs/):

- [Kom godt i gang](docs/getting-started.da.md)
- [Lande](docs/countries.da.md)
- [Sprog](docs/languages.da.md)
- [Valutaer](docs/currencies.da.md)
- [Kontinenter](docs/continents.da.md)
- [Oversaettelser](docs/translations.da.md)

## Projektstruktur

```
infobits_intl/
├── data/                    # Kilde (JSON + SVG-aktiver)
│   ├── core/                # Kernedatafiler
│   ├── i18n/                # Oversaettelser efter entitetstype
│   └── assets/flags/        # SVG-landeflag
├── generator/               # Go-kodegenerator
├── packages/
│   ├── dart/                # Flutter-pakke
│   ├── typescript/          # npm-pakke
│   └── go/                  # Go-modul
├── generator.yaml           # Generatorkonfiguration
└── Makefile                 # Build-kommandoer
```

## Kommandoer

```bash
# Generer alle pakker
make generate

# Generer specifik pakke
make generate-dart
make generate-typescript
make generate-go

# Valider datafiler
make validate

# Koer tests
make test              # Alle pakker
make test-dart
make test-typescript
make test-go

# Lint kode
make lint              # Generator + Go-pakke
make lint-generator
make lint-go
```

## CI/CD

GitHub Actions-workflows er konfigureret til:

- **Lint** (`lint.yml`): Koerer ved push/PR til master
  - Go-linting (generator og Go-pakke)
  - Dart-analyse
  - TypeScript-typetjek

- **Test** (`test.yml`): Koerer ved push/PR til master
  - Generatortests og datavalidering
  - Go-, Dart- og TypeScript-pakketests

- **Deploy** (`deploy.yml`): Manuel workflow dispatch
  - Publicer til npm og/eller pub.dev
  - Understoetter dry-run-tilstand
  - Valgfri versionsoverskrivning

## Udvikling

### Tilfoejelse af nye data

1. Tilfoej post til den relevante `data/core/*.json`-fil
2. Tilfoej oversaettelser til `data/i18n/{entity}/*.json`-filer
3. For lande: tilfoej SVG-flag til `data/assets/flags/`
4. Koer `make generate`

### Arbejdsgang

1. Rediger kildefiler i `data/`
2. Koer `make validate` for at kontrollere dataintegritet
3. Koer `make generate` for at regenerere alle pakker
4. Koer `make test` for at verificere genereret kode
5. Commit aendringer

## Bidrag

### Navngivning af branches

Brug beskrivende branchnavne med praefikser:

- `feature/` - Nye funktioner (f.eks. `feature/add-region-data`)
- `fix/` - Fejlrettelser (f.eks. `fix/currency-symbol-encoding`)
- `docs/` - Dokumentationsopdateringer (f.eks. `docs/api-examples`)
- `refactor/` - Koderefaktorering (f.eks. `refactor/generator-templates`)
- `chore/` - Vedligeholdelsesopgaver (f.eks. `chore/update-dependencies`)

### Conventional Commits

Dette projekt foelger [Conventional Commits](https://www.conventionalcommits.org/). Format:

```
<type>[valgfrit scope]: <beskrivelse>

[valgfri tekst]

[valgfri(e) footer(s)]
```

**Typer:**

| Type | Beskrivelse |
|------|-------------|
| `feat` | Ny funktion |
| `fix` | Fejlrettelse |
| `docs` | Kun dokumentation |
| `style` | Kodestil (formatering, semikoloner osv.) |
| `refactor` | Koderefaktorering (ingen funktion/rettelse) |
| `perf` | Ydelsesforbedring |
| `test` | Tilfoejelse eller opdatering af tests |
| `build` | Byggesystem eller afhaengigheder |
| `ci` | CI/CD-konfiguration |
| `chore` | Andre vedligeholdelsesopgaver |

**Scopes** (valgfrit): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Eksempler:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Opret en feature-branch fra `master`
2. Foretag aendringer i henhold til ovenstaaende konventioner
3. Sørg for at alle tests bestaar: `make test`
4. Sørg for at linting bestaar: `make lint`
5. Aaben en PR med en klar titel og beskrivelse
6. Link eventuelle relaterede issues

**PR-titelformat:** Brug conventional commit-format til PR-titlen.

**PR-beskrivelsesskabelon:**

```markdown
## Resumé
Kort beskrivelse af aendringer.

## Aendringer
- Aendring 1
- Aendring 2

## Test
Hvordan aendringerne blev testet.

## Relaterede issues
Fixes #123
```

## Licens

MIT
