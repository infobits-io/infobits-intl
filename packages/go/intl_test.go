package intl

import (
	"strings"
	"testing"

	"github.com/infobits-io/infobits-intl-go/i18n"
)

// Country tests

func TestAllCountries(t *testing.T) {
	t.Parallel()

	countries := AllCountries()
	if len(countries) == 0 {
		t.Error("Expected countries, got none")
	}

	if len(countries) < 200 {
		t.Errorf("Expected at least 200 countries, got %d", len(countries))
	}
}

func TestCountryByAlpha2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		code     string
		expected string
	}{
		{"US", "USA"},
		{"us", "USA"},
		{"GB", "GBR"},
		{"DE", "DEU"},
		{"FR", "FRA"},
		{"JP", "JPN"},
	}

	for _, tt := range tests {
		country, ok := CountryByAlpha2(tt.code)
		if !ok {
			t.Errorf("CountryByAlpha2(%q) not found", tt.code)

			continue
		}

		if country.Alpha3Code != tt.expected {
			t.Errorf("CountryByAlpha2(%q).Alpha3Code = %q, want %q", tt.code, country.Alpha3Code, tt.expected)
		}
	}
}

func TestCountryByAlpha2NotFound(t *testing.T) {
	t.Parallel()

	_, ok := CountryByAlpha2("XX")
	if ok {
		t.Error("Expected XX not to be found")
	}
}

func TestCountryByAlpha3(t *testing.T) {
	t.Parallel()

	tests := []struct {
		code     string
		expected string
	}{
		{"USA", "US"},
		{"usa", "US"},
		{"GBR", "GB"},
		{"DEU", "DE"},
		{"FRA", "FR"},
		{"JPN", "JP"},
	}

	for _, tt := range tests {
		country, ok := CountryByAlpha3(tt.code)
		if !ok {
			t.Errorf("CountryByAlpha3(%q) not found", tt.code)

			continue
		}

		if country.Alpha2Code != tt.expected {
			t.Errorf("CountryByAlpha3(%q).Alpha2Code = %q, want %q", tt.code, country.Alpha2Code, tt.expected)
		}
	}
}

func TestCountryCodeString(t *testing.T) {
	t.Parallel()

	code := CountryUS
	if code.String() != "US" {
		t.Errorf("CountryUS.String() = %q, want %q", code.String(), "US")
	}
}

func TestCountryCodeCountry(t *testing.T) {
	t.Parallel()

	code := CountryUS

	country := code.Country()
	if country.Alpha2Code != "US" {
		t.Errorf("CountryUS.Country().Alpha2Code = %q, want %q", country.Alpha2Code, "US")
	}

	if country.Alpha3Code != "USA" {
		t.Errorf("CountryUS.Country().Alpha3Code = %q, want %q", country.Alpha3Code, "USA")
	}
}

func TestCountryCodeEmojiFlag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		code     CountryCode
		expected string
	}{
		{CountryUS, "🇺🇸"},
		{CountryGB, "🇬🇧"},
		{CountryFR, "🇫🇷"},
		{CountryDE, "🇩🇪"},
		{CountryJP, "🇯🇵"},
	}

	for _, tt := range tests {
		emoji := tt.code.EmojiFlag()
		if emoji != tt.expected {
			t.Errorf("%s.EmojiFlag() = %q, want %q", tt.code, emoji, tt.expected)
		}
	}
}

func TestCountriesByContinent(t *testing.T) {
	t.Parallel()

	europeCountries := CountriesByContinent("europe")
	if len(europeCountries) == 0 {
		t.Error("Expected European countries, got none")
	}

	for _, c := range europeCountries {
		if c.Continent != "europe" {
			t.Errorf("Country %s has continent %q, want %q", c.Alpha2Code, c.Continent, "europe")
		}
	}
}

func TestCountryFields(t *testing.T) {
	t.Parallel()

	country, ok := CountryByAlpha2("US")
	if !ok {
		t.Fatal("US not found")
	}

	if country.Alpha2Code != "US" {
		t.Errorf("Alpha2Code = %q, want %q", country.Alpha2Code, "US")
	}

	if country.Alpha3Code != "USA" {
		t.Errorf("Alpha3Code = %q, want %q", country.Alpha3Code, "USA")
	}

	if country.NumericCode != 840 {
		t.Errorf("NumericCode = %d, want %d", country.NumericCode, 840)
	}

	if country.NativeName == "" {
		t.Error("NativeName is empty")
	}

	if country.MainLanguage == "" {
		t.Error("MainLanguage is empty")
	}

	if len(country.Languages) == 0 {
		t.Error("Languages is empty")
	}

	if country.Continent == "" {
		t.Error("Continent is empty")
	}

	if country.Currency == "" {
		t.Error("Currency is empty")
	}
}

// Language tests

func TestAllLanguages(t *testing.T) {
	t.Parallel()

	languages := AllLanguages()
	if len(languages) == 0 {
		t.Error("Expected languages, got none")
	}

	if len(languages) < 100 {
		t.Errorf("Expected at least 100 languages, got %d", len(languages))
	}
}

