---
title: Flaggen
sidebar_position: 7
description: Arbeiten mit SVG-Laenderflaggen und Emoji-Flaggen — Rendering, Zugriff und Anzeige von Flaggen in Ihrer Anwendung.
---

# Flaggen

Das Paket enthält 256 SVG-Länderflaggen, minifiziert und als Strings eingebettet. Flaggen sind nach ISO 3166-1 Alpha-2-Ländercodes indiziert.

## SVG-Flaggen

### Eine Flagge nach Ländercode abrufen

```go
svg, ok := intl.GetFlag("US")
if ok {
    fmt.Println(len(svg)) // SVG-Markup-Länge
}

// Oder direkt auf die Map zugreifen
svg = intl.Flags["US"]
```

```dart
// Über die globale Map
final svg = countryFlags['US'];

// Über eine Country-Instanz
final country = Country.fromAlpha2Code('US');
final flagSvg = country?.flagSvg;
```

```typescript
import { getFlag, flags } from 'infobits-intl';

// Mit der Suchfunktion
const svg = getFlag('US');

// Oder direkt auf die Map zugreifen
const svgDirect = flags['US'];
```

### SVG-Flaggen rendern

Die SVG-Strings können direkt in Ihrer Anwendung verwendet werden:

```go
// In ein HTML-Template schreiben
fmt.Fprintf(w, `<div class="flag">%s</div>`, intl.Flags["NO"])

// In eine Datei speichern
os.WriteFile("flag.svg", []byte(intl.Flags["NO"]), 0644)
```

```dart
// Als Flutter-Widget mit dem eingebauten Flag-Helper rendern
final country = Country.fromAlpha2Code('NO');
country?.flag(
  shape: FlagShape.rectangle,
  width: 32,
  height: 24,
);

// Oder den rohen SVG-String verwenden
final svg = countryFlags['NO'];
```

```typescript
import { getFlag } from 'infobits-intl';

// In einer Webanwendung
const svg = getFlag('NO');
if (svg) {
  document.getElementById('flag-container')!.innerHTML = svg;
}
```

## Emoji-Flaggen

Emoji-Flaggen werden aus Ländercodes mithilfe regionaler Indikatorsymbole generiert.

```go
code := intl.CountryUS
fmt.Println(code.EmojiFlag())

code = intl.CountryNO
fmt.Println(code.EmojiFlag())
```

```dart
final country = Country.fromAlpha2Code('US');
print(country?.emojiFlag);

// Oder direkt über ein Enum
print(Country.norway.emojiFlag);
```

```typescript
import { getEmojiFlag } from 'infobits-intl';

console.log(getEmojiFlag('US'));
console.log(getEmojiFlag('NO'));
```

## Sprachflaggen

Sprachen können über ihren `defaultFlagCode` eine repräsentative Flagge haben.

```go
lang, ok := intl.LanguageByCode("fr")
if ok && lang.DefaultFlagCode != "" {
    flag, flagOk := intl.GetFlag(lang.DefaultFlagCode)
    if flagOk {
        // Die französische Flagge als SVG verwenden
    }
}
```

```dart
final lang = Language.fromCode('fr');
final svg = lang?.flagSvg; // Wird über defaultFlagCode aufgelöst
```

```typescript
import { getLanguageByCode, getFlag } from 'infobits-intl';

const lang = getLanguageByCode('fr');
if (lang?.defaultFlagCode) {
  const svg = getFlag(lang.defaultFlagCode);
}
```

## Alle verfügbaren Flaggen

Über alle verfügbaren Flaggen iterieren:

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
