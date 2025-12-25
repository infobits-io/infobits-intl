# @infobits/intl

Internationalization data for countries, languages, currencies, and continents with translations and SVG flags.

## Installation

```bash
npm install @infobits/intl
```

## Usage

### Countries

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Get a country by code
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Get all countries
console.log(countries.length); // 248

// Filter countries
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Languages

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Get a language by code
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Get all languages
console.log(languages.length); // 185
```

### Currencies

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Get a currency by code
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Get all currencies
console.log(currencies.length); // 179
```

### Continents

```typescript
import { continents } from '@infobits/intl';

// Get all continents
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Flags

```typescript
import { flags } from '@infobits/intl';

// Get SVG flag string
const usaFlag = flags['US'];

// Use in HTML
document.getElementById('flag').innerHTML = usaFlag;

// Use in React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Translations

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Get translated country name
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Get translated language name
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Get translated currency name
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Features

- 248 countries with ISO 3166-1 codes
- 185 languages with ISO 639-1 codes
- 179 currencies with ISO 4217 codes
- 7 continents
- SVG country flags (inline embedded)
- Multi-language translations
- Full TypeScript support

## License

MIT
