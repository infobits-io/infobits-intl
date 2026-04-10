---
title: Reference API
sidebar_position: 8
description: Reference complete de l'API pour toutes les fonctions, types, enums et cartes de donnees exportes en Go, Dart et TypeScript.
---

# Reference API

Une reference complete de toutes les API publiques disponibles dans chaque package.

## Types

### Country

| Champ | Type | Description |
|-------|------|-------------|
| `id` | string | Identifiant unique (ex. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alpha-2 (ex. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alpha-3 (ex. `"USA"`) |
| `numericCode` | number | Code numerique ISO 3166-1 (ex. `840`) |
| `nativeName` | string | Nom natif du pays |
| `capital` | string | Capitale |
| `mainLanguage` | string | Code de la langue principale |
| `languages` | string[] | Codes de toutes les langues parlees |
| `tld` | string | Domaine de premier niveau |
| `callingCode` | number | Indicatif telephonique international |
| `continent` | string | Identifiant du continent |
| `currency` | string | Code de la devise |

### Language

| Champ | Type | Description |
|-------|------|-------------|
| `id` | string | Identifiant unique (ex. `"english"`) |
| `code` | string | Code ISO 639-1 (ex. `"en"`) |
| `nativeName` | string | Nom dans la langue elle-meme |
| `dialects` | LanguageDialect[] | Variantes regionales |
| `defaultFlagCode` | string? | Code pays pour le drapeau representatif |

### LanguageDialect

| Champ | Type | Description |
|-------|------|-------------|
| `id` | string | Identifiant unique |
| `code` | string | Code du dialecte |
| `nativeName` | string | Nom natif du dialecte |
| `flagCode` | string? | Code pays pour le drapeau du dialecte |

### Currency

| Champ | Type | Description |
|-------|------|-------------|
| `id` | string | Identifiant unique (ex. `"usd"`) |
| `code` | string | Code ISO 4217 (ex. `"USD"`) |
| `nativeName` | string | Nom au singulier (ex. `"US Dollar"`) |
| `nativeNamePlural` | string | Nom au pluriel (ex. `"US Dollars"`) |
| `symbol` | string | Symbole de la devise (ex. `"$"`) |

### Continent

| Champ | Type | Description |
|-------|------|-------------|
| `id` | string | Identifiant unique (ex. `"europe"`) |
| `code` | string | Code a deux lettres (ex. `"EU"`) |
| `name` | string | Nom d'affichage en anglais |

## Fonctions de recherche

### Pays

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

### Langues

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

### Devises

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

## Drapeaux

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

## Drapeaux emoji

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

Chaque package fournit des constantes typees pour tous les codes :

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

## Fonctions de traduction

Toutes les fonctions de traduction prennent un code d'entite et un code de locale, et retournent le nom localise.

Locales disponibles : `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

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

## Cartes de traductions

Pour un acces en masse, les cartes de traductions brutes sont disponibles :

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
