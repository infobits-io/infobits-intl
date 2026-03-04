---
title: Sprachen
sidebar_position: 3
description: Arbeiten mit Sprachdaten — Suche nach Code, alle Sprachen auflisten, Zugriff auf Dialekte und Eigennamen.
---

# Sprachen

Das Paket enthält 185+ Sprachen mit ISO 639-1-Codes.

## Spracheigenschaften

Jede Sprache hat die folgenden Felder:

| Feld | Beschreibung |
|------|--------------|
| `code` | ISO 639-1-Sprachcode (z.B. `"en"`) |
| `nativeName` | Name in der Sprache selbst (z.B. `"English"`) |
| `dialects` | Liste der Sprachdialekte |
| `defaultFlagCode` | Ländercode für eine repräsentative Flagge |

Jeder Dialekt hat:

| Feld | Beschreibung |
|------|--------------|
| `code` | Dialektcode |
| `nativeName` | Eigenname des Dialekts |
| `flagCode` | Ländercode für die Flagge des Dialekts |

## Suche nach Code

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

## Alle Sprachen auflisten

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

## Dialekte

Sprachen können Dialekte haben, die regionale Varianten darstellen.

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

## Sprachflaggen

Sprachen können eine repräsentative Flagge haben, basierend auf ihrem Standard-Flaggencode.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // Länge des SVG-Strings
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // SVG-String oder null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Typsichere Sprachcodes

```go
code := intl.LanguageFR
lang := code.Language()
fmt.Println(lang.NativeName) // Français
```

```dart
print(Language.french.nativeName); // Français
print(Language.french.code);       // fr
```

```typescript
import { LanguageCode, languages } from 'infobits-intl';

const french = languages[LanguageCode.FR];
console.log(french.nativeName); // Français
```
