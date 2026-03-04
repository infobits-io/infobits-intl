# infobits-intl-go

Internationaliseringsdata for lande, sprog, valutaer og kontinenter med oversættelser og SVG-flag.

## Installation

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Brug

### Lande

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Hent et land efter kode
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Hent alle lande
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Filtrer lande efter kontinent
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Sprog

```go
// Hent et sprog efter kode
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Hent alle sprog
fmt.Println(len(intl.Languages)) // 185
```

### Valutaer

```go
// Hent en valuta efter kode
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Hent alle valutaer
fmt.Println(len(intl.Currencies)) // 179
```

### Kontinenter

```go
// Hent alle kontinenter
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Brug kontinentkonstanter
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Flag

```go
// Hent SVG-flagstreng
usaFlag := intl.Flags["US"]

// Brug i HTTP-handler
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Oversættelser

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Hent oversat landenavn
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Hent oversat sprognavn
languageName := i18n.LanguageName("en", "es") // Inglés

// Hent oversat valutanavn
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Funktioner

- 248 lande med ISO 3166-1-koder
- 185 sprog med ISO 639-1-koder
- 179 valutaer med ISO 4217-koder
- 7 kontinenter
- SVG-landeflag (inline indlejret)
- Flersprogede oversættelser

## Licens

MIT
