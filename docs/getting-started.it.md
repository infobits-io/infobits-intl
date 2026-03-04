---
title: Per Iniziare
sidebar_position: 1
description: Panoramica di infobits-intl e istruzioni di installazione per Go, Dart e TypeScript.
---

# Per Iniziare

infobits-intl fornisce dati completi per l'internazionalizzazione, tra cui paesi, lingue, valute e continenti. Ogni entita include codici standard ISO, metadati, bandiere SVG dei paesi e traduzioni in piu lingue.

## Cosa Include

- **248+ paesi** con codici ISO 3166-1 alpha-2/alpha-3, capitali, prefissi telefonici, TLD e altro
- **185+ lingue** con codici ISO 639-1, nomi nativi e dialetti
- **179+ valute** con codici ISO 4217, simboli e forme plurali
- **7 continenti** con raggruppamenti per paese
- **Bandiere SVG dei paesi** incorporate come stringhe
- **Traduzioni** per tutte le entita in 7 lingue: danese, tedesco, inglese, spagnolo, francese, italiano e cinese

## Installazione

```go
go get github.com/infobits-io/infobits-intl-go
```

```dart
dart pub add infobits_intl
```

```typescript
npm install infobits-intl
```

## Esempio Rapido

Cerca un paese e stampa i suoi dettagli:

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

## Costanti di Codice Type-Safe

Ogni pacchetto fornisce costanti tipizzate per codici di paesi, lingue, valute e continenti, abilitando l'autocompletamento e la sicurezza a tempo di compilazione.

```go
// Usa costanti tipizzate per i codici paese
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart usa enum per la sicurezza dei tipi
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
