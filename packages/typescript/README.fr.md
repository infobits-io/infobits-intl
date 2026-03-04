# @infobits/intl

Données d'internationalisation pour les pays, langues, devises et continents avec traductions et drapeaux SVG.

## Installation

```bash
npm install @infobits/intl
```

## Utilisation

### Pays

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Obtenir un pays par code
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Obtenir tous les pays
console.log(countries.length); // 248

// Filtrer les pays
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Langues

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Obtenir une langue par code
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Obtenir toutes les langues
console.log(languages.length); // 185
```

### Devises

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Obtenir une devise par code
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Obtenir toutes les devises
console.log(currencies.length); // 179
```

### Continents

```typescript
import { continents } from '@infobits/intl';

// Obtenir tous les continents
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Drapeaux

```typescript
import { flags } from '@infobits/intl';

// Obtenir la chaîne SVG du drapeau
const usaFlag = flags['US'];

// Utiliser en HTML
document.getElementById('flag').innerHTML = usaFlag;

// Utiliser en React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Traductions

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Obtenir le nom de pays traduit
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Obtenir le nom de langue traduit
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Obtenir le nom de devise traduit
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Fonctionnalités

- 248 pays avec codes ISO 3166-1
- 185 langues avec codes ISO 639-1
- 179 devises avec codes ISO 4217
- 7 continents
- Drapeaux SVG des pays (intégrés en ligne)
- Traductions multilingues
- Support complet de TypeScript

## Licence

MIT
