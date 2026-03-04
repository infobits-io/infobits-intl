---
title: Continenti
sidebar_position: 5
description: Lavorare con i dati dei continenti — ricerca per codice, elenco di tutti i continenti, ottenere i paesi per continente.
---

# Continenti

Il pacchetto include tutti i 7 continenti con codici standardizzati.

## Proprieta dei Continenti

Ogni continente ha i seguenti campi:

| Campo | Descrizione |
|-------|-------------|
| `code` | Codice continente a due lettere |
| `name` | Nome visualizzato in inglese |

## Continenti Disponibili

| Codice | Nome |
|--------|------|
| `AF` | Africa |
| `AQ` | Antartide |
| `AS` | Asia |
| `EU` | Europa |
| `NA` | Nord America |
| `OS` | Oceania |
| `SA` | Sud America |

## Ricerca per Codice

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

## Elenco di Tutti i Continenti

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

## Ottenere i Paesi per Continente

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

## Continente di un Paese

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // not guaranteed; continent field is ID-based
    }
    // Il campo Continent del paese e un ID come "southAmerica"
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

## Codici Continente Type-Safe

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

## Lingue in un Continente (Dart)

Il pacchetto Dart fornisce anche l'accesso alle lingue parlate in un continente.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
