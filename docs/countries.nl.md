---
title: Landen
sidebar_position: 2
description: Werken met landgegevens — opzoeken op code, alle landen weergeven, filteren op continent en SVG-vlaggen gebruiken.
---

# Landen

Het pakket bevat 248+ landen met volledige ISO 3166-1-metadata.

## Landeigenschappen

Elk land heeft de volgende velden:

| Veld | Beschrijving |
|------|--------------|
| `alpha2Code` | ISO 3166-1 alfa-2-code (bijv. `"US"`) |
| `alpha3Code` | ISO 3166-1 alfa-3-code (bijv. `"USA"`) |
| `numericCode` | ISO 3166-1 numerieke code (bijv. `840`) |
| `nativeName` | Oorspronkelijke landnaam |
| `capital` | Hoofdstad |
| `mainLanguage` | Primaire taalcode |
| `languages` | Lijst van gesproken taalcodes |
| `tld` | Topniveaudomein (bijv. `".us"`) |
| `callingCode` | Internationaal landnummer (bijv. `1`) |
| `continent` | Continentidentificatie |
| `currency` | Primaire valutacode |

## Opzoeken op code

### Alfa-2-code

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

### Alfa-3-code

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

### Numerieke code (alleen Dart)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Alle landen weergeven

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

## Filteren op continent

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

## SVG-vlaggen

Elk land heeft een inline SVG-vlag beschikbaar als string.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg bevat de volledige <svg>...</svg>-opmaak
    fmt.Println(len(svg))
}

// Of benader de map rechtstreeks
svg = intl.Flags["US"]
```

```dart
// Toegang via de countryFlags-map
final svg = countryFlags['US'];

// Of via de Country-eigenschap
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Weergeven als Flutter-widget
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Met de opzoekfunctie
const svg = getFlag('US');

// Of benader de map rechtstreeks
const svgDirect = flags['US'];
```

## Emojivlaggen

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

## Landmetadata

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
