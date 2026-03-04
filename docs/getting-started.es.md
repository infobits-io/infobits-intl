---
title: Primeros pasos
sidebar_position: 1
description: Descripcion general de infobits-intl e instrucciones de instalacion para Go, Dart y TypeScript.
---

# Primeros pasos

infobits-intl proporciona datos completos de internacionalizacion, incluyendo paises, idiomas, monedas y continentes. Cada entidad incluye codigos estandar ISO, metadatos, banderas SVG de paises y traducciones en multiples idiomas.

## Que incluye

- **248+ paises** con codigos ISO 3166-1 alpha-2/alpha-3, capitales, codigos de llamada, TLDs y mas
- **185+ idiomas** con codigos ISO 639-1, nombres nativos y dialectos
- **179+ monedas** con codigos ISO 4217, simbolos y formas plurales
- **7 continentes** con agrupaciones de paises
- **Banderas SVG de paises** integradas como cadenas de texto
- **Traducciones** para todas las entidades en 7 idiomas: danes, aleman, ingles, espanol, frances, italiano y chino

## Instalacion

```go
go get github.com/infobits-io/infobits-intl-go
```

```dart
dart pub add infobits_intl
```

```typescript
npm install infobits-intl
```

## Ejemplo rapido

Buscar un pais e imprimir sus detalles:

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    country, ok := intl.CountryByAlpha2("US")
    if ok {
        fmt.Println(country.NativeName)  // United States
        fmt.Println(country.Alpha3Code)  // USA
        fmt.Println(country.Capital)     // Washington, D.C.
    }
}
```

```dart
import 'package:infobits_intl/infobits_intl.dart';

void main() {
  final country = Country.fromAlpha2Code('US');
  if (country != null) {
    print(country.nativeName);  // United States
    print(country.alpha3Code);  // USA
    print(country.capital);     // Washington, D.C.
  }
}
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('US');
if (country) {
  console.log(country.nativeName);  // United States
  console.log(country.alpha3Code);  // USA
  console.log(country.capital);     // Washington, D.C.
}
```

## Constantes de codigo con seguridad de tipos

Cada paquete proporciona constantes tipadas para codigos de paises, idiomas, monedas y continentes, permitiendo autocompletado y seguridad en tiempo de compilacion.

```go
// Usar constantes tipadas de codigos de pais
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart usa enums para seguridad de tipos
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
