---
title: Continenten
sidebar_position: 5
description: Werken met continentgegevens — opzoeken op code, alle continenten weergeven, landen per continent ophalen.
---

# Continenten

Het pakket bevat alle 7 continenten met gestandaardiseerde codes.

## Continenteigenschappen

Elk continent heeft de volgende velden:

| Veld | Beschrijving |
|------|--------------|
| `code` | Tweeletterige continentcode |
| `name` | Engelse weergavenaam |

## Beschikbare continenten

| Code | Naam |
|------|------|
| `AF` | Afrika |
| `AQ` | Antarctica |
| `AS` | Azie |
| `EU` | Europa |
| `NA` | Noord-Amerika |
| `OS` | Oceanie |
| `SA` | Zuid-Amerika |

## Opzoeken op code

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

## Alle continenten weergeven

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

## Landen per continent ophalen

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

## Continent van een land

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // niet gegarandeerd; continentveld is ID-gebaseerd
    }
    // Het Continent-veld van het land is een ID zoals "southAmerica"
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

## Typeveilige continentcodes

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

## Talen op een continent (Dart)

Het Dart-pakket biedt ook toegang tot talen die op een continent gesproken worden.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
