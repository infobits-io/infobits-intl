---
title: Idiomas
sidebar_position: 3
description: Trabajar con datos de idiomas — busqueda por codigo, listar todos los idiomas, acceder a dialectos y nombres nativos.
---

# Idiomas

El paquete incluye 185+ idiomas con codigos ISO 639-1.

## Propiedades de idioma

Cada idioma tiene los siguientes campos:

| Campo | Descripcion |
|-------|-------------|
| `code` | Codigo de idioma ISO 639-1 (ej. `"en"`) |
| `nativeName` | Nombre en el propio idioma (ej. `"English"`) |
| `dialects` | Lista de dialectos del idioma |
| `defaultFlagCode` | Codigo de pais para una bandera representativa |

Cada dialecto tiene:

| Campo | Descripcion |
|-------|-------------|
| `code` | Codigo del dialecto |
| `nativeName` | Nombre nativo del dialecto |
| `flagCode` | Codigo de pais para la bandera del dialecto |

## Busqueda por codigo

```go
lang, ok := intl.LanguageByCode("en")
if ok {
    fmt.Println(lang.NativeName) // English
    fmt.Println(lang.Code)       // en
}
```

```dart
final lang = Language.fromCode('en');
print(lang?.nativeName); // English
print(lang?.code);       // en
```

```typescript
import { getLanguageByCode } from 'infobits-intl';

const lang = getLanguageByCode('en');
console.log(lang?.nativeName); // English
console.log(lang?.code);       // en
```

## Listar todos los idiomas

```go
allLanguages := intl.AllLanguages()
fmt.Println(len(allLanguages)) // 185
```

```dart
final allLanguages = Language.values;
print(allLanguages.length); // 187
```

```typescript
import { languages } from 'infobits-intl';

const allLanguages = Object.values(languages);
console.log(allLanguages.length);
```

## Dialectos

Los idiomas pueden tener dialectos que representan variaciones regionales.

```go
lang, ok := intl.LanguageByCode("en")
if ok {
    for _, d := range lang.Dialects {
        fmt.Printf("%s (%s)\n", d.NativeName, d.FlagCode)
    }
}
```

```dart
final lang = Language.fromCode('en');
if (lang != null) {
  for (final d in lang.dialects) {
    print('${d.nativeName} (${d.flagCode})');
  }
}
```

```typescript
import { getLanguageByCode } from 'infobits-intl';

const lang = getLanguageByCode('en');
lang?.dialects.forEach(d => {
  console.log(`${d.nativeName} (${d.flagCode})`);
});
```

## Banderas de idioma

Los idiomas pueden tener una bandera representativa basada en su codigo de bandera predeterminado.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // longitud de la cadena SVG
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // cadena SVG o null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Codigos de idioma con seguridad de tipos

```go
code := intl.LanguageFR
lang := code.Language()
fmt.Println(lang.NativeName) // Francais
```

```dart
print(Language.french.nativeName); // Francais
print(Language.french.code);       // fr
```

```typescript
import { LanguageCode, languages } from 'infobits-intl';

const french = languages[LanguageCode.FR];
console.log(french.nativeName); // Francais
```
