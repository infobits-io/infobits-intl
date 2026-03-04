---
title: Countries
sidebar_position: 2
description: Working with country data — lookup by code, list all countries, filter by continent, and access SVG flags.
---

# Countries

The package includes 248+ countries with full ISO 3166-1 metadata.

## Country Properties

Each country has the following fields:

| Field | Description |
|-------|-------------|
| `alpha2Code` | ISO 3166-1 alpha-2 code (e.g. `"US"`) |
| `alpha3Code` | ISO 3166-1 alpha-3 code (e.g. `"USA"`) |
| `numericCode` | ISO 3166-1 numeric code (e.g. `840`) |
| `nativeName` | Native country name |
| `capital` | Capital city |
| `mainLanguage` | Primary language code |
| `languages` | List of spoken language codes |
| `tld` | Top-level domain (e.g. `".us"`) |
| `callingCode` | International calling code (e.g. `1`) |
| `continent` | Continent identifier |
| `currency` | Primary currency code |

## Lookup by Code

### Alpha-2 Code

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

### Alpha-3 Code

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

### Numeric Code (Dart only)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## List All Countries

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

## Filter by Continent

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

## SVG Flags

Each country has an inline SVG flag available as a string.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg contains the full <svg>...</svg> markup
    fmt.Println(len(svg))
}

// Or access the map directly
svg = intl.Flags["US"]
```

```dart
// Access via the countryFlags map
final svg = countryFlags['US'];

// Or via the Country property
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Render as a Flutter widget
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Using the lookup function
const svg = getFlag('US');

// Or access the map directly
const svgDirect = flags['US'];
```

## Emoji Flags

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

## Country Metadata

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
