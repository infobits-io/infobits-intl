---
title: Übersetzungen
sidebar_position: 6
description: Wie i18n-Übersetzungen funktionieren — lokalisierte Namen für Länder, Sprachen, Währungen, Kontinente und Hauptstädte abrufen.
---

# Übersetzungen

Alle Entitätsnamen (Länder, Sprachen, Währungen, Kontinente und Hauptstädte) können in mehreren Sprachen abgerufen werden.

## Verfügbare Sprachen

| Code | Sprache |
|------|---------|
| `da` | Dänisch |
| `de` | Deutsch |
| `en` | Englisch |
| `es` | Spanisch |
| `fr` | Französisch |
| `it` | Italienisch |
| `zh` | Chinesisch |

## Ländernamen

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

// Mit einem Locale-Objekt
final name = country?.displayNameFromLocale(Locale('de'));
print(name); // Vereinigte Staaten

// Mit BuildContext (in einem Widget)
final contextName = country?.displayName(context);
```

```typescript
import { getCountriesName } from 'infobits-intl';

const name = getCountriesName('US', 'de');
console.log(name); // Vereinigte Staaten

const nameFr = getCountriesName('FR', 'es');
console.log(nameFr); // Francia
```

## Sprachnamen

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

## Währungsnamen

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

## Kontinentnamen

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

## Hauptstadtnamen

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

## Direkter Zugriff auf Übersetzungs-Maps

Für Massenoperationen können Sie auf die vollständigen Übersetzungs-Maps zugreifen.

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Auf alle deutschen Länderübersetzungen zugreifen
germanCountries := i18n.CountriesTranslations["de"]
for code, name := range germanCountries {
    fmt.Printf("%s: %s\n", code, name)
}

// Weitere verfügbare Maps:
// i18n.LanguagesTranslations
// i18n.CurrenciesTranslations
// i18n.ContinentsTranslations
// i18n.CapitalsTranslations
```

```dart
// Übersetzungs-Maps sind über Delegate-Klassen zugänglich
final germanCountries = CountriesTranslationsDelegate.translations['de'];
germanCountries?.forEach((code, name) {
  print('$code: $name');
});

// Weitere verfügbare Delegates:
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

// Auf alle deutschen Länderübersetzungen zugreifen
const germanCountries = countriesTranslations['de'];
for (const [code, name] of Object.entries(germanCountries)) {
  console.log(`${code}: ${name}`);
}
```
