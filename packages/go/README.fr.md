# infobits-intl-go

Données d'internationalisation pour les pays, langues, devises et continents avec traductions et drapeaux SVG.

## Installation

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Utilisation

### Pays

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Obtenir un pays par code
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Obtenir tous les pays
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Filtrer les pays par continent
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Langues

```go
// Obtenir une langue par code
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Obtenir toutes les langues
fmt.Println(len(intl.Languages)) // 185
```

### Devises

```go
// Obtenir une devise par code
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Obtenir toutes les devises
fmt.Println(len(intl.Currencies)) // 179
```

### Continents

```go
// Obtenir tous les continents
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Utiliser les constantes de continent
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Drapeaux

```go
// Obtenir la chaîne SVG du drapeau
usaFlag := intl.Flags["US"]

// Utiliser dans un handler HTTP
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Traductions

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Obtenir le nom de pays traduit
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Obtenir le nom de langue traduit
languageName := i18n.LanguageName("en", "es") // Inglés

// Obtenir le nom de devise traduit
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Fonctionnalités

- 248 pays avec codes ISO 3166-1
- 185 langues avec codes ISO 639-1
- 179 devises avec codes ISO 4217
- 7 continents
- Drapeaux SVG des pays (intégrés en ligne)
- Traductions multilingues

## Licence

MIT
