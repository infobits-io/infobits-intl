---
title: API Reference
sidebar_position: 8
description: Complete API reference for all exported functions, types, enums, and data maps across Go, Dart, and TypeScript.
---

# API Reference

A complete reference of all public APIs available in each package.

## Types

### Country

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier (e.g. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alpha-2 (e.g. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alpha-3 (e.g. `"USA"`) |
| `numericCode` | number | ISO 3166-1 numeric (e.g. `840`) |
| `nativeName` | string | Native country name |
| `capital` | string | Capital city |
| `mainLanguage` | string | Primary language code |
| `languages` | string[] | All spoken language codes |
| `tld` | string | Top-level domain |
| `callingCode` | number | International calling code |
| `continent` | string | Continent identifier |
| `currency` | string | Currency code |

### Language

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier (e.g. `"english"`) |
| `code` | string | ISO 639-1 code (e.g. `"en"`) |
| `nativeName` | string | Name in the language itself |
| `dialects` | LanguageDialect[] | Regional variations |
| `defaultFlagCode` | string? | Country code for representative flag |

### LanguageDialect

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier |
| `code` | string | Dialect code |
| `nativeName` | string | Native name of the dialect |
| `flagCode` | string? | Country code for the dialect's flag |

### Currency

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier (e.g. `"usd"`) |
| `code` | string | ISO 4217 code (e.g. `"USD"`) |
| `nativeName` | string | Singular name (e.g. `"US Dollar"`) |
| `nativeNamePlural` | string | Plural name (e.g. `"US Dollars"`) |
| `symbol` | string | Currency symbol (e.g. `"$"`) |

### Continent

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier (e.g. `"europe"`) |
| `code` | string | Two-letter code (e.g. `"EU"`) |
| `name` | string | English display name |

## Lookup Functions

### Countries

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

### Languages

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

### Currencies

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

### Continents

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

## Flags

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

## Emoji Flags

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

## Enums / Constants

Each package provides type-safe constants for all codes:

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

## Translation Functions

All translation functions take an entity code and a locale code, returning the localized name.

Available locales: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

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

## Translation Maps

For bulk access, raw translation maps are available:

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
