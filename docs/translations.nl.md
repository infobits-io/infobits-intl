---
title: Vertalingen
sidebar_position: 6
description: Hoe i18n-vertalingen werken — gelokaliseerde namen ophalen voor landen, talen, valuta's, continenten en hoofdsteden.
---

# Vertalingen

Alle entiteitsnamen (landen, talen, valuta's, continenten en hoofdsteden) kunnen in meerdere talen worden opgehaald.

## Beschikbare taalgebieden

| Code | Taal |
|------|------|
| `da` | Deens |
| `de` | Duits |
| `en` | Engels |
| `es` | Spaans |
| `fr` | Frans |
| `it` | Italiaans |
| `zh` | Chinees |

## Landnamen

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

// Met een Locale-object
final name = country?.displayNameFromLocale(Locale('de'));
print(name); // Vereinigte Staaten

// Met BuildContext (in een widget)
final contextName = country?.displayName(context);
```

```typescript
import { getCountriesName } from 'infobits-intl';

const name = getCountriesName('US', 'de');
console.log(name); // Vereinigte Staaten

const nameFr = getCountriesName('FR', 'es');
console.log(nameFr); // Francia
```

## Taalnamen

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

## Valutanamen

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

## Continentnamen

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

## Hoofdstednamen

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

## Vertaalmappen rechtstreeks benaderen

Voor bulkbewerkingen kunt u de volledige vertaalmappen benaderen.

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Toegang tot alle Duitse landvertalingen
germanCountries := i18n.CountriesTranslations["de"]
for code, name := range germanCountries {
    fmt.Printf("%s: %s\n", code, name)
}

// Andere beschikbare mappen:
// i18n.LanguagesTranslations
// i18n.CurrenciesTranslations
// i18n.ContinentsTranslations
// i18n.CapitalsTranslations
```

```dart
// Vertaalmappen zijn toegankelijk via delegate-klassen
final germanCountries = CountriesTranslationsDelegate.translations['de'];
germanCountries?.forEach((code, name) {
  print('$code: $name');
});

// Andere beschikbare delegates:
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

// Toegang tot alle Duitse landvertalingen
const germanCountries = countriesTranslations['de'];
for (const [code, name] of Object.entries(germanCountries)) {
  console.log(`${code}: ${name}`);
}
```
