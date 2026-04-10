# Infobits Intl voor Dart/Flutter

Internationaliseringsgegevens voor landen, talen, valuta's en continenten met vertalingen en SVG-vlaggen.

## Installatie

```bash
flutter pub add infobits_intl
```

## Gebruik

### Landen

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Een land ophalen op basis van code
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Alle landen ophalen
final allCountries = Countries.values;

// Landen ophalen per continent
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Talen

```dart
// Een taal ophalen op basis van code
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Alle talen ophalen
final allLanguages = Languages.values;
```

### Valuta's

```dart
// Een valuta ophalen op basis van code
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Alle valuta's ophalen
final allCurrencies = Currencies.values;
```

### Continenten

```dart
// Alle continenten ophalen
final allContinents = Continents.values;

// Continenteigenschappen opvragen
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Vlaggen

```dart
import 'package:flutter_svg/flutter_svg.dart';

// SVG-vlagtekenreeks ophalen
final flagSvg = countryFlags['US'];

// Vlag weergeven in Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Vertalingen

```dart
// Vertaalde landnaam ophalen
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Vertaalde taalnaam ophalen
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Vertaalde valutanaam ophalen
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Kenmerken

- 248 landen met ISO 3166-1-codes
- 185 talen met ISO 639-1-codes
- 179 valuta's met ISO 4217-codes
- 7 continenten
- SVG-landvlaggen (inline ingebed)
- Meertalige vertalingen

## Licentie

MIT
