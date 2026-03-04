---
title: Continents
sidebar_position: 5
description: Working with continent data — lookup by code, list all continents, get countries by continent.
---

# Continents

The package includes all 7 continents with standardized codes.

## Continent Properties

Each continent has the following fields:

| Field | Description |
|-------|-------------|
| `code` | Two-letter continent code |
| `name` | English display name |

## Available Continents

| Code | Name |
|------|------|
| `AF` | Africa |
| `AQ` | Antarctica |
| `AS` | Asia |
| `EU` | Europe |
| `NA` | North America |
| `OS` | Oceania |
| `SA` | South America |

## Lookup by Code

```go
continent, ok := intl.ContinentByCode("EU")
if ok {
    fmt.Println(continent.Name) // Europe
    fmt.Println(continent.Code) // EU
}
```

```dart
final continent = Continent.fromCode('EU');
print(continent?.name); // Europe
print(continent?.code); // EU
```

```typescript
import { getContinentByCode } from 'infobits-intl';

const continent = getContinentByCode('EU');
console.log(continent?.name); // Europe
console.log(continent?.code); // EU
```

## List All Continents

```go
allContinents := intl.AllContinents()
for _, c := range allContinents {
    fmt.Printf("%s (%s)\n", c.Name, c.Code)
}
```

```dart
for (final c in Continent.values) {
  print('${c.name} (${c.code})');
}
```

```typescript
import { continents } from 'infobits-intl';

Object.values(continents).forEach(c => {
  console.log(`${c.name} (${c.code})`);
});
```

## Get Countries by Continent

```go
countries := intl.CountriesByContinent("europe")
fmt.Println(len(countries))
for _, c := range countries {
    fmt.Println(c.NativeName)
}
```

```dart
final countries = Continent.europe.countries;
print(countries.length);
for (final c in countries) {
  print(c.nativeName);
}
```

```typescript
import { getCountriesByContinent } from 'infobits-intl';

const countries = getCountriesByContinent('europe');
console.log(countries.length);
countries.forEach(c => console.log(c.nativeName));
```

## Continent from a Country

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // not guaranteed; continent field is ID-based
    }
    // The country's Continent field is an ID like "southAmerica"
    fmt.Println(country.Continent)
}
```

```dart
final country = Country.fromAlpha2Code('BR');
print(country?.continent.name); // South America
print(country?.continent.code); // SA
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('BR');
console.log(country?.continent); // southAmerica
```

## Type-Safe Continent Codes

```go
code := intl.ContinentEU
continent := code.Continent()
fmt.Println(continent.Name) // Europe
```

```dart
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

```typescript
import { ContinentCode, continents } from 'infobits-intl';

const europe = continents[ContinentCode.EU];
console.log(europe.name); // Europe
```

## Languages in a Continent (Dart)

The Dart package also provides access to languages spoken in a continent.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
