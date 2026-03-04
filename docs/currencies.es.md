---
title: Monedas
sidebar_position: 4
description: Trabajar con datos de monedas — busqueda por codigo, listar todas las monedas, acceder a simbolos y nombres.
---

# Monedas

El paquete incluye 179+ monedas con codigos ISO 4217.

## Propiedades de moneda

Cada moneda tiene los siguientes campos:

| Campo | Descripcion |
|-------|-------------|
| `code` | Codigo de moneda ISO 4217 (ej. `"USD"`) |
| `nativeName` | Nombre en singular (ej. `"US Dollar"`) |
| `nativeNamePlural` | Nombre en plural (ej. `"US Dollars"`) |
| `symbol` | Simbolo de la moneda (ej. `"$"`) |

## Busqueda por codigo

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

## Listar todas las monedas

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

## Moneda de un pais

Cada pais tiene una moneda vinculada.

```go
country, ok := intl.CountryByAlpha2("JP")
if ok {
    currency, cOk := intl.CurrencyByCode(country.Currency)
    if cOk {
        fmt.Println(currency.NativeName) // Japanese Yen
        fmt.Println(currency.Symbol)     // ¥
    }
}
```

```dart
final country = Country.fromAlpha2Code('JP');
print(country?.currency.nativeName); // Japanese Yen
print(country?.currency.symbol);     // ¥
```

```typescript
import { getCountryByAlpha2, getCurrencyByCode } from 'infobits-intl';

const country = getCountryByAlpha2('JP');
if (country) {
  const currency = getCurrencyByCode(country.currency);
  console.log(currency?.nativeName); // Japanese Yen
  console.log(currency?.symbol);     // ¥
}
```

## Codigos de moneda con seguridad de tipos

```go
code := intl.CurrencyEUR
currency := code.Currency()
fmt.Println(currency.NativeName) // Euro
fmt.Println(currency.Symbol)     // €
```

```dart
print(Currency.eur.nativeName); // Euro
print(Currency.eur.symbol);     // €
```

```typescript
import { CurrencyCode, currencies } from 'infobits-intl';

const euro = currencies[CurrencyCode.EUR];
console.log(euro.nativeName); // Euro
console.log(euro.symbol);     // €
```
