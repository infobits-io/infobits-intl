---
title: Drapeaux
sidebar_position: 7
description: Travailler avec les drapeaux SVG de pays et les drapeaux emoji — rendu, acces et affichage des drapeaux dans votre application.
---

# Drapeaux

Le package inclut 256 drapeaux SVG de pays, minifies et integres sous forme de chaines de caracteres. Les drapeaux sont indexes par les codes pays ISO 3166-1 alpha-2.

## Drapeaux SVG

### Obtenir un drapeau par code pays

```go
svg, ok := intl.GetFlag("US")
if ok {
    fmt.Println(len(svg)) // SVG markup length
}

// Or access the map directly
svg = intl.Flags["US"]
```

```dart
// Via the global map
final svg = countryFlags['US'];

// Via a Country instance
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Using the lookup function
const svg = getFlag('US');

// Or access the map directly
const svgDirect = flags['US'];
```

### Afficher les drapeaux SVG

Les chaines SVG peuvent etre utilisees directement dans votre application :

```go
// Write to an HTML template
fmt.Fprintf(w, `<div class="flag">%s</div>`, intl.Flags["NO"])

// Save to a file
os.WriteFile("flag.svg", []byte(intl.Flags["NO"]), 0644)
```

```dart
// Render as a Flutter widget with the built-in flag helper
final country = Country.fromAlpha2Code('NO');
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);

// Or use the raw SVG string
final svg = countryFlags['NO'];
```

```typescript
import { getFlag } from 'infobits-intl';

// In a web application
const svg = getFlag('NO');
if (svg) {
  document.getElementById('flag-container')!.innerHTML = svg;
}
```

## Drapeaux emoji

Les drapeaux emoji sont generes a partir des codes pays en utilisant les symboles indicateurs regionaux.

```go
code := intl.CountryUS
fmt.Println(code.EmojiFlag())

code = intl.CountryNO
fmt.Println(code.EmojiFlag())
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.emojiFlag);

// Or directly from an enum
print(Country.norway.emojiFlag);
```

```typescript
import { getEmojiFlag } from 'infobits-intl';

console.log(getEmojiFlag('US'));
console.log(getEmojiFlag('NO'));
```

## Drapeaux de langues

Les langues peuvent avoir un drapeau representatif via leur `defaultFlagCode`.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        // Use the French flag SVG
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // Resolves via defaultFlagCode
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Tous les drapeaux disponibles

Parcourir tous les drapeaux disponibles :

```go
for code, svg := range intl.Flags {
    fmt.Printf("%s: %d bytes\n", code, len(svg))
}
```

```dart
countryFlags.forEach((code, svg) {
  print('$code: ${svg.length} bytes');
});
```

```typescript
import { flags } from 'infobits-intl';

for (const [code, svg] of Object.entries(flags)) {
  console.log(`${code}: ${svg.length} bytes`);
}
```
