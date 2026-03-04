---
title: Pays
sidebar_position: 2
description: Travailler avec les donnees de pays -- recherche par code, lister tous les pays, filtrer par continent et acceder aux drapeaux SVG.
---

# Pays

Le package inclut 248+ pays avec les metadonnees completes ISO 3166-1.

## Proprietes d'un pays

Chaque pays possede les champs suivants :

| Champ | Description |
|-------|-------------|
| `alpha2Code` | Code ISO 3166-1 alpha-2 (ex. `"US"`) |
| `alpha3Code` | Code ISO 3166-1 alpha-3 (ex. `"USA"`) |
| `numericCode` | Code numerique ISO 3166-1 (ex. `840`) |
| `nativeName` | Nom natif du pays |
| `capital` | Capitale |
| `mainLanguage` | Code de la langue principale |
| `languages` | Liste des codes de langues parlees |
| `tld` | Domaine de premier niveau (ex. `".us"`) |
| `callingCode` | Indicatif telephonique international (ex. `1`) |
| `continent` | Identifiant du continent |
| `currency` | Code de la devise principale |

## Recherche par code

### Code alpha-2

```go
country, ok := intl.CountryByAlpha2("US")
if ok {
    fmt.Println(country.NativeName) // United States
    fmt.Println(country.Capital)    // Washington, D.C.
}
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.nativeName); // United States
print(country?.capital);    // Washington, D.C.
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('US');
console.log(country?.nativeName); // United States
console.log(country?.capital);    // Washington, D.C.
```

### Code alpha-3

```go
country, ok := intl.CountryByAlpha3("USA")
if ok {
    fmt.Println(country.Alpha2Code) // US
}
```

```dart
final country = Country.fromAlpha3Code('USA');
print(country?.alpha2Code); // US
```

```typescript
import { getCountryByAlpha3 } from 'infobits-intl';

const country = getCountryByAlpha3('USA');
console.log(country?.alpha2Code); // US
```

### Code numerique (Dart uniquement)

```dart
final country = Country.fromNumericCode(840);
print(country?.nativeName); // United States
```

## Lister tous les pays

```go
allCountries := intl.AllCountries()
fmt.Println(len(allCountries)) // 248
```

```dart
final allCountries = Country.values;
print(allCountries.length); // 253
```

```typescript
import { countries, CountryCode } from 'infobits-intl';

const allCountries = Object.values(countries);
console.log(allCountries.length);
```

## Filtrer par continent

```go
european := intl.CountriesByContinent("europe")
for _, c := range european {
    fmt.Println(c.NativeName)
}
```

```dart
final european = Continent.europe.countries;
for (final c in european) {
  print(c.nativeName);
}
```

```typescript
import { getCountriesByContinent } from 'infobits-intl';

const european = getCountriesByContinent('europe');
european.forEach(c => console.log(c.nativeName));
```

## Drapeaux SVG

Chaque pays dispose d'un drapeau SVG en ligne disponible sous forme de chaine de caracteres.

```go
svg, ok := intl.GetFlag("US")
if ok {
    // svg contient le balisage complet <svg>...</svg>
    fmt.Println(len(svg))
}

// Ou acceder directement a la map
svg = intl.Flags["US"]
```

```dart
// Acces via la map countryFlags
final svg = countryFlags['US'];

// Ou via la propriete Country
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;

// Afficher comme widget Flutter
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// En utilisant la fonction de recherche
const svg = getFlag('US');

// Ou acceder directement a la map
const svgDirect = flags['US'];
```

## Drapeaux emoji

```go
code := intl.CountryUS
fmt.Println(code.EmojiFlag())
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.emojiFlag);
```

```typescript
import { getEmojiFlag } from 'infobits-intl';

console.log(getEmojiFlag('US'));
```

## Metadonnees de pays

```go
country, _ := intl.CountryByAlpha2("JP")
fmt.Println(country.CallingCode) // 81
fmt.Println(country.TLD)         // .jp
fmt.Println(country.Currency)    // JPY
fmt.Println(country.Continent)   // asia
fmt.Println(country.MainLanguage) // ja
```

```dart
final country = Country.fromAlpha2Code('JP');
print(country?.callingCode);          // 81
print(country?.tld);                  // .jp
print(country?.currency.code);        // JPY
print(country?.continent.code);       // AS
print(country?.mainLanguage.code);    // ja
```

```typescript
import { getCountryByAlpha2 } from 'infobits-intl';

const country = getCountryByAlpha2('JP');
console.log(country?.callingCode); // 81
console.log(country?.tld);         // .jp
console.log(country?.currency);    // JPY
console.log(country?.continent);   // asia
console.log(country?.mainLanguage); // ja
```
