---
title: Referencia de API
sidebar_position: 8
description: Referencia completa de la API para todas las funciones, tipos, enums y mapas de datos exportados en Go, Dart y TypeScript.
---

# Referencia de API

Una referencia completa de todas las APIs publicas disponibles en cada paquete.

## Tipos

### Country

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | string | Identificador unico (ej. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alpha-2 (ej. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alpha-3 (ej. `"USA"`) |
| `numericCode` | number | Numerico ISO 3166-1 (ej. `840`) |
| `nativeName` | string | Nombre nativo del pais |
| `capital` | string | Ciudad capital |
| `mainLanguage` | string | Codigo del idioma principal |
| `languages` | string[] | Codigos de todos los idiomas hablados |
| `tld` | string | Dominio de nivel superior |
| `callingCode` | number | Codigo de llamada internacional |
| `continent` | string | Identificador de continente |
| `currency` | string | Codigo de moneda |

### Language

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | string | Identificador unico (ej. `"english"`) |
| `code` | string | Codigo ISO 639-1 (ej. `"en"`) |
| `nativeName` | string | Nombre en el propio idioma |
| `dialects` | LanguageDialect[] | Variaciones regionales |
| `defaultFlagCode` | string? | Codigo de pais para la bandera representativa |

### LanguageDialect

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | string | Identificador unico |
| `code` | string | Codigo de dialecto |
| `nativeName` | string | Nombre nativo del dialecto |
| `flagCode` | string? | Codigo de pais para la bandera del dialecto |

### Currency

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | string | Identificador unico (ej. `"usd"`) |
| `code` | string | Codigo ISO 4217 (ej. `"USD"`) |
| `nativeName` | string | Nombre en singular (ej. `"US Dollar"`) |
| `nativeNamePlural` | string | Nombre en plural (ej. `"US Dollars"`) |
| `symbol` | string | Simbolo de la moneda (ej. `"$"`) |

### Continent

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | string | Identificador unico (ej. `"europe"`) |
| `code` | string | Codigo de dos letras (ej. `"EU"`) |
| `name` | string | Nombre para mostrar en ingles |

## Funciones de busqueda

### Paises

```go
// Go
intl.CountryByAlpha2(code string) (Country, bool)
intl.CountryByAlpha3(code string) (Country, bool)
intl.AllCountries() []Country
intl.CountriesByContinent(continent string) []Country
```

```dart
// Dart
Country.fromAlpha2Code(String code) -> Country?
Country.fromAlpha3Code(String code) -> Country?
Country.fromNumericCode(int code) -> Country?
Country.values -> List<Country>
Continent.countries -> List<Country>
```

```typescript
// TypeScript
getCountryByAlpha2(code: string): Country | undefined
getCountryByAlpha3(code: string): Country | undefined
getCountriesByContinent(continent: string): Country[]
getCountriesInContinent(code: string): Country[]
```

### Idiomas

```go
// Go
intl.LanguageByCode(code string) (Language, bool)
intl.AllLanguages() []Language
```

```dart
// Dart
Language.fromCode(String code) -> Language?
Language.values -> List<Language>
```

```typescript
// TypeScript
getLanguageByCode(code: string): Language | undefined
```

### Monedas

```go
// Go
intl.CurrencyByCode(code string) (Currency, bool)
intl.AllCurrencies() []Currency
```

```dart
// Dart
Currency.fromCode(String code) -> Currency?
Currency.values -> List<Currency>
```

```typescript
// TypeScript
getCurrencyByCode(code: string): Currency | undefined
```

### Continentes

```go
// Go
intl.ContinentByCode(code string) (Continent, bool)
intl.AllContinents() []Continent
```

```dart
// Dart
Continent.fromCode(String code) -> Continent?
Continent.values -> List<Continent>
```

```typescript
// TypeScript
getContinentByCode(code: string): Continent | undefined
```

## Banderas

```go
// Go
intl.GetFlag(alpha2 string) (string, bool)
intl.Flags  // map[string]string
```

```dart
// Dart
countryFlags  // Map<String, String>
country.flagSvg  // String?
country.flag(shape: FlagShape, width: double, height: double)  // Widget
```

```typescript
// TypeScript
getFlag(alpha2: string): string | undefined
flags  // Record<string, string>
```

## Banderas emoji

```go
// Go
code.EmojiFlag() string  // on CountryCode
```

```dart
// Dart
country.emojiFlag  // String
```

```typescript
// TypeScript
getEmojiFlag(alpha2: string): string
```

## Enums / Constantes

Cada paquete proporciona constantes con tipo seguro para todos los codigos:

```go
// Go — typed string constants
intl.CountryUS    // CountryCode("US")
intl.LanguageFR   // LanguageCode("FR")
intl.CurrencyEUR  // CurrencyCode("EUR")
intl.ContinentEU  // ContinentCode("EU")
```

```dart
// Dart — enums
Country.unitedStates
Language.french
Currency.eur
Continent.europe
```

```typescript
// TypeScript — string enums
CountryCode.US    // "US"
LanguageCode.FR   // "FR"
CurrencyCode.EUR  // "EUR"
ContinentCode.EU  // "EU"
```

## Funciones de traduccion

Todas las funciones de traduccion reciben un codigo de entidad y un codigo de idioma, y devuelven el nombre localizado.

Idiomas disponibles: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

```go
// Go — all in the i18n sub-package
i18n.GetCountriesName(code, locale string) (string, bool)
i18n.GetLanguagesName(code, locale string) (string, bool)
i18n.GetCurrenciesName(code, locale string) (string, bool)
i18n.GetContinentsName(code, locale string) (string, bool)
i18n.GetCapitalsName(code, locale string) (string, bool)
```

```dart
// Dart — via entity instances
country.displayName(BuildContext context)  // uses app locale
country.displayNameFromLocale(Locale locale)
country.displayCapitalFromLocale(Locale locale)
language.displayNameFromLocale(Locale locale)
currency.displayNameFromLocale(Locale locale)
continent.displayNameFromLocale(Locale locale)
```

```typescript
// TypeScript
getCountriesName(code: string, locale: string): string | undefined
getLanguagesName(code: string, locale: string): string | undefined
getCurrenciesName(code: string, locale: string): string | undefined
getContinentsName(code: string, locale: string): string | undefined
getCapitalsName(code: string, locale: string): string | undefined
```

## Mapas de traducciones

Para acceso masivo, los mapas completos de traducciones estan disponibles:

```go
// Go
i18n.CountriesTranslations   // map[string]map[string]string
i18n.LanguagesTranslations
i18n.CurrenciesTranslations
i18n.ContinentsTranslations
i18n.CapitalsTranslations
```

```dart
// Dart
CountriesTranslationsDelegate.translations
LanguagesTranslationsDelegate.translations
CurrenciesTranslationsDelegate.translations
ContinentsTranslationsDelegate.translations
CapitalsTranslationsDelegate.translations
```

```typescript
// TypeScript
countriesTranslations   // Record<string, Record<string, string>>
languagesTranslations
currenciesTranslations
continentsTranslations
capitalsTranslations
```
