---
title: Kontinenter
sidebar_position: 5
description: Arbejde med kontinentdata -- opslag efter kode, liste over alle kontinenter, hent lande efter kontinent.
---

# Kontinenter

Pakken indeholder alle 7 kontinenter med standardiserede koder.

## Kontinentegenskaber

Hvert kontinent har foelgende felter:

| Felt | Beskrivelse |
|------|-------------|
| `code` | Tobogstavs kontinentkode |
| `name` | Engelsk visningsnavn |

## Tilgaengelige kontinenter

| Kode | Navn |
|------|------|
| `AF` | Afrika |
| `AQ` | Antarktis |
| `AS` | Asien |
| `EU` | Europa |
| `NA` | Nordamerika |
| `OS` | Oceanien |
| `SA` | Sydamerika |

## Opslag efter kode

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

## Liste over alle kontinenter

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

## Hent lande efter kontinent

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

## Kontinent fra et land

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // ikke garanteret; kontinent-feltet er ID-baseret
    }
    // Landets Continent-felt er et ID som "southAmerica"
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

## Typesikre kontinentkoder

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

## Sprog i et kontinent (Dart)

Dart-pakken giver ogs adgang til sprog, der tales i et kontinent.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
