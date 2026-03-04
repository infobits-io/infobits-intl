---
title: Traduzioni
sidebar_position: 6
description: Come funzionano le traduzioni i18n — ottenere nomi localizzati per paesi, lingue, valute, continenti e capitali.
---

# Traduzioni

Tutti i nomi delle entita (paesi, lingue, valute, continenti e capitali) possono essere recuperati in piu lingue.

## Lingue Disponibili

| Codice | Lingua |
|--------|--------|
| `da` | Danese |
| `de` | Tedesco |
| `en` | Inglese |
| `es` | Spagnolo |
| `fr` | Francese |
| `it` | Italiano |
| `zh` | Cinese |

## Nomi dei Paesi

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

// Usando un oggetto Locale
final name = country?.displayNameFromLocale(Locale('de'));
print(name); // Vereinigte Staaten

// Usando BuildContext (in un widget)
final contextName = country?.displayName(context);
```

```typescript
import { getCountriesName } from 'infobits-intl';

const name = getCountriesName('US', 'de');
console.log(name); // Vereinigte Staaten

const nameFr = getCountriesName('FR', 'es');
console.log(nameFr); // Francia
```

## Nomi delle Lingue

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

## Nomi delle Valute

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

## Nomi dei Continenti

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

## Nomi delle Capitali

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

## Accesso Diretto alle Mappe di Traduzione

Per operazioni in blocco, puoi accedere direttamente alle mappe di traduzione complete.

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Accedi a tutte le traduzioni dei paesi in tedesco
germanCountries := i18n.CountriesTranslations["de"]
for code, name := range germanCountries {
    fmt.Printf("%s: %s\n", code, name)
}

// Altre mappe disponibili:
// i18n.LanguagesTranslations
// i18n.CurrenciesTranslations
// i18n.ContinentsTranslations
// i18n.CapitalsTranslations
```

```dart
// Le mappe di traduzione sono accessibili tramite classi delegate
final germanCountries = CountriesTranslationsDelegate.translations['de'];
germanCountries?.forEach((code, name) {
  print('$code: $name');
});

// Altre classi delegate disponibili:
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

// Accedi a tutte le traduzioni dei paesi in tedesco
const germanCountries = countriesTranslations['de'];
for (const [code, name] of Object.entries(germanCountries)) {
  console.log(`${code}: ${name}`);
}
```
