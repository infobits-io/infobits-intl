# Infobits Intl for Dart/Flutter

Internationalization data for countries, languages, currencies, and continents with translations and SVG flags.

## Installation

```yaml
dependencies:
  infobits_intl: ^1.0.0
```

```bash
flutter pub get
```

## Usage

### Countries

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Get a country by code
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Get all countries
final allCountries = Countries.values;

// Get countries by continent
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Languages

```dart
// Get a language by code
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Get all languages
final allLanguages = Languages.values;
```

### Currencies

```dart
// Get a currency by code
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Get all currencies
final allCurrencies = Currencies.values;
```

### Continents

```dart
// Get all continents
final allContinents = Continents.values;

// Access continent properties
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Flags

```dart
import 'package:flutter_svg/flutter_svg.dart';

// Get SVG flag string
final flagSvg = countryFlags['US'];

// Display flag in Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Translations

```dart
// Get translated country name
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Get translated language name
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Get translated currency name
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Features

- 248 countries with ISO 3166-1 codes
- 185 languages with ISO 639-1 codes
- 179 currencies with ISO 4217 codes
- 7 continents
- SVG country flags (inline embedded)
- Multi-language translations

## License

MIT
