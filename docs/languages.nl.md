---
title: Talen
sidebar_position: 3
description: Werken met taalgegevens — opzoeken op code, alle talen weergeven, dialecten en oorspronkelijke namen gebruiken.
---

# Talen

Het pakket bevat 185+ talen met ISO 639-1-codes.

## Taaleigenschappen

Elke taal heeft de volgende velden:

| Veld | Beschrijving |
|------|--------------|
| `code` | ISO 639-1-taalcode (bijv. `"en"`) |
| `nativeName` | Naam in de taal zelf (bijv. `"English"`) |
| `dialects` | Lijst van taaldialecten |
| `defaultFlagCode` | Landcode voor een representatieve vlag |

Elk dialect heeft:

| Veld | Beschrijving |
|------|--------------|
| `code` | Dialectcode |
| `nativeName` | Oorspronkelijke naam van het dialect |
| `flagCode` | Landcode voor de vlag van het dialect |

## Opzoeken op code

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

## Alle talen weergeven

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

## Dialecten

Talen kunnen dialecten hebben die regionale variaties vertegenwoordigen.

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

## Taalvlaggen

Talen kunnen een representatieve vlag hebben op basis van hun standaard vlagcode.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // lengte van SVG-string
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // SVG-string of null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Typeveilige taalcodes

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
