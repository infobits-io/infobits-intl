# infobits-intl-go

Dati di internazionalizzazione per paesi, lingue, valute e continenti con traduzioni e bandiere SVG.

## Installazione

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Utilizzo

### Paesi

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Ottenere un paese per codice
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Ottenere tutti i paesi
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Filtrare i paesi per continente
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Lingue

```go
// Ottenere una lingua per codice
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Ottenere tutte le lingue
fmt.Println(len(intl.Languages)) // 185
```

### Valute

```go
// Ottenere una valuta per codice
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Ottenere tutte le valute
fmt.Println(len(intl.Currencies)) // 179
```

### Continenti

```go
// Ottenere tutti i continenti
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Usare le costanti dei continenti
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Bandiere

```go
// Ottenere la stringa SVG della bandiera
usaFlag := intl.Flags["US"]

// Usare in un handler HTTP
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Traduzioni

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Ottenere il nome del paese tradotto
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Ottenere il nome della lingua tradotto
languageName := i18n.LanguageName("en", "es") // Inglés

// Ottenere il nome della valuta tradotto
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Funzionalità

- 248 paesi con codici ISO 3166-1
- 185 lingue con codici ISO 639-1
- 179 valute con codici ISO 4217
- 7 continenti
- Bandiere SVG dei paesi (incorporate in linea)
- Traduzioni multilingue

## Licenza

MIT
