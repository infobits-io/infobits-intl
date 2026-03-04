---
title: Valute
sidebar_position: 4
description: Lavorare con i dati delle valute — ricerca per codice, elenco di tutte le valute, accesso a simboli e nomi.
---

# Valute

Il pacchetto include oltre 179 valute con codici ISO 4217.

## Proprieta delle Valute

Ogni valuta ha i seguenti campi:

| Campo | Descrizione |
|-------|-------------|
| `code` | Codice valuta ISO 4217 (es. `"USD"`) |
| `nativeName` | Nome singolare (es. `"US Dollar"`) |
| `nativeNamePlural` | Nome plurale (es. `"US Dollars"`) |
| `symbol` | Simbolo della valuta (es. `"$"`) |

## Ricerca per Codice

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

## Elenco di Tutte le Valute

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

## Valuta di un Paese

Ogni paese ha una valuta collegata.

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

## Codici Valuta Type-Safe

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
