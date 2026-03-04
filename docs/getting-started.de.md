---
title: Erste Schritte
sidebar_position: 1
description: Überblick über infobits-intl und Installationsanweisungen für Go, Dart und TypeScript.
---

# Erste Schritte

infobits-intl bietet umfassende Internationalisierungsdaten, darunter Länder, Sprachen, Währungen und Kontinente. Jede Entität enthält ISO-standardisierte Codes, Metadaten, SVG-Länderflaggen und Übersetzungen in mehreren Sprachen.

## Was enthalten ist

- **248+ Länder** mit ISO 3166-1 Alpha-2/Alpha-3-Codes, Hauptstädten, Vorwahlen, TLDs und mehr
- **185+ Sprachen** mit ISO 639-1-Codes, Eigennamen und Dialekten
- **179+ Währungen** mit ISO 4217-Codes, Symbolen und Pluralformen
- **7 Kontinente** mit Ländergruppierungen
- **SVG-Länderflaggen** als Strings eingebettet
- **Übersetzungen** für alle Entitäten in 7 Sprachen: Dänisch, Deutsch, Englisch, Spanisch, Französisch, Italienisch und Chinesisch

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

## Schnellbeispiel

Ein Land nachschlagen und seine Details ausgeben:

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

## Typsichere Code-Konstanten

Jedes Paket stellt typisierte Konstanten für Länder-, Sprach-, Währungs- und Kontinent-Codes bereit, die Autovervollständigung und Kompilierzeit-Sicherheit ermöglichen.

```go
// Typisierte Ländercode-Konstanten verwenden
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart verwendet Enums für Typsicherheit
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```
