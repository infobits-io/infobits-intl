---
title: Sprog
sidebar_position: 3
description: Arbejde med sprogdata -- opslag efter kode, liste over alle sprog, adgang til dialekter og oprindelige navne.
---

# Sprog

Pakken indeholder 185+ sprog med ISO 639-1-koder.

## Sprogegenskaber

Hvert sprog har foelgende felter:

| Felt | Beskrivelse |
|------|-------------|
| `code` | ISO 639-1-sprogkode (f.eks. `"en"`) |
| `nativeName` | Navn p selve sproget (f.eks. `"English"`) |
| `dialects` | Liste over sprogdialekter |
| `defaultFlagCode` | Landekode for et repraesentativt flag |

Hver dialekt har:

| Felt | Beskrivelse |
|------|-------------|
| `code` | Dialektkode |
| `nativeName` | Dialektens oprindelige navn |
| `flagCode` | Landekode for dialektens flag |

## Opslag efter kode

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

## Liste over alle sprog

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

## Dialekter

Sprog kan have dialekter, der repraesenterer regionale variationer.

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

## Sprogflag

Sprog kan have et repraesentativt flag baseret p deres standardflagkode.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // SVG-strenglaengde
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // SVG-streng eller null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Typesikre sprogkoder

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
