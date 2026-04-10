---
title: Flags
sidebar_position: 7
description: Working with SVG country flags and emoji flags — rendering, accessing, and displaying flags in your application.
---

# Flags

The package includes 256 SVG country flags, minified and embedded as strings. Flags are keyed by ISO 3166-1 alpha-2 country codes.

## SVG Flags

### Get a Flag by Country Code

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

```php
use Infobits\Intl\Flags;

// Using the lookup function (case insensitive)
$svg = Flags::svg('US');

// Get all flags as an associative array
$allFlags = Flags::all();
$svg = $allFlags['US'];
```

### Rendering SVG Flags

The SVG strings can be used directly in your application:

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

```php
use Infobits\Intl\Flags;

// Output in a template
$svg = Flags::svg('NO');
echo '<div class="flag">' . $svg . '</div>';

// Save to a file
file_put_contents('flag.svg', Flags::svg('NO'));
```

## Emoji Flags

Emoji flags are generated from country codes using regional indicator symbols.

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

```php
use Infobits\Intl\Country;

echo Country::US->emojiFlag();
echo Country::NO->emojiFlag();
```

## Language Flags

Languages can have a representative flag via their `defaultFlagCode`.

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

```php
use Infobits\Intl\Language;
use Infobits\Intl\Flags;

$lang = Language::tryFrom('fr');
$flagCode = $lang?->defaultFlagCode();
if ($flagCode !== null) {
    $svg = Flags::svg($flagCode);
}
```

## All Available Flags

Iterate over all available flags:

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

```php
use Infobits\Intl\Flags;

foreach (Flags::all() as $code => $svg) {
    echo "$code: " . strlen($svg) . " bytes";
}
```
