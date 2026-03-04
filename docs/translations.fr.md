---
title: Traductions
sidebar_position: 6
description: Fonctionnement des traductions i18n -- obtenir les noms localises des pays, langues, devises, continents et capitales.
---

# Traductions

Tous les noms d'entites (pays, langues, devises, continents et capitales) peuvent etre recuperes dans plusieurs langues.

## Langues disponibles

| Code | Langue |
|------|--------|
| `da` | Danois |
| `de` | Allemand |
| `en` | Anglais |
| `es` | Espagnol |
| `fr` | Francais |
| `it` | Italien |
| `zh` | Chinois |

## Noms de pays

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

// En utilisant un objet Locale
final name = country?.displayNameFromLocale(Locale('de'));
print(name); // Vereinigte Staaten

// En utilisant BuildContext (dans un widget)
final contextName = country?.displayName(context);
```

```typescript
import { getCountriesName } from 'infobits-intl';

const name = getCountriesName('US', 'de');
console.log(name); // Vereinigte Staaten

const nameFr = getCountriesName('FR', 'es');
console.log(nameFr); // Francia
```

## Noms de langues

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

## Noms de devises

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

## Noms de continents

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

## Noms de capitales

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

## Acces direct aux maps de traductions

Pour les operations en masse, vous pouvez acceder directement aux maps de traductions completes.

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Acceder a toutes les traductions allemandes des pays
germanCountries := i18n.CountriesTranslations["de"]
for code, name := range germanCountries {
    fmt.Printf("%s: %s\n", code, name)
}

// Autres maps disponibles :
// i18n.LanguagesTranslations
// i18n.CurrenciesTranslations
// i18n.ContinentsTranslations
// i18n.CapitalsTranslations
```

```dart
// Les maps de traductions sont accessibles via les classes de delegues
final germanCountries = CountriesTranslationsDelegate.translations['de'];
germanCountries?.forEach((code, name) {
  print('$code: $name');
});

// Autres delegues disponibles :
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

// Acceder a toutes les traductions allemandes des pays
const germanCountries = countriesTranslations['de'];
for (const [code, name] of Object.entries(germanCountries)) {
  console.log(`${code}: ${name}`);
}
```
