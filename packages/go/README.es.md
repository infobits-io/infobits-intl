# infobits-intl-go

Datos de internacionalización para países, idiomas, monedas y continentes con traducciones y banderas SVG.

## Instalación

```bash
go get github.com/infobits-io/infobits-intl-go
```

## Uso

### Países

```go
package main

import (
    "fmt"
    intl "github.com/infobits-io/infobits-intl-go"
)

func main() {
    // Obtener un país por código
    usa := intl.CountryByAlpha2Code("US")
    fmt.Println(usa.Name)        // United States
    fmt.Println(usa.Alpha3Code)  // USA
    fmt.Println(usa.Capital)     // Washington, D.C.
    fmt.Println(usa.CallingCode) // 1

    // Obtener todos los países
    allCountries := intl.Countries
    fmt.Println(len(allCountries)) // 248

    // Filtrar países por continente
    for _, c := range intl.Countries {
        if c.Continent == intl.ContinentEurope {
            fmt.Println(c.Name)
        }
    }
}
```

### Idiomas

```go
// Obtener un idioma por código
english := intl.LanguageByCode("en")
fmt.Println(english.Name)       // English
fmt.Println(english.NativeName) // English

// Obtener todos los idiomas
fmt.Println(len(intl.Languages)) // 185
```

### Monedas

```go
// Obtener una moneda por código
usd := intl.CurrencyByCode("USD")
fmt.Println(usd.Name)   // US Dollar
fmt.Println(usd.Symbol) // $

// Obtener todas las monedas
fmt.Println(len(intl.Currencies)) // 179
```

### Continentes

```go
// Obtener todos los continentes
for _, c := range intl.Continents {
    fmt.Println(c.Name, c.Code)
}

// Usar constantes de continente
fmt.Println(intl.ContinentEurope.Name) // Europe
```

### Banderas

```go
// Obtener cadena SVG de bandera
usaFlag := intl.Flags["US"]

// Usar en un handler HTTP
func flagHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Write([]byte(intl.Flags[code]))
}
```

### Traducciones

```go
import "github.com/infobits-io/infobits-intl-go/i18n"

// Obtener nombre de país traducido
countryName := i18n.CountryName("US", "de") // Vereinigte Staaten

// Obtener nombre de idioma traducido
languageName := i18n.LanguageName("en", "es") // Inglés

// Obtener nombre de moneda traducido
currencyName := i18n.CurrencyName("USD", "fr") // Dollar américain
```

## Características

- 248 países con códigos ISO 3166-1
- 185 idiomas con códigos ISO 639-1
- 179 monedas con códigos ISO 4217
- 7 continentes
- Banderas SVG de países (incrustadas en línea)
- Traducciones multilingües

## Licencia

MIT
