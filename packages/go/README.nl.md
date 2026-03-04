# infobits-intl-go

Internationalisatiegegevens voor landen, talen, valuta's en continenten met vertalingen en SVG-vlaggen.

## Installatie

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Gebruik

### Landen

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Een land ophalen op code
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Alle landen ophalen
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Landen filteren op continent
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Talen

```go
// Een taal ophalen op code
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Alle talen ophalen
fmt.Println(len(intl.Languages)) // 185
```

### Valuta's

```go
// Een valuta ophalen op code
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Alle valuta's ophalen
fmt.Println(len(intl.Currencies)) // 179
```

### Continenten

```go
// Alle continenten ophalen
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Continentconstanten gebruiken
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Vlaggen

```go
// SVG-vlaggenstring ophalen
usaFlag := intl.Flags["US"]

// Gebruiken in een HTTP-handler
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Vertalingen

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Vertaalde landnaam ophalen
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Vertaalde taalnaam ophalen
languageName := i18n.LanguageName("en", "es") // Inglés

// Vertaalde valutanaam ophalen
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Kenmerken

- 248 landen met ISO 3166-1-codes
- 185 talen met ISO 639-1-codes
- 179 valuta's met ISO 4217-codes
- 7 continenten
- SVG-landvlaggen (inline ingebed)
- Meertalige vertalingen

## Licentie

MIT
