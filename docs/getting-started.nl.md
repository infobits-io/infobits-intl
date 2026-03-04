---
title: Aan de slag
sidebar_position: 1
description: Overzicht van infobits-intl en installatie-instructies voor Go, Dart en TypeScript.
---

# Aan de slag

infobits-intl biedt uitgebreide internationalisatiegegevens, waaronder landen, talen, valuta's en continenten. Elke entiteit bevat ISO-standaardcodes, metadata, SVG-landvlaggen en vertalingen in meerdere talen.

## Wat is inbegrepen

- **248+ landen** met ISO 3166-1 alfa-2/alfa-3-codes, hoofdsteden, landnummers, TLD's en meer
- **185+ talen** met ISO 639-1-codes, oorspronkelijke namen en dialecten
- **179+ valuta's** met ISO 4217-codes, symbolen en meervoudsvormen
- **7 continenten** met landengroeperingen
- **SVG-landvlaggen** ingebed als strings
- **Vertalingen** voor alle entiteiten in 7 taalgebieden: Deens, Duits, Engels, Spaans, Frans, Italiaans en Chinees

## Installatie

```go
go get github.com/infobits-io/infobits-intl-go
```

```dart
dart pub add infobits_intl
```

```typescript
npm install infobits-intl
```

## Snel voorbeeld

Zoek een land op en toon de details:

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

## Typeveilige codeconstanten

Elk pakket biedt getypeerde constanten voor land-, taal-, valuta- en continentcodes, waardoor automatisch aanvullen en controle tijdens het compileren mogelijk is.

```go
// Gebruik getypeerde landcodeconstanten
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart gebruikt enums voor typeveiligheid
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
