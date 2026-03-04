# Infobits Intl

Paquete de internacionalizacion multientorno que proporciona datos de paises, idiomas, monedas y continentes con traducciones y banderas SVG.

## Paquetes

| Paquete | Plataforma | Registro | Documentacion |
|---------|------------|----------|---------------|
| `infobits_intl` | Dart/Flutter | [pub.dev](https://pub.dev/packages/infobits_intl) | [README](packages/dart/README.md) |
| `@infobits/intl` | TypeScript/JavaScript | [npm](https://www.npmjs.com/package/@infobits/intl) | [README](packages/typescript/README.md) |
| `github.com/infobits-io/infobits-intl-go` | Go | [pkg.go.dev](https://pkg.go.dev/github.com/infobits-io/infobits-intl-go) | [README](packages/go/README.md) |

## Caracteristicas

- 248 paises con codigos ISO, capitales, codigos de llamada, TLDs
- 185 idiomas con codigos ISO 639-1
- 179 monedas con codigos ISO 4217 y simbolos
- 7 continentes
- Banderas de paises en SVG (incrustadas en linea)
- Traducciones multilingues (danes, aleman, ingles, espanol, frances, italiano, chino)

## Documentacion

La documentacion completa de la API con ejemplos para los tres lenguajes esta disponible en [packages.infobits.io](https://packages.infobits.io) y en el directorio [`docs/`](docs/):

- [Primeros pasos](docs/getting-started.es.md)
- [Paises](docs/countries.es.md)
- [Idiomas](docs/languages.es.md)
- [Monedas](docs/currencies.es.md)
- [Continentes](docs/continents.es.md)
- [Traducciones](docs/translations.es.md)

## Estructura del proyecto

```
infobits_intl/
├── data/                    # Fuente de datos (JSON + activos SVG)
│   ├── core/                # Archivos de datos principales
│   ├── i18n/                # Traducciones por tipo de entidad
│   └── assets/flags/        # Banderas de paises en SVG
├── generator/               # Generador de codigo en Go
├── packages/
│   ├── dart/                # Paquete Flutter
│   ├── typescript/          # Paquete npm
│   └── go/                  # Modulo Go
├── generator.yaml           # Configuracion del generador
└── Makefile                 # Comandos de compilacion
```

## Comandos

```bash
# Generar todos los paquetes
make generate

# Generar un paquete especifico
make generate-dart
make generate-typescript
make generate-go

# Validar archivos de datos
make validate

# Ejecutar pruebas
make test              # Todos los paquetes
make test-dart
make test-typescript
make test-go

# Analizar codigo
make lint              # Generador + paquete Go
make lint-generator
make lint-go
```

## CI/CD

Los workflows de GitHub Actions estan configurados para:

- **Lint** (`lint.yml`): Se ejecuta en push/PR a master
  - Linting de Go (generador y paquete Go)
  - Analisis de Dart
  - Verificacion de tipos de TypeScript

- **Test** (`test.yml`): Se ejecuta en push/PR a master
  - Pruebas del generador y validacion de datos
  - Pruebas de paquetes Go, Dart y TypeScript

- **Deploy** (`deploy.yml`): Dispatch manual del workflow
  - Publicacion en npm y/o pub.dev
  - Compatible con modo dry-run
  - Anulacion de version opcional

## Desarrollo

### Agregar nuevos datos

1. Agregar entrada al archivo `data/core/*.json` correspondiente
2. Agregar traducciones a los archivos `data/i18n/{entity}/*.json`
3. Para paises: agregar bandera SVG a `data/assets/flags/`
4. Ejecutar `make generate`

### Flujo de trabajo

1. Editar archivos fuente en `data/`
2. Ejecutar `make validate` para verificar la integridad de los datos
3. Ejecutar `make generate` para regenerar todos los paquetes
4. Ejecutar `make test` para verificar el codigo generado
5. Hacer commit de los cambios

## Contribuir

### Nomenclatura de ramas

Use nombres de rama descriptivos con prefijos:

- `feature/` - Nuevas funcionalidades (ej. `feature/add-region-data`)
- `fix/` - Correcciones de errores (ej. `fix/currency-symbol-encoding`)
- `docs/` - Actualizaciones de documentacion (ej. `docs/api-examples`)
- `refactor/` - Refactorizacion de codigo (ej. `refactor/generator-templates`)
- `chore/` - Tareas de mantenimiento (ej. `chore/update-dependencies`)

### Conventional Commits

Este proyecto sigue [Conventional Commits](https://www.conventionalcommits.org/). Formato:

```
<type>[scope opcional]: <descripcion>

[cuerpo opcional]

[pie(s) de pagina opcional(es)]
```

**Tipos:**

| Tipo | Descripcion |
|------|-------------|
| `feat` | Nueva funcionalidad |
| `fix` | Correccion de error |
| `docs` | Solo documentacion |
| `style` | Estilo de codigo (formato, punto y coma, etc.) |
| `refactor` | Refactorizacion de codigo (sin funcionalidad/correccion) |
| `perf` | Mejora de rendimiento |
| `test` | Agregar o actualizar pruebas |
| `build` | Sistema de compilacion o dependencias |
| `ci` | Configuracion de CI/CD |
| `chore` | Otras tareas de mantenimiento |

**Scopes** (opcional): `dart`, `typescript`, `go`, `generator`, `data`, `ci`

**Ejemplos:**

```bash
feat(dart): add region subdivision support
fix(generator): handle special characters in currency symbols
docs: update installation instructions
chore(ci): add Node.js 22 to test matrix
```

### Pull Requests

1. Crear una rama de funcionalidad desde `master`
2. Realizar cambios siguiendo las convenciones anteriores
3. Asegurarse de que todas las pruebas pasen: `make test`
4. Asegurarse de que el linting pase: `make lint`
5. Abrir un PR con un titulo y descripcion claros
6. Vincular los issues relacionados

**Formato del titulo del PR:** Use el formato de conventional commit para el titulo del PR.

**Plantilla de descripcion del PR:**

```markdown
## Resumen
Breve descripcion de los cambios.

## Cambios
- Cambio 1
- Cambio 2

## Pruebas
Como se probaron los cambios.

## Issues relacionados
Fixes #123
```

## Licencia

MIT
