# Infobits Intl pour Dart/Flutter

Données d'internationalisation pour les pays, langues, devises et continents avec traductions et drapeaux SVG.

## Installation

```yaml
dependencies:
  infobits_intl: ^1.0.0
```

```bash
flutter pub get
```

## Utilisation

### Pays

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Obtenir un pays par code
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Obtenir tous les pays
final allCountries = Countries.values;

// Obtenir les pays par continent
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Langues

```dart
// Obtenir une langue par code
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Obtenir toutes les langues
final allLanguages = Languages.values;
```

### Devises

```dart
// Obtenir une devise par code
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Obtenir toutes les devises
final allCurrencies = Currencies.values;
```

### Continents

```dart
// Obtenir tous les continents
final allContinents = Continents.values;

// Accéder aux propriétés du continent
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Drapeaux

```dart
import 'package:flutter_svg/flutter_svg.dart';

// Obtenir la chaîne SVG du drapeau
final flagSvg = countryFlags['US'];

// Afficher le drapeau dans Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Traductions

```dart
// Obtenir le nom traduit du pays
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Obtenir le nom traduit de la langue
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Obtenir le nom traduit de la devise
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Fonctionnalités

- 248 pays avec codes ISO 3166-1
- 185 langues avec codes ISO 639-1
- 179 devises avec codes ISO 4217
- 7 continents
- Drapeaux de pays en SVG (intégrés en ligne)
- Traductions multilingues

## Licence

MIT
