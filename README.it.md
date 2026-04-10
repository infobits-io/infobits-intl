# Infobits Intl

Pacchetto di internazionalizzazione multi-framework che fornisce dati su paesi, lingue, valute e continenti con traduzioni e bandiere SVG.

## Pacchetti

| Pacchetto | Piattaforma | Registro | Documentazione |
|-----------|-------------|----------|----------------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Funzionalita

- 248 paesi con codici ISO, capitali, prefissi telefonici, TLD
- 185 lingue con codici ISO 639-1
- 179 valute con codici ISO 4217 e simboli
- 7 continenti
- Bandiere dei paesi in SVG (incorporate inline)
- Traduzioni multilingue (danese, tedesco, inglese, spagnolo, francese, italiano, cinese)

## Documentazione

La documentazione completa dell'API con esempi per tutti e tre i linguaggi e disponibile su [docs.infobits.io](https://docs.infobits.io) e nella directory [`docs/`](docs/):

- [Per Iniziare](docs/getting-started.it.md)
- [Paesi](docs/countries.it.md)
- [Lingue](docs/languages.it.md)
- [Valute](docs/currencies.it.md)
- [Continenti](docs/continents.it.md)
- [Traduzioni](docs/translations.it.md)

## Struttura del progetto

```
infobits_intl/
├── data/                    # Fonte di verita (JSON + risorse SVG)
│   ├── core/                # File di dati principali
│   ├── i18n/                # Traduzioni per tipo di entita
│   └── assets/flags/        # Bandiere dei paesi in SVG
├── generator/               # Generatore di codice in Go
├── packages/
│   ├── dart/                # Pacchetto Flutter
│   ├── typescript/          # Pacchetto npm
│   └── go/                  # Modulo Go
├── generator.yaml           # Configurazione del generatore
└── Makefile                 # Comandi di build
```

## Comandi

```bash
# Generare tutti i pacchetti
make generate

# Generare un pacchetto specifico
make generate-dart
make generate-typescript
make generate-go

# Validare i file di dati
make validate

# Eseguire i test
make test              # Tutti i pacchetti
make test-dart
make test-typescript
make test-go

# Controllare il codice
make lint              # Generatore + pacchetto Go
make lint-generator
make lint-go
```

## CI/CD

I workflow di GitHub Actions sono configurati per:

- **Lint** (`lint.yml`): Eseguito su push/PR verso master
  - Linting Go (generatore e pacchetto Go)
  - Analisi Dart
  - Controllo dei tipi TypeScript

- **Test** (`test.yml`): Eseguito su push/PR verso master
  - Test del generatore e validazione dei dati
  - Test dei pacchetti Go, Dart e TypeScript

- **Deploy** (`deploy.yml`): Attivazione manuale
  - Pubblicazione su npm e/o pub.dev
  - Supporto modalita dry-run
  - Override della versione opzionale

## Sviluppo

### Aggiungere nuovi dati

1. Aggiungere una voce nel file `data/core/*.json` appropriato
2. Aggiungere le traduzioni nei file `data/i18n/{entity}/*.json`
3. Per i paesi: aggiungere la bandiera SVG in `data/assets/flags/`
4. Eseguire `make generate`

### Flusso di lavoro

1. Modificare i file sorgente in `data/`
2. Eseguire `make validate` per verificare l'integrita dei dati
3. Eseguire `make generate` per rigenerare tutti i pacchetti
4. Eseguire `make test` per verificare il codice generato
5. Effettuare il commit delle modifiche

## Contribuire

### Denominazione dei branch

Utilizzare nomi di branch descrittivi con prefissi:

- `feature/` - Nuove funzionalita (es. `feature/add-region-data`)
- `fix/` - Correzioni di bug (es. `fix/currency-symbol-encoding`)
- `docs/` - Aggiornamenti della documentazione (es. `docs/api-examples`)
- `refactor/` - Refactoring del codice (es. `refactor/generator-templates`)
- `chore/` - Attivita di manutenzione (es. `chore/update-dependencies`)

### Conventional Commits

Questo progetto segue la convenzione [Conventional Commits](https://www.conventionalcommits.org/). Formato:

```
<type>[scope opzionale]: <descrizione>

[corpo opzionale]

[piede di pagina opzionale]
```

**Tipi:**

| Tipo | Descrizione |
|------|-------------|
| `feat` | Nuova funzionalita |
| `fix` | Correzione di bug |
| `docs` | Solo documentazione |
| `style` | Stile del codice (formattazione, punti e virgola, ecc.) |
| `refactor` | Refactoring del codice (nessuna funzionalita/correzione) |
| `perf` | Miglioramento delle prestazioni |
| `test` | Aggiunta o aggiornamento di test |
| `build` | Sistema di build o dipendenze |
| `ci` | Configurazione CI/CD |
| `chore` | Altre attivita di manutenzione |

**Scope** (opzionale): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Esempi:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Request

1. Creare un branch da `master`
2. Apportare le modifiche seguendo le convenzioni sopra indicate
3. Assicurarsi che tutti i test passino: `make test`
4. Assicurarsi che il linting passi: `make lint`
5. Aprire una PR con un titolo e una descrizione chiari
6. Collegare le issue correlate

**Formato del titolo della PR:** Utilizzare il formato conventional commit per il titolo della PR.

**Modello di descrizione della PR:**

```markdown
## Riepilogo
Breve descrizione delle modifiche.

## Modifiche
- Modifica 1
- Modifica 2

## Test
Come sono state testate le modifiche.

## Issue correlate
Fixes #123
```

## Licenza

MIT
