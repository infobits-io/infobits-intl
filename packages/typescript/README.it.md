# @infobits/intl

Dati di internazionalizzazione per paesi, lingue, valute e continenti con traduzioni e bandiere SVG.

## Installazione

```bash
npm install @infobits/intl
```

## Utilizzo

### Paesi

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Ottenere un paese per codice
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Ottenere tutti i paesi
console.log(countries.length); // 248

// Filtrare i paesi
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Lingue

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Ottenere una lingua per codice
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Ottenere tutte le lingue
console.log(languages.length); // 185
```

### Valute

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Ottenere una valuta per codice
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Ottenere tutte le valute
console.log(currencies.length); // 179
```

### Continenti

```typescript
import { continents } from '@infobits/intl';

// Ottenere tutti i continenti
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Bandiere

```typescript
import { flags } from '@infobits/intl';

// Ottenere la stringa SVG della bandiera
const usaFlag = flags['US'];

// Usare in HTML
document.getElementById('flag').innerHTML = usaFlag;

// Usare in React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Traduzioni

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Ottenere il nome del paese tradotto
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Ottenere il nome della lingua tradotto
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Ottenere il nome della valuta tradotto
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Funzionalità

- 248 paesi con codici ISO 3166-1
- 185 lingue con codici ISO 639-1
- 179 valute con codici ISO 4217
- 7 continenti
- Bandiere SVG dei paesi (incorporate in linea)
- Traduzioni multilingue
- Supporto completo per TypeScript

## Licenza

MIT
