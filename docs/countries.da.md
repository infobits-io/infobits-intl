---
title: Lande
sidebar_position: 2
description: Arbejde med landedata -- opslag efter kode, liste over alle lande, filtrering efter kontinent og adgang til SVG-flag.
---

# Lande

Pakken indeholder 248+ lande med fuld ISO 3166-1-metadata.

## Landeegenskaber

Hvert land har foelgende felter:

| Felt | Beskrivelse |
|------|-------------|
| `alpha2Code` | ISO 3166-1 alpha-2-kode (f.eks. `"US"`) |
| `alpha3Code` | ISO 3166-1 alpha-3-kode (f.eks. `"USA"`) |
| `numericCode` | ISO 3166-1 numerisk kode (f.eks. `840`) |
| `nativeName` | Oprindeligt landenavn |
| `capital` | Hovedstad |
| `mainLanguage` | Primaer sprogkode |
| `languages` | Liste over talte sprogkoder |
| `tld` | Topdomaene (f.eks. `".us"`) |
| `callingCode` | Internationalt opkaldsnummer (f.eks. `1`) |
| `continent` | Kontinent-identifikator |
| `currency` | Primaer valutakode |

## Opslag efter kode

### Alpha-2-kode

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

### Alpha-3-kode

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

### Numerisk kode (kun Dart)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Liste over alle lande

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

## Filtrer efter kontinent

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

## SVG-flag

Hvert land har et inline SVG-flag tilgaengeligt som en streng.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg indeholder den fulde <svg>...</svg>-markup
    fmt.Println(len(svg))
}

// Eller tilg direkte via map'et
svg = intl.Flags["US"]
```

```dart
// Tilg via countryFlags-map'et
final svg = countryFlags['US'];

// Eller via Country-egenskaben
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Renderer som en Flutter-widget
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Brug opslagsfunktionen
const svg = getFlag('US');

// Eller tilg map'et direkte
const svgDirect = flags['US'];
```

## Emoji-flag

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

## Landemetadata

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
