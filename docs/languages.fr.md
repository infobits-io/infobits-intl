---
title: Langues
sidebar_position: 3
description: Travailler avec les donnees de langues -- recherche par code, lister toutes les langues, acceder aux dialectes et noms natifs.
---

# Langues

Le package inclut 185+ langues avec les codes ISO 639-1.

## Proprietes d'une langue

Chaque langue possede les champs suivants :

| Champ | Description |
|-------|-------------|
| `code` | Code de langue ISO 639-1 (ex. `"en"`) |
| `nativeName` | Nom dans la langue elle-meme (ex. `"English"`) |
| `dialects` | Liste des dialectes de la langue |
| `defaultFlagCode` | Code pays pour un drapeau representatif |

Chaque dialecte possede :

| Champ | Description |
|-------|-------------|
| `code` | Code du dialecte |
| `nativeName` | Nom natif du dialecte |
| `flagCode` | Code pays pour le drapeau du dialecte |

## Recherche par code

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

## Lister toutes les langues

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

## Dialectes

Les langues peuvent avoir des dialectes representant des variations regionales.

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

## Drapeaux de langues

Les langues peuvent avoir un drapeau representatif base sur leur code de drapeau par defaut.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        fmt.Println(len(flag)) // longueur de la chaine SVG
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // chaine SVG ou null
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Codes de langues type-safe

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
