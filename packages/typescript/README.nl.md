# @infobits/intl

Internationalisatiegegevens voor landen, talen, valuta's en continenten met vertalingen en SVG-vlaggen.

## Installatie

```bash
npm install @infobits/intl
```

## Gebruik

### Landen

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Een land ophalen op code
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Alle landen ophalen
console.log(countries.length); // 248

// Landen filteren
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Talen

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Een taal ophalen op code
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Alle talen ophalen
console.log(languages.length); // 185
```

### Valuta's

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Een valuta ophalen op code
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Alle valuta's ophalen
console.log(currencies.length); // 179
```

### Continenten

```typescript
import { continents } from '@infobits/intl';

// Alle continenten ophalen
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Vlaggen

```typescript
import { flags } from '@infobits/intl';

// SVG-vlagreeks ophalen
const usaFlag = flags['US'];

// Gebruiken in HTML
document.getElementById('flag').innerHTML = usaFlag;

// Gebruiken in React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Vertalingen

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Vertaalde landnaam ophalen
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Vertaalde taalnaam ophalen
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Vertaalde valutanaam ophalen
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Functies

- 248 landen met ISO 3166-1-codes
- 185 talen met ISO 639-1-codes
- 179 valuta's met ISO 4217-codes
- 7 continenten
- SVG-landvlaggen (inline ingebed)
- Meertalige vertalingen
- Volledige TypeScript-ondersteuning

## Licentie

MIT
