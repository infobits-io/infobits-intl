# Infobits Intl per Dart/Flutter

Dati di internazionalizzazione per paesi, lingue, valute e continenti con traduzioni e bandiere SVG.

## Installazione

```yaml
dependencies:
  infobits_intl: ^1.0.0
```

```bash
flutter pub get
```

## Utilizzo

### Paesi

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Ottenere un paese tramite codice
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Ottenere tutti i paesi
final allCountries = Countries.values;

// Ottenere i paesi per continente
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Lingue

```dart
// Ottenere una lingua tramite codice
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Ottenere tutte le lingue
final allLanguages = Languages.values;
```

### Valute

```dart
// Ottenere una valuta tramite codice
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Ottenere tutte le valute
final allCurrencies = Currencies.values;
```

### Continenti

```dart
// Ottenere tutti i continenti
final allContinents = Continents.values;

// Accedere alle proprietà del continente
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Bandiere

```dart
import 'package:flutter_svg/flutter_svg.dart';

// Ottenere la stringa SVG della bandiera
final flagSvg = countryFlags['US'];

// Visualizzare la bandiera in Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Traduzioni

```dart
// Ottenere il nome tradotto del paese
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Ottenere il nome tradotto della lingua
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Ottenere il nome tradotto della valuta
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Funzionalità

- 248 paesi con codici ISO 3166-1
- 185 lingue con codici ISO 639-1
- 179 valute con codici ISO 4217
- 7 continenti
- Bandiere dei paesi in SVG (incorporate inline)
- Traduzioni multilingue

## Licenza

MIT
