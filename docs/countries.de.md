---
title: Länder
sidebar_position: 2
description: Arbeiten mit Länderdaten — Suche nach Code, alle Länder auflisten, nach Kontinent filtern und auf SVG-Flaggen zugreifen.
---

# Länder

Das Paket enthält 248+ Länder mit vollständigen ISO 3166-1-Metadaten.

## Ländereigenschaften

Jedes Land hat die folgenden Felder:

| Feld | Beschreibung |
|------|--------------|
| `alpha2Code` | ISO 3166-1 Alpha-2-Code (z.B. `"US"`) |
| `alpha3Code` | ISO 3166-1 Alpha-3-Code (z.B. `"USA"`) |
| `numericCode` | ISO 3166-1 numerischer Code (z.B. `840`) |
| `nativeName` | Eigenname des Landes |
| `capital` | Hauptstadt |
| `mainLanguage` | Primärer Sprachcode |
| `languages` | Liste der gesprochenen Sprachcodes |
| `tld` | Top-Level-Domain (z.B. `".us"`) |
| `callingCode` | Internationale Vorwahl (z.B. `1`) |
| `continent` | Kontinent-Bezeichner |
| `currency` | Primärer Währungscode |

## Suche nach Code

### Alpha-2-Code

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

### Alpha-3-Code

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

### Numerischer Code (nur Dart)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Alle Länder auflisten

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

## Nach Kontinent filtern

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

## SVG-Flaggen

Für jedes Land ist eine eingebettete SVG-Flagge als String verfügbar.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg enthält das vollständige <svg>...</svg>-Markup
    fmt.Println(len(svg))
}

// Oder direkt auf die Map zugreifen
svg = intl.Flags["US"]
```

```dart
// Zugriff über die countryFlags-Map
final svg = countryFlags['US'];

// Oder über die Country-Eigenschaft
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Als Flutter-Widget rendern
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Mit der Lookup-Funktion
const svg = getFlag('US');

// Oder direkt auf die Map zugreifen
const svgDirect = flags['US'];
```

## Emoji-Flaggen

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

## Länder-Metadaten

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
