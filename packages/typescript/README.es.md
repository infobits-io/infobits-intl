# @infobits/intl

Datos de internacionalización para países, idiomas, monedas y continentes con traducciones y banderas SVG.

## Instalación

```bash
npm install @infobits/intl
```

## Uso

### Países

```typescript
import { countries, getCountryByAlpha2Code } from '@infobits/intl';

// Obtener un país por código
const usa = getCountryByAlpha2Code('US');
console.log(usa?.name); // United States
console.log(usa?.alpha3Code); // USA
console.log(usa?.capital); // Washington, D.C.
console.log(usa?.callingCode); // 1

// Obtener todos los países
console.log(countries.length); // 248

// Filtrar países
const europeanCountries = countries.filter(c => c.continent === 'europe');
```

### Idiomas

```typescript
import { languages, getLanguageByCode } from '@infobits/intl';

// Obtener un idioma por código
const english = getLanguageByCode('en');
console.log(english?.name); // English
console.log(english?.nativeName); // English

// Obtener todos los idiomas
console.log(languages.length); // 185
```

### Monedas

```typescript
import { currencies, getCurrencyByCode } from '@infobits/intl';

// Obtener una moneda por código
const usd = getCurrencyByCode('USD');
console.log(usd?.name); // US Dollar
console.log(usd?.symbol); // $

// Obtener todas las monedas
console.log(currencies.length); // 179
```

### Continentes

```typescript
import { continents } from '@infobits/intl';

// Obtener todos los continentes
continents.forEach(c => {
  console.log(c.name, c.code);
});
```

### Banderas

```typescript
import { flags } from '@infobits/intl';

// Obtener cadena SVG de bandera
const usaFlag = flags['US'];

// Usar en HTML
document.getElementById('flag').innerHTML = usaFlag;

// Usar en React
function Flag({ code }: { code: string }) {
  return <div dangerouslySetInnerHTML={{ __html: flags[code] }} />;
}
```

### Traducciones

```typescript
import {
  getCountryTranslation,
  getLanguageTranslation,
  getCurrencyTranslation
} from '@infobits/intl';

// Obtener nombre de país traducido
const countryName = getCountryTranslation('US', 'de'); // Vereinigte Staaten

// Obtener nombre de idioma traducido
const languageName = getLanguageTranslation('en', 'es'); // Inglés

// Obtener nombre de moneda traducido
const currencyName = getCurrencyTranslation('USD', 'fr'); // Dollar américain
```

## Características

- 248 países con códigos ISO 3166-1
- 185 idiomas con códigos ISO 639-1
- 179 monedas con códigos ISO 4217
- 7 continentes
- Banderas SVG de países (incrustadas en línea)
- Traducciones multilingües
- Soporte completo de TypeScript

## Licencia

MIT
