# Infobits Intl für Dart/Flutter

Internationalisierungsdaten für Länder, Sprachen, Währungen und Kontinente mit Übersetzungen und SVG-Flaggen.

## Installation

```yaml
dependencies:
  infobits_intl: ^1.0.0
```

```bash
flutter pub get
```

## Verwendung

### Länder

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Ein Land nach Code abrufen
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Alle Länder abrufen
final allCountries = Countries.values;

// Länder nach Kontinent abrufen
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Sprachen

```dart
// Eine Sprache nach Code abrufen
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Alle Sprachen abrufen
final allLanguages = Languages.values;
```

### Währungen

```dart
// Eine Währung nach Code abrufen
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Alle Währungen abrufen
final allCurrencies = Currencies.values;
```

### Kontinente

```dart
// Alle Kontinente abrufen
final allContinents = Continents.values;

// Auf Kontinent-Eigenschaften zugreifen
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Flaggen

```dart
import 'package:flutter_svg/flutter_svg.dart';

// SVG-Flaggenzeichenkette abrufen
final flagSvg = countryFlags['US'];

// Flagge in Flutter anzeigen
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Übersetzungen

```dart
// Übersetzten Ländernamen abrufen
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Übersetzten Sprachnamen abrufen
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Übersetzten Währungsnamen abrufen
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Funktionen

- 248 Länder mit ISO 3166-1-Codes
- 185 Sprachen mit ISO 639-1-Codes
- 179 Währungen mit ISO 4217-Codes
- 7 Kontinente
- SVG-Länderflaggen (inline eingebettet)
- Mehrsprachige Übersetzungen

## Lizenz

MIT
