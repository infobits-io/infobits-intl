---
title: Getting Started
sidebar_position: 1
description: Overview of infobits-intl and installation instructions for Go, Dart, TypeScript, and PHP.
---

# Getting Started

infobits-intl provides comprehensive internationalization data including countries, languages, currencies, and continents. Each entity comes with ISO-standard codes, metadata, SVG country flags, and translations in multiple languages.

## What's Included

- **248+ countries** with ISO 3166-1 alpha-2/alpha-3 codes, capitals, calling codes, TLDs, and more
- **185+ languages** with ISO 639-1 codes, native names, and dialects
- **179+ currencies** with ISO 4217 codes, symbols, and plural forms
- **7 continents** with country groupings
- **SVG country flags** embedded as strings
- **Translations** for all entities in 7 locales: Danish, German, English, Spanish, French, Italian, and Chinese

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

```php
composer require infobits/intl
```

## Quick Example

Look up a country and print its details:

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

```php
use Infobits\Intl\Country;

$country = Country::tryFrom('US');
if ($country !== null) {
    echo $country->nativeName();  // United States
    echo $country->alpha3Code();  // USA
    echo $country->capital();     // Washington, D.C.
}
```

## Type-Safe Code Constants

Each package provides typed constants for country, language, currency, and continent codes, enabling autocompletion and compile-time safety.

```go
// Use typed country code constants
code := intl.CountryUS
fmt.Println(code.Country().NativeName) // United States
fmt.Println(code.EmojiFlag())          // flag emoji
```

```dart
// Dart uses enums for type safety
print(Country.unitedStates.nativeName); // United States
print(Country.unitedStates.emojiFlag);  // flag emoji
```

```typescript
import { CountryCode, countries } from 'infobits-intl';

const usa = countries[CountryCode.US];
console.log(usa.nativeName); // United States
```

```php
use Infobits\Intl\Country;

// PHP backed enums provide type safety
echo Country::US->nativeName(); // United States
echo Country::US->emojiFlag();  // flag emoji
```