func TestLanguageByCode(t *testing.T) {
	t.Parallel()

	tests := []string{"en", "EN", "da", "de", "fr", "es", "ja", "zh"}

	for _, code := range tests {
		lang, ok := LanguageByCode(code)
		if !ok {
			t.Errorf("LanguageByCode(%q) not found", code)

			continue
		}

		if lang.Code == "" {
			t.Errorf("LanguageByCode(%q).Code is empty", code)
		}

		if lang.NativeName == "" {
			t.Errorf("LanguageByCode(%q).NativeName is empty", code)
		}
	}
}

func TestLanguageCodeString(t *testing.T) {
	t.Parallel()

	code := LanguageEN
	if code.String() != "en" {
		t.Errorf("LanguageEN.String() = %q, want %q", code.String(), "en")
	}
}

// Currency tests

func TestAllCurrencies(t *testing.T) {
	t.Parallel()

	currencies := AllCurrencies()
	if len(currencies) == 0 {
		t.Error("Expected currencies, got none")
	}

	if len(currencies) < 100 {
		t.Errorf("Expected at least 100 currencies, got %d", len(currencies))
	}
}

func TestCurrencyByCode(t *testing.T) {
	t.Parallel()

	tests := []string{"USD", "usd", "EUR", "GBP", "JPY", "CNY"}

	for _, code := range tests {
		curr, ok := CurrencyByCode(code)
		if !ok {
			t.Errorf("CurrencyByCode(%q) not found", code)

			continue
		}

		if curr.Code == "" {
			t.Errorf("CurrencyByCode(%q).Code is empty", code)
		}

		if curr.Symbol == "" {
			t.Errorf("CurrencyByCode(%q).Symbol is empty", code)
		}

		if curr.NativeName == "" {
			t.Errorf("CurrencyByCode(%q).NativeName is empty", code)
		}
	}
}

func TestCurrencyCodeString(t *testing.T) {
	t.Parallel()

	code := CurrencyUSD
	if code.String() != "USD" {
		t.Errorf("CurrencyUSD.String() = %q, want %q", code.String(), "USD")
	}
}

// Continent tests

func TestAllContinents(t *testing.T) {
	t.Parallel()

	continents := AllContinents()
	if len(continents) != 7 {
		t.Errorf("Expected 7 continents, got %d", len(continents))
	}
}

func TestContinentByCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		code string
		name string
	}{
		{"EU", "Europe"},
		{"eu", "Europe"},
		{"NA", "North America"},
		{"AS", "Asia"},
		{"AF", "Africa"},
	}

	for _, tt := range tests {
		cont, ok := ContinentByCode(tt.code)
		if !ok {
			t.Errorf("ContinentByCode(%q) not found", tt.code)

			continue
		}

		if cont.Name != tt.name {
			t.Errorf("ContinentByCode(%q).Name = %q, want %q", tt.code, cont.Name, tt.name)
		}
	}
}

// Flag tests

func TestCountryHasFlag(t *testing.T) {
	t.Parallel()

	tests := []string{"US", "GB", "FR", "DE", "JP"}

	for _, code := range tests {
		flag, ok := GetFlag(code)
		if !ok || flag == "" {
			t.Errorf("No flag for %s", code)

			continue
		}

		if !strings.HasPrefix(flag, "<svg") {
			t.Errorf("Flag for %s doesn't start with <svg: %s...", code, flag[:min(50, len(flag))])
		}
	}
}

func TestAllCountriesHaveFlags(t *testing.T) {
	t.Parallel()

	countries := AllCountries()
	missingFlags := 0

	for _, c := range countries {
		if _, ok := GetFlag(c.Alpha2Code); !ok {
			missingFlags++

			t.Logf("Missing flag for %s (%s)", c.NativeName, c.Alpha2Code)
		}
	}

	if missingFlags > 0 {
		t.Errorf("%d countries are missing flags", missingFlags)
	}
}

// Translation tests

func TestCountryTranslation(t *testing.T) {
	t.Parallel()

	translations := i18n.CountriesTranslations["en"]
	if translations == nil {
		t.Skip("No English country translations available")
	}

	name := translations["US"]
	if name == "" {
		t.Error("Expected US country name in English")
	}
}

func TestLanguageTranslation(t *testing.T) {
	t.Parallel()

	translations := i18n.LanguagesTranslations["en"]
	if translations == nil {
		t.Skip("No English language translations available")
	}

	name := translations["en"]
	if name == "" {
		t.Error("Expected English language name in English")
	}
}

func TestCurrencyTranslation(t *testing.T) {
	t.Parallel()

	translations := i18n.CurrenciesTranslations["en"]
	if translations == nil {
		t.Skip("No English currency translations available")
	}

	name := translations["USD"]
	if name == "" {
		t.Error("Expected USD currency name in English")
	}
}

func TestContinentTranslation(t *testing.T) {
	t.Parallel()

	translations := i18n.ContinentsTranslations["en"]
	if translations == nil {
		t.Skip("No English continent translations available")
	}

	name := translations["EU"]
	if name == "" {
		t.Error("Expected EU continent name in English")
	}
}
