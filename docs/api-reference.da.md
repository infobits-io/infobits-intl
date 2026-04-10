---
title: API-reference
sidebar_position: 8
description: Komplet API-reference for alle eksporterede funktioner, typer, enums og datamaps paa tvaers af Go, Dart og TypeScript.
---

# API-reference

En komplet reference over alle offentlige API'er tilgaengelige i hver pakke.

## Typer

### Country

| Felt | Type | Beskrivelse |
|------|------|-------------|
| `id` | string | Unik identifikator (f.eks. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alpha-2 (f.eks. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alpha-3 (f.eks. `"USA"`) |
| `numericCode` | number | ISO 3166-1 numerisk (f.eks. `840`) |
| `nativeName` | string | Oprindeligt landenavn |
| `capital` | string | Hovedstad |
| `mainLanguage` | string | Primaer sprogkode |
| `languages` | string[] | Alle talte sprogkoder |
| `tld` | string | Topdomaene |
| `callingCode` | number | Internationalt opkaldsnummer |
| `continent` | string | Kontinent-identifikator |
| `currency` | string | Valutakode |

### Language

| Felt | Type | Beskrivelse |
|------|------|-------------|
| `id` | string | Unik identifikator (f.eks. `"english"`) |
| `code` | string | ISO 639-1-kode (f.eks. `"en"`) |
| `nativeName` | string | Navn paa selve sproget |
| `dialects` | LanguageDialect[] | Regionale varianter |
| `defaultFlagCode` | string? | Landekode for repraesentativt flag |

### LanguageDialect

| Felt | Type | Beskrivelse |
|------|------|-------------|
| `id` | string | Unik identifikator |
| `code` | string | Dialektkode |
| `nativeName` | string | Oprindeligt navn paa dialekten |
| `flagCode` | string? | Landekode for dialektens flag |

### Currency

| Felt | Type | Beskrivelse |
|------|------|-------------|
| `id` | string | Unik identifikator (f.eks. `"usd"`) |
| `code` | string | ISO 4217-kode (f.eks. `"USD"`) |
| `nativeName` | string | Ental-navn (f.eks. `"US Dollar"`) |
| `nativeNamePlural` | string | Flertal-navn (f.eks. `"US Dollars"`) |
| `symbol` | string | Valutasymbol (f.eks. `"$"`) |

### Continent

| Felt | Type | Beskrivelse |
|------|------|-------------|
| `id` | string | Unik identifikator (f.eks. `"europe"`) |
| `code` | string | Tobogstavskode (f.eks. `"EU"`) |
| `name` | string | Engelsk visningsnavn |

## Opslagsfunktioner

### Lande

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

### Sprog

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

### Valutaer

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

### Kontinenter

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

## Flag

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

## Emoji-flag

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

## Enums / Konstanter

Hver pakke tilbyder typesikre konstanter for alle koder:

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

## Oversaettelsesfunktioner

Alle oversaettelsesfunktioner tager en enhedskode og en lokalitetskode og returnerer det lokaliserede navn.

Tilgaengelige lokaliteter: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

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

## Oversaettelsesmaps

For masseadgang er de raa oversaettelsesmaps tilgaengelige:

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
