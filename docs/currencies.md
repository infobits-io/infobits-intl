---
title: Currencies
sidebar_position: 4
description: Working with currency data — lookup by code, list all currencies, access symbols and names.
---

# Currencies

The package includes 179+ currencies with ISO 4217 codes.

## Currency Properties

Each currency has the following fields:

| Field | Description |
|-------|-------------|
| `code` | ISO 4217 currency code (e.g. `"USD"`) |
| `nativeName` | Singular name (e.g. `"US Dollar"`) |
| `nativeNamePlural` | Plural name (e.g. `"US Dollars"`) |
| `symbol` | Currency symbol (e.g. `"$"`) |

## Lookup by Code

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

## List All Currencies

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

## Currency for a Country

Each country has a linked currency.

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

## Type-Safe Currency Codes

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
