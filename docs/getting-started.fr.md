---
title: Premiers pas
sidebar_position: 1
description: Apercu de infobits-intl et instructions d'installation pour Go, Dart et TypeScript.
---

# Premiers pas

infobits-intl fournit des donnees d'internationalisation completes incluant les pays, les langues, les devises et les continents. Chaque entite est accompagnee de codes conformes aux normes ISO, de metadonnees, de drapeaux SVG de pays et de traductions en plusieurs langues.

## Contenu inclus

- **248+ pays** avec codes ISO 3166-1 alpha-2/alpha-3, capitales, indicatifs telephoniques, TLDs, et plus encore
- **185+ langues** avec codes ISO 639-1, noms natifs et dialectes
- **179+ devises** avec codes ISO 4217, symboles et formes plurielles
- **7 continents** avec regroupements par pays
- **Drapeaux SVG des pays** integres sous forme de chaines de caracteres
- **Traductions** pour toutes les entites dans 7 langues : danois, allemand, anglais, espagnol, francais, italien et chinois

## Installation

```go
go get github.com/infobits-io/infobits-intl-go
```

```dart
dart pub add infobits_intl
```

```typescript
npm install infobits-intl
```

## Exemple rapide

Rechercher un pays et afficher ses details :

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

## Constantes de code type-safe

Chaque package fournit des constantes typees pour les codes de pays, langues, devises et continents, permettant l'autocompletion et la verification a la compilation.

```go
// Utiliser des constantes de code pays typees
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart utilise des enums pour la securite de type
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
