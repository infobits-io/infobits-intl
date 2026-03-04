---
title: Continentes
sidebar_position: 5
description: Trabajar con datos de continentes — busqueda por codigo, listar todos los continentes, obtener paises por continente.
---

# Continentes

El paquete incluye los 7 continentes con codigos estandarizados.

## Propiedades de continente

Cada continente tiene los siguientes campos:

| Campo | Descripcion |
|-------|-------------|
| `code` | Codigo de continente de dos letras |
| `name` | Nombre para mostrar en ingles |

## Continentes disponibles

| Codigo | Nombre |
|--------|--------|
| `AF` | Africa |
| `AQ` | Antartida |
| `AS` | Asia |
| `EU` | Europa |
| `NA` | America del Norte |
| `OS` | Oceania |
| `SA` | America del Sur |

## Busqueda por codigo

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

## Listar todos los continentes

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

## Obtener paises por continente

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

## Continente de un pais

```go
country, ok := intl.CountryByAlpha2("BR")
if ok {
    continent, cOk := intl.ContinentByCode(country.Continent)
    if cOk {
        fmt.Println(continent.Name) // no garantizado; el campo continent esta basado en ID
    }
    // El campo Continent del pais es un ID como "southAmerica"
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

## Codigos de continente con seguridad de tipos

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

## Idiomas en un continente (Dart)

El paquete de Dart tambien proporciona acceso a los idiomas hablados en un continente.

```dart
final langs = Continent.europe.languages;
for (final lang in langs) {
  print('${lang.code}: ${lang.nativeName}');
}
```
