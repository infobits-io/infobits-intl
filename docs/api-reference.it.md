---
title: Riferimento API
sidebar_position: 8
description: Riferimento completo dell'API per tutte le funzioni, tipi, enum e mappe dati esportati in Go, Dart e TypeScript.
---

# Riferimento API

Un riferimento completo di tutte le API pubbliche disponibili in ciascun pacchetto.

## Tipi

### Country

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | string | Identificatore univoco (es. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alpha-2 (es. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alpha-3 (es. `"USA"`) |
| `numericCode` | number | ISO 3166-1 numerico (es. `840`) |
| `nativeName` | string | Nome nativo del paese |
| `capital` | string | Capitale |
| `mainLanguage` | string | Codice della lingua principale |
| `languages` | string[] | Codici di tutte le lingue parlate |
| `tld` | string | Dominio di primo livello |
| `callingCode` | number | Prefisso telefonico internazionale |
| `continent` | string | Identificatore del continente |
| `currency` | string | Codice della valuta |

### Language

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | string | Identificatore univoco (es. `"english"`) |
| `code` | string | Codice ISO 639-1 (es. `"en"`) |
| `nativeName` | string | Nome nella lingua stessa |
| `dialects` | LanguageDialect[] | Varianti regionali |
| `defaultFlagCode` | string? | Codice paese per la bandiera rappresentativa |

### LanguageDialect

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | string | Identificatore univoco |
| `code` | string | Codice del dialetto |
| `nativeName` | string | Nome nativo del dialetto |
| `flagCode` | string? | Codice paese per la bandiera del dialetto |

### Currency

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | string | Identificatore univoco (es. `"usd"`) |
| `code` | string | Codice ISO 4217 (es. `"USD"`) |
| `nativeName` | string | Nome singolare (es. `"US Dollar"`) |
| `nativeNamePlural` | string | Nome plurale (es. `"US Dollars"`) |
| `symbol` | string | Simbolo della valuta (es. `"$"`) |

### Continent

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | string | Identificatore univoco (es. `"europe"`) |
| `code` | string | Codice a due lettere (es. `"EU"`) |
| `name` | string | Nome visualizzato in inglese |

## Funzioni di Ricerca

### Paesi

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

### Lingue

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

### Valute

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

### Continenti

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

## Bandiere

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

## Bandiere Emoji

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

## Enum / Costanti

Ogni pacchetto fornisce costanti type-safe per tutti i codici:

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

## Funzioni di Traduzione

Tutte le funzioni di traduzione accettano un codice entita e un codice locale, restituendo il nome localizzato.

Lingue disponibili: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

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

## Mappe di Traduzione

Per l'accesso in blocco, sono disponibili le mappe di traduzione grezze:

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
