# @infobits/intl

Internationalisierungsdaten für Länder, Sprachen, Währungen und Kontinente mit Übersetzungen und SVG-Flaggen.

## Installation

```bash
npm install @infobits/intl
```

## Verwendung

### Länder

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Ein Land nach Code abrufen
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Alle Länder abrufen
console.log(countries.length); // 248

// Länder filtern
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Sprachen

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Eine Sprache nach Code abrufen
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Alle Sprachen abrufen
console.log(languages.length); // 185
```

### Währungen

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Eine Währung nach Code abrufen
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Alle Währungen abrufen
console.log(currencies.length); // 179
```

### Kontinente

```typescript
import { continents } from '@infobits/intl';

// Alle Kontinente abrufen
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Flaggen

```typescript
import { flags } from '@infobits/intl';

// SVG-Flaggenzeichenkette abrufen
const usaFlag = flags['US'];

// In HTML verwenden
document.getElementById('flag').innerHTML = usaFlag;

// In React verwenden
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Übersetzungen

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Übersetzten Ländernamen abrufen
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Übersetzten Sprachnamen abrufen
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Übersetzten Währungsnamen abrufen
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Funktionen

- 248 Länder mit ISO 3166-1-Codes
- 185 Sprachen mit ISO 639-1-Codes
- 179 Währungen mit ISO 4217-Codes
- 7 Kontinente
- SVG-Länderflaggen (inline eingebettet)
- Mehrsprachige Übersetzungen
- Vollständige TypeScript-Unterstützung

## Lizenz

MIT
