package data

// Country represents a country with all its metadata.
type Country struct {
	ID           string   `json:"id"`
	Alpha2Code   string   `json:"alpha2Code"`
	Alpha3Code   string   `json:"alpha3Code"`
	NumericCode  int      `json:"numericCode"`
	NativeName   string   `json:"nativeName"`
	Capital      string   `json:"capital"`
	MainLanguage string   `json:"mainLanguage"`
	Languages    []string `json:"languages"`
	TLD          string   `json:"tld"`
	CallingCode  int      `json:"callingCode"`
	Continent    string   `json:"continent"`
	Currency     string   `json:"currency"`
}

// Language represents a language with its metadata.
type Language struct {
	ID              string            `json:"id"`
	Code            string            `json:"code"`
	NativeName      string            `json:"nativeName"`
	Dialects        []LanguageDialect `json:"dialects,omitempty"`
	DefaultFlagCode string            `json:"defaultFlagCode,omitempty"`
}

// LanguageDialect represents a dialect of a language.
type LanguageDialect struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	NativeName string `json:"nativeName"`
	FlagCode   string `json:"flagCode,omitempty"`
}

// Currency represents a currency with its metadata.
type Currency struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	NativeName       string `json:"nativeName"`
	NativeNamePlural string `json:"nativeNamePlural"`
	Symbol           string `json:"symbol"`
}

// Continent represents a continent.
type Continent struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// IntlData contains all internationalization data.
type IntlData struct {
	Countries  []Country   `json:"countries"`
	Languages  []Language  `json:"languages"`
	Currencies []Currency  `json:"currencies"`
	Continents []Continent `json:"continents"`
}

// CountriesFile is the JSON file structure for countries.
type CountriesFile struct {
	Countries []Country `json:"countries"`
}

// LanguagesFile is the JSON file structure for languages.
type LanguagesFile struct {
	Languages []Language `json:"languages"`
}

// CurrenciesFile is the JSON file structure for currencies.
type CurrenciesFile struct {
	Currencies []Currency `json:"currencies"`
}

// ContinentsFile is the JSON file structure for continents.
type ContinentsFile struct {
	Continents []Continent `json:"continents"`
}
