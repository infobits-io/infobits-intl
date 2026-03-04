---
title: Devises
sidebar_position: 4
description: Travailler avec les donnees de devises -- recherche par code, lister toutes les devises, acceder aux symboles et noms.
---

# Devises

Le package inclut 179+ devises avec les codes ISO 4217.

## Proprietes d'une devise

Chaque devise possede les champs suivants :

| Champ | Description |
|-------|-------------|
| `code` | Code de devise ISO 4217 (ex. `"USD"`) |
| `nativeName` | Nom singulier (ex. `"US Dollar"`) |
| `nativeNamePlural` | Nom pluriel (ex. `"US Dollars"`) |
| `symbol` | Symbole de la devise (ex. `"$"`) |

## Recherche par code

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

## Lister toutes les devises

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

## Devise d'un pays

Chaque pays est associe a une devise.

```go
country, ok := intl.CountryByAlpha2("JP")
if ok {
    currency, cOk := intl.CurrencyByCode(country.Currency)
    if cOk {
        fmt.Println(currency.NativeName) // Japanese Yen
        fmt.Println(currency.Symbol)     // yen symbol
    }
}
```

```dart
final country = Country.fromAlpha2Code('JP');
print(country?.currency.nativeName); // Japanese Yen
print(country?.currency.symbol);     // yen symbol
```

```typescript
import { getCountryByAlpha2, getCurrencyByCode } from 'infobits-intl';

const country = getCountryByAlpha2('JP');
if (country) {
  const currency = getCurrencyByCode(country.currency);
  console.log(currency?.nativeName); // Japanese Yen
  console.log(currency?.symbol);     // yen symbol
}
```

## Codes de devises type-safe

```go
code := intl.CurrencyEUR
currency := code.Currency()
fmt.Println(currency.NativeName) // Euro
fmt.Println(currency.Symbol)     // euro sign
```

```dart
print(Currency.eur.nativeName); // Euro
print(Currency.eur.symbol);     // euro sign
```

```typescript
import { CurrencyCode, currencies } from 'infobits-intl';

const euro = currencies[CurrencyCode.EUR];
console.log(euro.nativeName); // Euro
console.log(euro.symbol);     // euro sign
```
