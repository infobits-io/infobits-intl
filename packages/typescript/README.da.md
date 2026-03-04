# @infobits/intl

Internationaliseringsdata for lande, sprog, valutaer og kontinenter med oversættelser og SVG-flag.

## Installation

```bash
npm install @infobits/intl
```

## Brug

### Lande

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Hent et land efter kode
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Hent alle lande
console.log(countries.length); // 248

// Filtrer lande
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Sprog

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Hent et sprog efter kode
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Hent alle sprog
console.log(languages.length); // 185
```

### Valutaer

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Hent en valuta efter kode
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Hent alle valutaer
console.log(currencies.length); // 179
```

### Kontinenter

```typescript
import { continents } from '@infobits/intl';

// Hent alle kontinenter
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Flag

```typescript
import { flags } from '@infobits/intl';

// Hent SVG-flagstreng
const usaFlag = flags['US'];

// Brug i HTML
document.getElementById('flag').innerHTML = usaFlag;

// Brug i React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Oversættelser

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Hent oversat landenavn
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Hent oversat sprognavn
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Hent oversat valutanavn
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Funktioner

- 248 lande med ISO 3166-1-koder
- 185 sprog med ISO 639-1-koder
- 179 valutaer med ISO 4217-koder
- 7 kontinenter
- SVG-landeflag (inline indlejret)
- Flersprogede oversættelser
- Fuld TypeScript-understøttelse

## Licens

MIT
