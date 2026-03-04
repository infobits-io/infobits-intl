# Infobits Intl para Dart/Flutter

Datos de internacionalización para países, idiomas, monedas y continentes con traducciones y banderas SVG.

## Instalación

```yaml
dependencies:
  infobits_intl: ^1.0.0
```

```bash
flutter pub get
```

## Uso

### Países

```dart
import 'package:infobits_intl/infobits_intl.dart';

// Obtener un país por código
final usa = Countries.byAlpha2Code('US');
print(usa?.name); // United States
print(usa?.alpha3Code); // USA
print(usa?.capital); // Washington, D.C.
print(usa?.callingCode); // 1

// Obtener todos los países
final allCountries = Countries.values;

// Obtener países por continente
final europeanCountries = Countries.values
    .where((c) => c.continent == Continent.europe);
```

### Idiomas

```dart
// Obtener un idioma por código
final english = Languages.byCode('en');
print(english?.name); // English
print(english?.nativeName); // English

// Obtener todos los idiomas
final allLanguages = Languages.values;
```

### Monedas

```dart
// Obtener una moneda por código
final usd = Currencies.byCode('USD');
print(usd?.name); // US Dollar
print(usd?.symbol); // $

// Obtener todas las monedas
final allCurrencies = Currencies.values;
```

### Continentes

```dart
// Obtener todos los continentes
final allContinents = Continents.values;

// Acceder a las propiedades del continente
print(Continent.europe.name); // Europe
print(Continent.europe.code); // EU
```

### Banderas

```dart
import 'package:flutter_svg/flutter_svg.dart';

// Obtener la cadena SVG de la bandera
final flagSvg = countryFlags['US'];

// Mostrar la bandera en Flutter
SvgPicture.string(
  countryFlags['US']!,
  width: 32,
  height: 24,
);
```

### Traducciones

```dart
// Obtener el nombre traducido del país
final countryName = CountryTranslations.getName('US', 'de'); // Vereinigte Staaten

// Obtener el nombre traducido del idioma
final languageName = LanguageTranslations.getName('en', 'es'); // Inglés

// Obtener el nombre traducido de la moneda
final currencyName = CurrencyTranslations.getName('USD', 'fr'); // Dollar américain
```

## Características

- 248 países con códigos ISO 3166-1
- 185 idiomas con códigos ISO 639-1
- 179 monedas con códigos ISO 4217
- 7 continentes
- Banderas de países en SVG (incrustadas en línea)
- Traducciones multilingües

## Licencia

MIT
