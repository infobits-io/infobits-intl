# Infobits Intl til Dart/Flutter

Internationaliseringsdata for lande, sprog, valutaer og kontinenter med oversættelser og SVG-flag.

## Installation

```bash
flutter pub add infobits_intl
```

## Brug

### Lande

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Hent et land efter kode
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Hent alle lande
final allCountries = Countries.values;

// Hent lande efter kontinent
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Sprog

```dart
// Hent et sprog efter kode
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Hent alle sprog
final allLanguages = Languages.values;
```

### Valutaer

```dart
// Hent en valuta efter kode
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Hent alle valutaer
final allCurrencies = Currencies.values;
```

### Kontinenter

```dart
// Hent alle kontinenter
final allContinents = Continents.values;

// Tilgå kontinentegenskaber
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Flag

```dart
import 'package:flutter_svg/flutter_svg.dart';

// Hent SVG-flagstreng
final flagSvg = countryFlags['US'];

// Vis flag i Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Oversættelser

```dart
// Hent oversat landenavn
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Hent oversat sprognavn
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Hent oversat valutanavn
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Funktioner

- 248 lande med ISO 3166-1-koder
- 185 sprog med ISO 639-1-koder
- 179 valutaer med ISO 4217-koder
- 7 kontinenter
- SVG-landeflag (indlejret inline)
- Flersprogede oversættelser

## Licens

MIT
