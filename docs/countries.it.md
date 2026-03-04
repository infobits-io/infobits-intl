---
title: Paesi
sidebar_position: 2
description: Lavorare con i dati dei paesi — ricerca per codice, elenco di tutti i paesi, filtraggio per continente e accesso alle bandiere SVG.
---

# Paesi

Il pacchetto include oltre 248 paesi con metadati completi ISO 3166-1.

## Proprieta dei Paesi

Ogni paese ha i seguenti campi:

| Campo | Descrizione |
|-------|-------------|
| `alpha2Code` | Codice ISO 3166-1 alpha-2 (es. `"US"`) |
| `alpha3Code` | Codice ISO 3166-1 alpha-3 (es. `"USA"`) |
| `numericCode` | Codice numerico ISO 3166-1 (es. `840`) |
| `nativeName` | Nome nativo del paese |
| `capital` | Capitale |
| `mainLanguage` | Codice della lingua principale |
| `languages` | Elenco dei codici delle lingue parlate |
| `tld` | Dominio di primo livello (es. `".us"`) |
| `callingCode` | Prefisso telefonico internazionale (es. `1`) |
| `continent` | Identificatore del continente |
| `currency` | Codice della valuta principale |

## Ricerca per Codice

### Codice Alpha-2

```go
country, ok := intl.CountryByAlpha2("US")
if ok {
    fmt.Println(country.NativeName) // United States
    fmt.Println(country.Capital)    // Washington, D.C.
}
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.nativeName); // United States
print(country?.capital);    // Washington, D.C.
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('US');
console.log(country?.nativeName); // United States
console.log(country?.capital);    // Washington, D.C.
```

### Codice Alpha-3

```go
country, ok := intl.CountryByAlpha3("USA")
if ok {
    fmt.Println(country.Alpha2Code) // US
}
```

```dart
final country = Country.fromAlpha3Code('USA');
print(country?.alpha2Code); // US
```

```typescript
import { getCountryByAlpha3 } from 'infobits-intl';

const country = getCountryByAlpha3('USA');
console.log(country?.alpha2Code); // US
```

### Codice Numerico (solo Dart)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Elenco di Tutti i Paesi

```go
allCountries := intl.AllCountries()
fmt.Println(len(allCountries)) // 248
```

```dart
final allCountries = Country.values;
print(allCountries.length); // 253
```

```typescript
import { countries, CountryCode } from 'infobits-intl';

const allCountries = Object.values(countries);
console.log(allCountries.length);
```

## Filtrare per Continente

```go
european := intl.CountriesByContinent("europe")
for _, c := range european {
    fmt.Println(c.NativeName)
}
```

```dart
final european = Continent.europe.countries;
for (final c in european) {
  print(c.nativeName);
}
```

```typescript
import { getCountriesByContinent } from 'infobits-intl';

const european = getCountriesByContinent('europe');
european.forEach(c => console.log(c.nativeName));
```

## Bandiere SVG

Ogni paese ha una bandiera SVG inline disponibile come stringa.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg contiene il markup completo <svg>...</svg>
    fmt.Println(len(svg))
}

// Oppure accedi direttamente alla mappa
svg = intl.Flags["US"]
```

```dart
// Accesso tramite la mappa countryFlags
final svg = countryFlags['US'];

// Oppure tramite la proprieta Country
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Renderizza come widget Flutter
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Usando la funzione di ricerca
const svg = getFlag('US');

// Oppure accedi direttamente alla mappa
const svgDirect = flags['US'];
```

## Bandiere Emoji

```go
code := intl.CountryUS
fmt.Println(code.EmojiFlag())
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.emojiFlag);
```

```typescript
import { getEmojiFlag } from 'infobits-intl';

console.log(getEmojiFlag('US'));
```

## Metadati del Paese

```go
country, _ := intl.CountryByAlpha2("JP")
fmt.Println(country.CallingCode) // 81
fmt.Println(country.TLD)         // .jp
fmt.Println(country.Currency)    // JPY
fmt.Println(country.Continent)   // asia
fmt.Println(country.MainLanguage) // ja
```

```dart
final country = Country.fromAlpha2Code('JP');
print(country?.callingCode);          // 81
print(country?.tld);                  // .jp
print(country?.currency.code);        // JPY
print(country?.continent.code);       // AS
print(country?.mainLanguage.code);    // ja
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('JP');
console.log(country?.callingCode); // 81
console.log(country?.tld);         // .jp
console.log(country?.currency);    // JPY
console.log(country?.continent);   // asia
console.log(country?.mainLanguage); // ja
```
