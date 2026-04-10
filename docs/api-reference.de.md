---
title: API-Referenz
sidebar_position: 8
description: Vollstaendige API-Referenz fuer alle exportierten Funktionen, Typen, Enums und Datenmaps fuer Go, Dart und TypeScript.
---

# API-Referenz

Eine vollständige Referenz aller öffentlichen APIs in jedem Paket.

## Typen

### Country

| Feld | Typ | Beschreibung |
|------|-----|--------------|
| `id` | string | Eindeutiger Bezeichner (z.B. `"unitedStates"`) |
| `alpha2Code` | string | ISO 3166-1 Alpha-2 (z.B. `"US"`) |
| `alpha3Code` | string | ISO 3166-1 Alpha-3 (z.B. `"USA"`) |
| `numericCode` | number | ISO 3166-1 numerisch (z.B. `840`) |
| `nativeName` | string | Einheimischer Ländername |
| `capital` | string | Hauptstadt |
| `mainLanguage` | string | Primärer Sprachcode |
| `languages` | string[] | Alle gesprochenen Sprachcodes |
| `tld` | string | Top-Level-Domain |
| `callingCode` | number | Internationale Vorwahl |
| `continent` | string | Kontinent-Bezeichner |
| `currency` | string | Währungscode |

### Language

| Feld | Typ | Beschreibung |
|------|-----|--------------|
| `id` | string | Eindeutiger Bezeichner (z.B. `"english"`) |
| `code` | string | ISO 639-1 Code (z.B. `"en"`) |
| `nativeName` | string | Name in der jeweiligen Sprache |
| `dialects` | LanguageDialect[] | Regionale Varianten |
| `defaultFlagCode` | string? | Ländercode für repräsentative Flagge |

### LanguageDialect

| Feld | Typ | Beschreibung |
|------|-----|--------------|
| `id` | string | Eindeutiger Bezeichner |
| `code` | string | Dialektcode |
| `nativeName` | string | Einheimischer Name des Dialekts |
| `flagCode` | string? | Ländercode für die Flagge des Dialekts |

### Currency

| Feld | Typ | Beschreibung |
|------|-----|--------------|
| `id` | string | Eindeutiger Bezeichner (z.B. `"usd"`) |
| `code` | string | ISO 4217 Code (z.B. `"USD"`) |
| `nativeName` | string | Singular-Name (z.B. `"US Dollar"`) |
| `nativeNamePlural` | string | Plural-Name (z.B. `"US Dollars"`) |
| `symbol` | string | Währungssymbol (z.B. `"$"`) |

### Continent

| Feld | Typ | Beschreibung |
|------|-----|--------------|
| `id` | string | Eindeutiger Bezeichner (z.B. `"europe"`) |
| `code` | string | Zweibuchstaben-Code (z.B. `"EU"`) |
| `name` | string | Englischer Anzeigename |

## Suchfunktionen

### Länder

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

### Sprachen

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

### Währungen

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

### Kontinente

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

## Flaggen

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

## Emoji-Flaggen

```go
// Go
code.EmojiFlag() string  // auf CountryCode
```

```dart
// Dart
country.emojiFlag  // String
```

```typescript
// TypeScript
getEmojiFlag(alpha2: string): string
```

## Enums / Konstanten

Jedes Paket stellt typsichere Konstanten für alle Codes bereit:

```go
// Go — typisierte String-Konstanten
intl.CountryUS    // CountryCode("US")
intl.LanguageFR   // LanguageCode("FR")
intl.CurrencyEUR  // CurrencyCode("EUR")
intl.ContinentEU  // ContinentCode("EU")
```

```dart
// Dart — Enums
Country.unitedStates
Language.french
Currency.eur
Continent.europe
```

```typescript
// TypeScript — String-Enums
CountryCode.US    // "US"
LanguageCode.FR   // "FR"
CurrencyCode.EUR  // "EUR"
ContinentCode.EU  // "EU"
```

## Übersetzungsfunktionen

Alle Übersetzungsfunktionen nehmen einen Entitätscode und einen Locale-Code entgegen und geben den lokalisierten Namen zurück.

Verfügbare Sprachen: `da`, `de`, `en`, `es`, `fr`, `it`, `zh`

```go
// Go — alle im i18n-Unterpaket
i18n.GetCountriesName(code, locale string) (string, bool)
i18n.GetLanguagesName(code, locale string) (string, bool)
i18n.GetCurrenciesName(code, locale string) (string, bool)
i18n.GetContinentsName(code, locale string) (string, bool)
i18n.GetCapitalsName(code, locale string) (string, bool)
```

```dart
// Dart — über Entitäts-Instanzen
country.displayName(BuildContext context)  // verwendet App-Locale
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

## Übersetzungs-Maps

Für den Massenzugriff stehen die vollständigen Übersetzungs-Maps zur Verfügung:

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
