package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadAll loads all data from the data directory.
func LoadAll(dataDir string) (*IntlData, error) {
	coreDir := filepath.Join(dataDir, "core")

	countries, err := LoadCountries(filepath.Join(coreDir, "countries.json"))
	if err != nil {
		return nil, fmt.Errorf("loading countries: %w", err)
	}

	languages, err := LoadLanguages(filepath.Join(coreDir, "languages.json"))
	if err != nil {
		return nil, fmt.Errorf("loading languages: %w", err)
	}

	currencies, err := LoadCurrencies(filepath.Join(coreDir, "currencies.json"))
	if err != nil {
		return nil, fmt.Errorf("loading currencies: %w", err)
	}

	continents, err := LoadContinents(filepath.Join(coreDir, "continents.json"))
	if err != nil {
		return nil, fmt.Errorf("loading continents: %w", err)
	}

	return &IntlData{
		Countries:  countries,
		Languages:  languages,
		Currencies: currencies,
		Continents: continents,
	}, nil
}

// LoadCountries loads countries from JSON file.
func LoadCountries(path string) ([]Country, error) {
	var file CountriesFile
	if err := loadJSON(path, &file); err != nil {
		return nil, err
	}

	return file.Countries, nil
}

// LoadLanguages loads languages from JSON file.
func LoadLanguages(path string) ([]Language, error) {
	var file LanguagesFile
	if err := loadJSON(path, &file); err != nil {
		return nil, err
	}

	return file.Languages, nil
}

// LoadCurrencies loads currencies from JSON file.
func LoadCurrencies(path string) ([]Currency, error) {
	var file CurrenciesFile
	if err := loadJSON(path, &file); err != nil {
		return nil, err
	}

	return file.Currencies, nil
}

// LoadContinents loads continents from JSON file.
func LoadContinents(path string) ([]Continent, error) {
	var file ContinentsFile
	if err := loadJSON(path, &file); err != nil {
		return nil, err
	}

	return file.Continents, nil
}

// LoadTranslations loads translation files from a directory.
func LoadTranslations(dir string) (map[string]map[string]string, error) {
	translations := make(map[string]map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", dir, err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		locale := file.Name()[:len(file.Name())-5] // Remove .json extension

		var data map[string]string
		if err := loadJSON(filepath.Join(dir, file.Name()), &data); err != nil {
			return nil, fmt.Errorf("loading %s: %w", file.Name(), err)
		}

		translations[locale] = data
	}

	return translations, nil
}

// LoadCurrencyTranslations loads currency translation files (which have nested structure).
func LoadCurrencyTranslations(dir string) (map[string]map[string]string, error) {
	translations := make(map[string]map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", dir, err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		locale := file.Name()[:len(file.Name())-5] // Remove .json extension

		// Try nested format first
		var nestedData map[string]struct {
			Name       string `json:"name"`
			PluralName string `json:"pluralName"`
		}
		if err := loadJSON(filepath.Join(dir, file.Name()), &nestedData); err == nil && len(nestedData) > 0 {
			// Check if first entry has nested structure
			flat := make(map[string]string)
			for code, trans := range nestedData {
				flat[code] = trans.Name
			}

			translations[locale] = flat
		} else {
			// Fall back to flat format
			var flatData map[string]string
			if err := loadJSON(filepath.Join(dir, file.Name()), &flatData); err != nil {
				return nil, fmt.Errorf("loading %s: %w", file.Name(), err)
			}

			translations[locale] = flatData
		}
	}

	return translations, nil
}

// LoadFlags loads SVG flag content from a directory.
func LoadFlags(dir string) (map[string]string, error) {
	flags := make(map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", dir, err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".svg" {
			continue
		}

		code := file.Name()[:len(file.Name())-4] // Remove .svg extension
		code = toUpperCase(code)                 // Convert to uppercase (e.g., "af" -> "AF")

		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", file.Name(), err)
		}

		// Minify SVG
		minified := minifySVG(string(content))
		flags[code] = minified
	}

	return flags, nil
}

// minifySVG removes unnecessary whitespace and comments from SVG content.
func minifySVG(svg string) string {
	var result strings.Builder

	inTag := false
	inComment := false
	lastWasSpace := false

	for i := 0; i < len(svg); i++ {
		c := svg[i]

		// Check for comment start
		if i+3 < len(svg) && svg[i:i+4] == "<!--" {
			inComment = true
			i += 3

			continue
		}

		// Check for comment end
		if inComment {
			if i+2 < len(svg) && svg[i:i+3] == "-->" {
				inComment = false
				i += 2
			}

			continue
		}

		// Track tag state
		if c == '<' {
			inTag = true
			lastWasSpace = false

			result.WriteByte(c)

			continue
		}

		if c == '>' {
			inTag = false
			lastWasSpace = false

			result.WriteByte(c)

			continue
		}

		// Handle whitespace
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if inTag && !lastWasSpace {
				result.WriteByte(' ')

				lastWasSpace = true
			}

			continue
		}

		lastWasSpace = false

		result.WriteByte(c)
	}

	return strings.TrimSpace(result.String())
}

func loadJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}

	return nil
}

func toUpperCase(s string) string {
	result := make([]byte, len(s))
	for i := range len(s) {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			result[i] = c - 32
		} else {
			result[i] = c
		}
	}

	return string(result)
}
