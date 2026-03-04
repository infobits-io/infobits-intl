---
title: Paises
sidebar_position: 2
description: Trabajar con datos de paises — busqueda por codigo, listar todos los paises, filtrar por continente y acceder a banderas SVG.
---

# Paises

El paquete incluye 248+ paises con metadatos completos de ISO 3166-1.

## Propiedades de pais

Cada pais tiene los siguientes campos:

| Campo | Descripcion |
|-------|-------------|
| `alpha2Code` | Codigo ISO 3166-1 alpha-2 (ej. `"US"`) |
| `alpha3Code` | Codigo ISO 3166-1 alpha-3 (ej. `"USA"`) |
| `numericCode` | Codigo numerico ISO 3166-1 (ej. `840`) |
| `nativeName` | Nombre nativo del pais |
| `capital` | Ciudad capital |
| `mainLanguage` | Codigo del idioma principal |
| `languages` | Lista de codigos de idiomas hablados |
| `tld` | Dominio de nivel superior (ej. `".us"`) |
| `callingCode` | Codigo de llamada internacional (ej. `1`) |
| `continent` | Identificador de continente |
| `currency` | Codigo de moneda principal |

## Busqueda por codigo

### Codigo Alpha-2

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

### Codigo Alpha-3

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

### Codigo numerico (solo Dart)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Listar todos los paises

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

## Filtrar por continente

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

## Banderas SVG

Cada pais tiene una bandera SVG en linea disponible como cadena de texto.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg contiene el marcado completo <svg>...</svg>
    fmt.Println(len(svg))
}

// O acceder al mapa directamente
svg = intl.Flags["US"]
```

```dart
// Acceder mediante el mapa countryFlags
final svg = countryFlags['US'];

// O mediante la propiedad del pais
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Renderizar como widget de Flutter
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Usando la funcion de busqueda
const svg = getFlag('US');

// O acceder al mapa directamente
const svgDirect = flags['US'];
```

## Banderas emoji

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

## Metadatos de pais

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
