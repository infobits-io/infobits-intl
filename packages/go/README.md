# infobits-intl-go

Internationalization data for countries, languages, currencies, and continents with translations and SVG flags.

## Installation

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Usage

### Countries

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Get a country by code
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Get all countries
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Filter countries by continent
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Languages

```go
// Get a language by code
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Get all languages
fmt.Println(len(intl.Languages)) // 185
```

### Currencies

```go
// Get a currency by code
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Get all currencies
fmt.Println(len(intl.Currencies)) // 179
```

### Continents

```go
// Get all continents
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Use continent constants
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Flags

```go
// Get SVG flag string
usaFlag := intl.Flags["US"]

// Use in HTTP handler
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Translations

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Get translated country name
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Get translated language name
languageName := i18n.LanguageName("en", "es") // Inglés

// Get translated currency name
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Features

- 248 countries with ISO 3166-1 codes
- 185 languages with ISO 639-1 codes
- 179 currencies with ISO 4217 codes
- 7 continents
- SVG country flags (inline embedded)
- Multi-language translations

## License

MIT
