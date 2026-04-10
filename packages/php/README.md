# infobits/intl

Internationalization data for countries, languages, currencies, and continents with translations and SVG flags.

## Installation

```bash
composer require infobits/intl
```

Requires PHP 8.1+.

## Usage

### Countries

```php
use Infobits\Intl\Country;

// Get a country by alpha-2 code
$usa = Country::tryFrom('US');
echo $usa->nativeName();  // United States
echo $usa->alpha3Code();  // USA
echo $usa->capital();     // Washington, D.C.
echo $usa->callingCode(); // 1

// Get all countries
$allCountries = Country::cases();
echo count($allCountries); // 248

// Lookup by alpha-3 code
$country = Country::fromAlpha3('USA');

// Lookup by numeric code
$country = Country::fromNumeric(840);
```

### Languages

```php
use Infobits\Intl\Language;

// Get a language by code
$english = Language::tryFrom('en');
echo $english->nativeName(); // english
echo $english->code();       // en

// Case-insensitive lookup
$english = Language::fromCode('EN');

// Get all languages
echo count(Language::cases()); // 184
```

### Currencies

```php
use Infobits\Intl\Currency;

// Get a currency by code
$usd = Currency::tryFrom('USD');
echo $usd->nativeName();       // United States Dollar
echo $usd->nativeNamePlural(); // United States Dollars
echo $usd->symbol();           // $

// Case-insensitive lookup
$usd = Currency::fromCode('usd');

// Get all currencies
echo count(Currency::cases()); // 179
```

### Continents

```php
use Infobits\Intl\Continent;

// Get a continent by code
$europe = Continent::tryFrom('EU');
echo $europe->label(); // Europe
echo $europe->code();  // EU

// Get all continents
foreach (Continent::cases() as $continent) {
    echo $continent->label() . ' (' . $continent->code() . ')';
}
```

### Flags

```php
use Infobits\Intl\Flags;

// Get SVG flag string
$usFlag = Flags::svg('US');

// Use in HTML
echo '<div class="flag">' . $usFlag . '</div>';

// Get all flags
$allFlags = Flags::all();
echo count($allFlags); // 256
```

### Emoji Flags

```php
use Infobits\Intl\Country;

$usa = Country::tryFrom('US');
echo $usa->emojiFlag(); // flag emoji
```

### Translations

```php
use Infobits\Intl\I18n\CountriesTranslations;
use Infobits\Intl\I18n\LanguagesTranslations;
use Infobits\Intl\I18n\CurrenciesTranslations;
use Infobits\Intl\I18n\ContinentsTranslations;
use Infobits\Intl\I18n\CapitalsTranslations;

// Get translated country name
$name = CountriesTranslations::get('US', 'de'); // Vereinigte Staaten

// Get all translations for a locale
$germanCountries = CountriesTranslations::all('de');

// Get available locales
$locales = CountriesTranslations::locales();

// Other translation classes work the same way
$languageName = LanguagesTranslations::get('en', 'fr');
$currencyName = CurrenciesTranslations::get('USD', 'it');
$continentName = ContinentsTranslations::get('EU', 'da');
$capitalName = CapitalsTranslations::get('JP', 'es');
```

## Features

- 248 countries with ISO 3166-1 codes
- 184 languages with ISO 639-1 codes
- 179 currencies with ISO 4217 codes
- 7 continents
- SVG country flags (inline embedded)
- Multi-language translations
- PHP 8.1 backed enums for type safety

## License

MIT
