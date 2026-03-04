---
title: Translations
sidebar_position: 6
description: How i18n translations work — getting localized names for countries, languages, currencies, continents, and capitals.
---

# Translations

All entity names (countries, languages, currencies, continents, and capitals) can be retrieved in multiple languages.

## Available Locales

| Code | Language |
|------|----------|
| `da` | Danish |
| `de` | German |
| `en` | English |
| `es` | Spanish |
| `fr` | French |
| `it` | Italian |
| `zh` | Chinese |

## Country Names

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

name, ok := i18n.GetCountriesName("US", "de")
if ok {
    fmt.Println(name) // Vereinigte Staaten
}

name, _ = i18n.GetCountriesName("FR", "es")
fmt.Println(name) // Francia
```

```dart
final country = Country.fromAlpha2Code('US');

// Using a Locale object
final name = country?.displayNameFromLocale(Locale('de'));
print(name); // Vereinigte Staaten

// Using BuildContext (in a widget)
final contextName = country?.displayName(context);
```

```typescript
import { getCountriesName } from 'infobits-intl';

const name = getCountriesName('US', 'de');
console.log(name); // Vereinigte Staaten

const nameFr = getCountriesName('FR', 'es');
console.log(nameFr); // Francia
```

## Language Names

```go
name, ok := i18n.GetLanguagesName("en", "fr")
if ok {
    fmt.Println(name) // anglais
}
```

```dart
final lang = Language.fromCode('en');
final name = lang?.displayNameFromLocale(Locale('fr'));
print(name); // anglais
```

```typescript
import { getLanguagesName } from 'infobits-intl';

const name = getLanguagesName('en', 'fr');
console.log(name); // anglais
```

## Currency Names

```go
name, ok := i18n.GetCurrenciesName("USD", "it")
if ok {
    fmt.Println(name) // Dollaro statunitense
}
```

```dart
final currency = Currency.fromCode('USD');
final name = currency?.displayNameFromLocale(Locale('it'));
print(name); // Dollaro statunitense
```

```typescript
import { getCurrenciesName } from 'infobits-intl';

const name = getCurrenciesName('USD', 'it');
console.log(name); // Dollaro statunitense
```

## Continent Names

```go
name, ok := i18n.GetContinentsName("EU", "da")
if ok {
    fmt.Println(name) // Europa
}
```

```dart
final continent = Continent.fromCode('EU');
final name = continent?.displayNameFromLocale(Locale('da'));
print(name); // Europa
```

```typescript
import { getContinentsName } from 'infobits-intl';

const name = getContinentsName('EU', 'da');
console.log(name); // Europa
```

## Capital Names

```go
name, ok := i18n.GetCapitalsName("JP", "fr")
if ok {
    fmt.Println(name) // Tokyo
}
```

```dart
final country = Country.fromAlpha2Code('JP');
final name = country?.displayCapitalFromLocale(Locale('fr'));
print(name); // Tokyo
```

```typescript
import { getCapitalsName } from 'infobits-intl';

const name = getCapitalsName('JP', 'fr');
console.log(name); // Tokyo
```

## Accessing Translation Maps Directly

For bulk operations, you can access the full translation maps.

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Access all German country translations
germanCountries := i18n.CountriesTranslations["de"]
for code, name := range germanCountries {
    fmt.Printf("%s: %s\n", code, name)
}

// Other available maps:
// i18n.LanguagesTranslations
// i18n.CurrenciesTranslations
// i18n.ContinentsTranslations
// i18n.CapitalsTranslations
```

```dart
// Translation maps are accessible via delegate classes
final germanCountries = CountriesTranslationsDelegate.translations['de'];
germanCountries?.forEach((code, name) {
  print('$code: $name');
});

// Other available delegates:
// LanguagesTranslationsDelegate.translations
// CurrenciesTranslationsDelegate.translations
// ContinentsTranslationsDelegate.translations
// CapitalsTranslationsDelegate.translations
```

```typescript
import {
  countriesTranslations,
  languagesTranslations,
  currenciesTranslations,
  continentsTranslations,
  capitalsTranslations
} from 'infobits-intl';

// Access all German country translations
const germanCountries = countriesTranslations['de'];
for (const [code, name] of Object.entries(germanCountries)) {
  console.log(`${code}: ${name}`);
}
```
