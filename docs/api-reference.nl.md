---
title: API-referentie
sidebar_position: 8
description: Volledige API-referentie voor alle geexporteerde functies, typen, enums en datamaps in Go, Dart en TypeScript.
---

# API-referentie

Een volledig overzicht van alle publieke API's die beschikbaar zijn in elk pakket.

## Typen

### Country

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | string | Unieke identificatie (bijv. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 alfa-2 (bijv. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 alfa-3 (bijv. `"USA"`) |
| `numericCode` | number | ISO 3166-1 numeriek (bijv. `840`) |
| `nativeName` | string | Oorspronkelijke landnaam |
| `capital` | string | Hoofdstad |
| `mainLanguage` | string | Primaire taalcode |
| `languages` | string[] | Alle gesproken taalcodes |
| `tld` | string | Topniveaudomein |
| `callingCode` | number | Internationaal landnummer |
| `continent` | string | Continentidentificatie |
| `currency` | string | Valutacode |

### Language

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | string | Unieke identificatie (bijv. `"english"`) |
| `code` | string | ISO 639-1-code (bijv. `"en"`) |
| `nativeName` | string | Naam in de taal zelf |
| `dialects` | LanguageDialect[] | Regionale varianten |
| `defaultFlagCode` | string? | Landcode voor representatieve vlag |

### LanguageDialect

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | string | Unieke identificatie |
| `code` | string | Dialectcode |
| `nativeName` | string | Oorspronkelijke naam van het dialect |
| `flagCode` | string? | Landcode voor de vlag van het dialect |

### Currency

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | string | Unieke identificatie (bijv. `"usd"`) |
| `code` | string | ISO 4217-code (bijv. `"USD"`) |
| `nativeName` | string | Enkelvoudige naam (bijv. `"US Dollar"`) |
| `nativeNamePlural` | string | Meervoudige naam (bijv. `"US Dollars"`) |
| `symbol` | string | Valutasymbool (bijv. `"$"`) |

### Continent

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | string | Unieke identificatie (bijv. `"europe"`) |
| `code` | string | Tweeletterige code (bijv. `"EU"`) |
| `name` | string | Engelse weergavenaam |

## Opzoekfuncties

### Landen

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

### Talen

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

### Valuta's

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

### Continenten

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

## Vlaggen

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

## Emojivlaggen

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

## Enums / Constanten

Elk pakket biedt typeveilige constanten voor alle codes:

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

## Vertaalfuncties

Alle vertaalfuncties nemen een entiteitcode en een taalgebiedcode en retourneren de gelokaliseerde naam.

Beschikbare taalgebieden: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

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

## Vertaalmappen

Voor bulktoegang zijn de ruwe vertaalmappen beschikbaar:

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
