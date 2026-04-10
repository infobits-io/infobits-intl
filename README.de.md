# Infobits Intl

Multi-Framework-Internationalisierungspaket mit Laender-, Sprach-, Waehrungs- und Kontinentdaten einschliesslich Uebersetzungen und SVG-Flaggen.

## Pakete

| Paket | Plattform | Registry | Dokumentation |
|-------|-----------|----------|---------------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Funktionen

- 248 Laender mit ISO-Codes, Hauptstaedten, Vorwahlen, TLDs
- 185 Sprachen mit ISO-639-1-Codes
- 179 Waehrungen mit ISO-4217-Codes und Symbolen
- 7 Kontinente
- SVG-Laenderflaggen (inline eingebettet)
- Mehrsprachige Uebersetzungen (Daenisch, Deutsch, Englisch, Spanisch, Franzoesisch, Italienisch, Chinesisch)

## Dokumentation

Die vollstaendige API-Dokumentation mit Beispielen fuer alle drei Sprachen ist auf [docs.infobits.io](https://docs.infobits.io) und im Verzeichnis [`docs/`](docs/) verfuegbar:

- [Erste Schritte](docs/getting-started.de.md)
- [Laender](docs/countries.de.md)
- [Sprachen](docs/languages.de.md)
- [Waehrungen](docs/currencies.de.md)
- [Kontinente](docs/continents.de.md)
- [Uebersetzungen](docs/translations.de.md)

## Projektstruktur

```
infobits_intl/
├── data/                    # Datenquelle (JSON + SVG-Assets)
│   ├── core/                # Kerndatendateien
│   ├── i18n/                # Uebersetzungen nach Entitaetstyp
│   └── assets/flags/        # SVG-Laenderflaggen
├── generator/               # Go-Codegenerator
├── packages/
│   ├── dart/                # Flutter-Paket
│   ├── typescript/          # npm-Paket
│   └── go/                  # Go-Modul
├── generator.yaml           # Generatorkonfiguration
└── Makefile                 # Build-Befehle
```

## Befehle

```bash
# Alle Pakete generieren
make generate

# Bestimmtes Paket generieren
make generate-dart
make generate-typescript
make generate-go

# Datendateien validieren
make validate

# Tests ausfuehren
make test              # Alle Pakete
make test-dart
make test-typescript
make test-go

# Code linten
make lint              # Generator + Go-Paket
make lint-generator
make lint-go
```

## CI/CD

GitHub-Actions-Workflows sind konfiguriert fuer:

- **Lint** (`lint.yml`): Wird bei Push/PR auf master ausgefuehrt
  - Go-Linting (Generator und Go-Paket)
  - Dart-Analyse
  - TypeScript-Typueberpruefung

- **Test** (`test.yml`): Wird bei Push/PR auf master ausgefuehrt
  - Generatortests und Datenvalidierung
  - Go-, Dart- und TypeScript-Pakettests

- **Deploy** (`deploy.yml`): Manueller Workflow-Dispatch
  - Veroeffentlichung auf npm und/oder pub.dev
  - Unterstuetzt Dry-Run-Modus
  - Optionale Versionsaenderung

## Entwicklung

### Neue Daten hinzufuegen

1. Eintrag zur entsprechenden `data/core/*.json`-Datei hinzufuegen
2. Uebersetzungen zu `data/i18n/{entity}/*.json`-Dateien hinzufuegen
3. Fuer Laender: SVG-Flagge zu `data/assets/flags/` hinzufuegen
4. `make generate` ausfuehren

### Arbeitsablauf

1. Quelldateien in `data/` bearbeiten
2. `make validate` ausfuehren, um die Datenintegritaet zu pruefen
3. `make generate` ausfuehren, um alle Pakete neu zu generieren
4. `make test` ausfuehren, um den generierten Code zu verifizieren
5. Aenderungen committen

## Mitwirken

### Branch-Benennung

Verwenden Sie beschreibende Branch-Namen mit Praefixen:

- `feature/` - Neue Funktionen (z.B. `feature/add-region-data`)
- `fix/` - Fehlerbehebungen (z.B. `fix/currency-symbol-encoding`)
- `docs/` - Dokumentationsaktualisierungen (z.B. `docs/api-examples`)
- `refactor/` - Code-Refactoring (z.B. `refactor/generator-templates`)
- `chore/` - Wartungsaufgaben (z.B. `chore/update-dependencies`)

### Conventional Commits

Dieses Projekt folgt [Conventional Commits](https://www.conventionalcommits.org/). Format:

```
<type>[optionaler Scope]: <Beschreibung>

[optionaler Text]

[optionale(r) Footer]
```

**Typen:**

| Typ | Beschreibung |
|-----|--------------|
| `feat` | Neue Funktion |
| `fix` | Fehlerbehebung |
| `docs` | Nur Dokumentation |
| `style` | Codestil (Formatierung, Semikolons usw.) |
| `refactor` | Code-Refactoring (keine Funktion/Behebung) |
| `perf` | Leistungsverbesserung |
| `test` | Hinzufuegen oder Aktualisieren von Tests |
| `build` | Buildsystem oder Abhaengigkeiten |
| `ci` | CI/CD-Konfiguration |
| `chore` | Andere Wartungsaufgaben |

**Scopes** (optional): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Beispiele:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Einen Feature-Branch von `master` erstellen
2. Aenderungen gemaess den obigen Konventionen vornehmen
3. Sicherstellen, dass alle Tests bestehen: `make test`
4. Sicherstellen, dass das Linting besteht: `make lint`
5. Einen PR mit klarem Titel und Beschreibung oeffnen
6. Verknuepfte Issues verlinken

**PR-Titelformat:** Verwenden Sie das Conventional-Commit-Format fuer den PR-Titel.

**PR-Beschreibungsvorlage:**

```markdown
## Zusammenfassung
Kurze Beschreibung der Aenderungen.

## Aenderungen
- Aenderung 1
- Aenderung 2

## Tests
Wie die Aenderungen getestet wurden.

## Verwandte Issues
Fixes #123
```

## Lizenz

MIT
