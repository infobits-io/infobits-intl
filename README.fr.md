# Infobits Intl

Package d'internationalisation multi-framework fournissant des donnees sur les pays, langues, devises et continents avec traductions et drapeaux SVG.

## Packages

| Package | Plateforme | Registre | Docs |
|---------|------------|----------|------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Fonctionnalites

- 248 pays avec codes ISO, capitales, indicatifs telephoniques, TLD
- 185 langues avec codes ISO 639-1
- 179 devises avec codes ISO 4217 et symboles
- 7 continents
- Drapeaux de pays en SVG (integres en ligne)
- Traductions multilingues (danois, allemand, anglais, espagnol, francais, italien, chinois)

## Documentation

La documentation complete de l'API avec des exemples pour les trois langages est disponible sur [packages.infobits.io](https://packages.infobits.io) et dans le repertoire [`docs/`](docs/) :

- [Premiers pas](docs/getting-started.fr.md)
- [Pays](docs/countries.fr.md)
- [Langues](docs/languages.fr.md)
- [Devises](docs/currencies.fr.md)
- [Continents](docs/continents.fr.md)
- [Traductions](docs/translations.fr.md)

## Structure du projet

```
infobits_intl/
├── data/                    # Source de verite (JSON + ressources SVG)
│   ├── core/                # Fichiers de donnees principaux
│   ├── i18n/                # Traductions par type d'entite
│   └── assets/flags/        # Drapeaux de pays en SVG
├── generator/               # Generateur de code en Go
├── packages/
│   ├── dart/                # Package Flutter
│   ├── typescript/          # Package npm
│   └── go/                  # Module Go
├── generator.yaml           # Configuration du generateur
└── Makefile                 # Commandes de build
```

## Commandes

```bash
# Generer tous les packages
make generate

# Generer un package specifique
make generate-dart
make generate-typescript
make generate-go

# Valider les fichiers de donnees
make validate

# Lancer les tests
make test              # Tous les packages
make test-dart
make test-typescript
make test-go

# Verifier le code
make lint              # Generateur + package Go
make lint-generator
make lint-go
```

## CI/CD

Les workflows GitHub Actions sont configures pour :

- **Lint** (`lint.yml`) : S'execute lors des push/PR vers master
  - Linting Go (generateur et package Go)
  - Analyse Dart
  - Verification des types TypeScript

- **Test** (`test.yml`) : S'execute lors des push/PR vers master
  - Tests du generateur et validation des donnees
  - Tests des packages Go, Dart et TypeScript

- **Deploy** (`deploy.yml`) : Declenchement manuel
  - Publication sur npm et/ou pub.dev
  - Supporte le mode dry-run
  - Remplacement de version optionnel

## Developpement

### Ajouter de nouvelles donnees

1. Ajouter une entree dans le fichier `data/core/*.json` approprie
2. Ajouter les traductions dans les fichiers `data/i18n/{entity}/*.json`
3. Pour les pays : ajouter le drapeau SVG dans `data/assets/flags/`
4. Lancer `make generate`

### Flux de travail

1. Modifier les fichiers source dans `data/`
2. Lancer `make validate` pour verifier l'integrite des donnees
3. Lancer `make generate` pour regenerer tous les packages
4. Lancer `make test` pour verifier le code genere
5. Commiter les modifications

## Contribuer

### Nommage des branches

Utilisez des noms de branches descriptifs avec des prefixes :

- `feature/` - Nouvelles fonctionnalites (ex. `feature/add-region-data`)
- `fix/` - Corrections de bugs (ex. `fix/currency-symbol-encoding`)
- `docs/` - Mises a jour de la documentation (ex. `docs/api-examples`)
- `refactor/` - Refactorisation du code (ex. `refactor/generator-templates`)
- `chore/` - Taches de maintenance (ex. `chore/update-dependencies`)

### Conventional Commits

Ce projet suit la convention [Conventional Commits](https://www.conventionalcommits.org/). Format :

```
<type>[scope optionnel]: <description>

[corps optionnel]

[pied de page optionnel]
```

**Types :**

| Type | Description |
|------|-------------|
| `feat` | Nouvelle fonctionnalite |
| `fix` | Correction de bug |
| `docs` | Documentation uniquement |
| `style` | Style du code (formatage, points-virgules, etc.) |
| `refactor` | Refactorisation du code (sans fonctionnalite/correction) |
| `perf` | Amelioration des performances |
| `test` | Ajout ou mise a jour de tests |
| `build` | Systeme de build ou dependances |
| `ci` | Configuration CI/CD |
| `chore` | Autres taches de maintenance |

**Scopes** (optionnel) : `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Exemples :**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Creer une branche depuis `master`
2. Effectuer les modifications en suivant les conventions ci-dessus
3. S'assurer que tous les tests passent : `make test`
4. S'assurer que le linting passe : `make lint`
5. Ouvrir une PR avec un titre et une description clairs
6. Lier les issues associees

**Format du titre de la PR :** Utilisez le format conventional commit pour le titre de la PR.

**Modele de description de PR :**

```markdown
## Resume
Breve description des modifications.

## Modifications
- Modification 1
- Modification 2

## Tests
Comment les modifications ont ete testees.

## Issues associees
Fixes #123
```

## Licence

MIT
