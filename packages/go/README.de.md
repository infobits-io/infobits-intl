# infobits-intl-go

Internationalisierungsdaten für Länder, Sprachen, Währungen und Kontinente mit Übersetzungen und SVG-Flaggen.

## Installation

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Verwendung

### Länder

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Ein Land nach Code abrufen
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Alle Länder abrufen
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Länder nach Kontinent filtern
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Sprachen

```go
// Eine Sprache nach Code abrufen
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Alle Sprachen abrufen
fmt.Println(len(intl.Languages)) // 185
```

### Währungen

```go
// Eine Währung nach Code abrufen
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Alle Währungen abrufen
fmt.Println(len(intl.Currencies)) // 179
```

### Kontinente

```go
// Alle Kontinente abrufen
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Kontinentkonstanten verwenden
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Flaggen

```go
// SVG-Flaggenstring abrufen
usaFlag := intl.Flags["US"]

// In HTTP-Handler verwenden
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Übersetzungen

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Übersetzten Ländernamen abrufen
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Übersetzten Sprachnamen abrufen
languageName := i18n.LanguageName("en", "es") // Inglés

// Übersetzten Währungsnamen abrufen
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Funktionen

- 248 Länder mit ISO 3166-1-Codes
- 185 Sprachen mit ISO 639-1-Codes
- 179 Währungen mit ISO 4217-Codes
- 7 Kontinente
- SVG-Länderflaggen (inline eingebettet)
- Mehrsprachige Übersetzungen

## Lizenz

MIT
