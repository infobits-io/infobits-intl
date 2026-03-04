---
title: Languages
sidebar_position: 3
description: Working with language data — lookup by code, list all languages, access dialects and native names.
---

# Languages

The package includes 185+ languages with ISO 639-1 codes.

## Language Properties

Each language has the following fields:

| Field | Description |
|-------|-------------|
| `code` | ISO 639-1 language code (e.g. `"en"`) |
| `nativeName` | Name in the language itself (e.g. `"English"`) |
| `dialects` | List of language dialects |
| `defaultFlagCode` | Country code for a representative flag |

Each dialect has:

| Field | Description |
|-------|-------------|
| `code` | Dialect code |
| `nativeName` | Native name of the dialect |
| `flagCode` | Country code for the dialect's flag |

## Lookup by Code

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

## List All Languages

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

## Dialects

Languages can have dialects representing regional variations.

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

## Language Flags

Languages can have a representative flag based on their default flag code.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // SVG string length
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // SVG string or null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Type-Safe Language Codes

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
