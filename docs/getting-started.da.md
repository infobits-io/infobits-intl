---
title: Kom godt i gang
sidebar_position: 1
description: Oversigt over infobits-intl og installationsvejledning til Go, Dart og TypeScript.
---

# Kom godt i gang

infobits-intl leverer omfattende internationaliseringsdata, herunder lande, sprog, valutaer og kontinenter. Hver enhed kommer med ISO-standardkoder, metadata, SVG-landeflag og oversaettelser til flere sprog.

## Hvad er inkluderet

- **248+ lande** med ISO 3166-1 alpha-2/alpha-3-koder, hovedstaeder, opkaldskoder, TLD'er og mere
- **185+ sprog** med ISO 639-1-koder, oprindelige navne og dialekter
- **179+ valutaer** med ISO 4217-koder, symboler og flertalsformer
- **7 kontinenter** med landegrupperinger
- **SVG-landeflag** indlejret som strenge
- **Oversaettelser** for alle enheder i 7 lokaliteter: dansk, tysk, engelsk, spansk, fransk, italiensk og kinesisk

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

## Hurtigt eksempel

Sl et land op og udskriv dets detaljer:

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

## Typesikre kodekonstanter

Hver pakke tilbyder typede konstanter for lande-, sprog-, valuta- og kontinentkoder, som muliggoer autofuldfoerelse og kompileringstidssikkerhed.

```go
// Brug typede landekodekonstanter
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart bruger enums til typesikkerhed
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
