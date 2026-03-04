---
title: Valutaer
sidebar_position: 4
description: Arbejde med valutadata -- opslag efter kode, liste over alle valutaer, adgang til symboler og navne.
---

# Valutaer

Pakken indeholder 179+ valutaer med ISO 4217-koder.

## Valutaegenskaber

Hver valuta har foelgende felter:

| Felt | Beskrivelse |
|------|-------------|
| `code` | ISO 4217-valutakode (f.eks. `"USD"`) |
| `nativeName` | Ental-navn (f.eks. `"US Dollar"`) |
| `nativeNamePlural` | Flertal-navn (f.eks. `"US Dollars"`) |
| `symbol` | Valutasymbol (f.eks. `"$"`) |

## Opslag efter kode

```go
currency, ok := intl.CurrencyByCode("USD")
if ok {
    fmt.Println(currency.NativeName)       // US Dollar
    fmt.Println(currency.NativeNamePlural) // US Dollars
    fmt.Println(currency.Symbol)           // $
}
```

```dart
final currency = Currency.fromCode('USD');
print(currency?.nativeName);       // US Dollar
print(currency?.nativeNamePlural); // US Dollars
print(currency?.symbol);           // $
```

```typescript
import { getCurrencyByCode } from 'infobits-intl';

const currency = getCurrencyByCode('USD');
console.log(currency?.nativeName);       // US Dollar
console.log(currency?.nativeNamePlural); // US Dollars
console.log(currency?.symbol);           // $
```

## Liste over alle valutaer

```go
allCurrencies := intl.AllCurrencies()
fmt.Println(len(allCurrencies)) // 179
```

```dart
final allCurrencies = Currency.values;
print(allCurrencies.length); // 181
```

```typescript
import { currencies } from 'infobits-intl';

const allCurrencies = Object.values(currencies);
console.log(allCurrencies.length);
```

## Valuta for et land

Hvert land har en tilknyttet valuta.

```go
country, ok := intl.CountryByAlpha2("JP")
if ok {
    currency, cOk := intl.CurrencyByCode(country.Currency)
    if cOk {
        fmt.Println(currency.NativeName) // Japanese Yen
        fmt.Println(currency.Symbol)     // ??
    }
}
```

```dart
final country = Country.fromAlpha2Code('JP');
print(country?.currency.nativeName); // Japanese Yen
print(country?.currency.symbol);     // ??
```

```typescript
import { getCountryByAlpha2, getCurrencyByCode } from 'infobits-intl';

const country = getCountryByAlpha2('JP');
if (country) {
  const currency = getCurrencyByCode(country.currency);
  console.log(currency?.nativeName); // Japanese Yen
  console.log(currency?.symbol);     // ??
}
```

## Typesikre valutakoder

```go
code := intl.CurrencyEUR
currency := code.Currency()
fmt.Println(currency.NativeName) // Euro
fmt.Println(currency.Symbol)     // ???
```

```dart
print(Currency.eur.nativeName); // Euro
print(Currency.eur.symbol);     // ???
```

```typescript
import { CurrencyCode, currencies } from 'infobits-intl';

const euro = currencies[CurrencyCode.EUR];
console.log(euro.nativeName); // Euro
console.log(euro.symbol);     // ???
```
