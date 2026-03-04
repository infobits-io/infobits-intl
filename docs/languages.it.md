---
title: Lingue
sidebar_position: 3
description: Lavorare con i dati delle lingue — ricerca per codice, elenco di tutte le lingue, accesso ai dialetti e nomi nativi.
---

# Lingue

Il pacchetto include oltre 185 lingue con codici ISO 639-1.

## Proprieta delle Lingue

Ogni lingua ha i seguenti campi:

| Campo | Descrizione |
|-------|-------------|
| `code` | Codice lingua ISO 639-1 (es. `"en"`) |
| `nativeName` | Nome nella lingua stessa (es. `"English"`) |
| `dialects` | Elenco dei dialetti della lingua |
| `defaultFlagCode` | Codice paese per una bandiera rappresentativa |

Ogni dialetto ha:

| Campo | Descrizione |
|-------|-------------|
| `code` | Codice del dialetto |
| `nativeName` | Nome nativo del dialetto |
| `flagCode` | Codice paese per la bandiera del dialetto |

## Ricerca per Codice

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

## Elenco di Tutte le Lingue

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

## Dialetti

Le lingue possono avere dialetti che rappresentano variazioni regionali.

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

## Bandiere delle Lingue

Le lingue possono avere una bandiera rappresentativa basata sul loro codice bandiera predefinito.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // lunghezza della stringa SVG
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // stringa SVG o null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Codici Lingua Type-Safe

```go
code := intl.LanguageFR
lang := code.Language()
fmt.Println(lang.NativeName) // Fran??ais
```

```dart
print(Language.french.nativeName); // Fran??ais
print(Language.french.code);       // fr
```

```typescript
import { LanguageCode, languages } from 'infobits-intl';

const french = languages[LanguageCode.FR];
console.log(french.nativeName); // Fran??ais
```
