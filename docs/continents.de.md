---
title: Kontinente
sidebar_position: 5
description: Arbeiten mit Kontinentdaten — Suche nach Code, alle Kontinente auflisten, Länder nach Kontinent abrufen.
---

# Kontinente

Das Paket enthält alle 7 Kontinente mit standardisierten Codes.

## Kontinenteigenschaften

Jeder Kontinent hat die folgenden Felder:

| Feld | Beschreibung |
|------|--------------|
| `code` | Zweibuchstabiger Kontinent-Code |
| `name` | Englischer Anzeigename |

## Verfügbare Kontinente

| Code | Name |
|------|------|
| `AF` | Africa |
| `AQ` | Antarctica |
| `AS` | Asia |
| `EU` | Europe |
| `NA` | North America |
| `OS` | Oceania |
| `SA` | South America |

## Suche nach Code

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

## Alle Kontinente auflisten

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

## Länder nach Kontinent abrufen

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

## Kontinent eines Landes

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // nicht garantiert; Kontinent-Feld ist ID-basiert
    }
    // Das Continent-Feld des Landes ist eine ID wie "southAmerica"
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

## Typsichere Kontinent-Codes

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

## Sprachen in einem Kontinent (Dart)

Das Dart-Paket bietet auch Zugriff auf die in einem Kontinent gesprochenen Sprachen.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
