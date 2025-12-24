package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/spf13/cobra"
)

var validateDataDir string

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate JSON data files",
	Long:  `Validates the JSON data files for correctness and consistency.`,
	Run: func(cmd *cobra.Command, args []string) {
		errors := []string{}
		warnings := []string{}

		// Load all data
		intlData, err := data.LoadAll(validateDataDir)
		if err != nil {
			exitWithError("Loading data", err)
		}

		// Build lookup maps
		languageIDs := make(map[string]bool)
		for _, lang := range intlData.Languages {
			if languageIDs[lang.ID] {
				errors = append(errors, "Duplicate language ID: "+lang.ID)
			}
			languageIDs[lang.ID] = true
		}

		currencyIDs := make(map[string]bool)
		for _, curr := range intlData.Currencies {
			if currencyIDs[curr.ID] {
				errors = append(errors, "Duplicate currency ID: "+curr.ID)
			}
			currencyIDs[curr.ID] = true
		}

		continentIDs := make(map[string]bool)
		for _, cont := range intlData.Continents {
			if continentIDs[cont.ID] {
				errors = append(errors, "Duplicate continent ID: "+cont.ID)
			}
			continentIDs[cont.ID] = true
		}

		countryIDs := make(map[string]bool)
		alpha2Codes := make(map[string]bool)
		alpha3Codes := make(map[string]bool)
		numericCodes := make(map[int]bool)

		for _, country := range intlData.Countries {
			// Check for duplicate IDs
			if countryIDs[country.ID] {
				errors = append(errors, "Duplicate country ID: "+country.ID)
			}
			countryIDs[country.ID] = true

			// Check for duplicate codes
			if alpha2Codes[country.Alpha2Code] {
				errors = append(errors, "Duplicate alpha2 code: "+country.Alpha2Code)
			}
			alpha2Codes[country.Alpha2Code] = true

			if alpha3Codes[country.Alpha3Code] {
				errors = append(errors, "Duplicate alpha3 code: "+country.Alpha3Code)
			}
			alpha3Codes[country.Alpha3Code] = true

			if numericCodes[country.NumericCode] {
				errors = append(errors, fmt.Sprintf("Duplicate numeric code: %d (country: %s)", country.NumericCode, country.ID))
			}
			numericCodes[country.NumericCode] = true

			// Validate required fields
			if country.Alpha2Code == "" {
				errors = append(errors, fmt.Sprintf("Country %s missing alpha2Code", country.ID))
			} else if len(country.Alpha2Code) != 2 {
				errors = append(errors, fmt.Sprintf("Country %s has invalid alpha2Code length: %s", country.ID, country.Alpha2Code))
			}

			if country.Alpha3Code == "" {
				errors = append(errors, fmt.Sprintf("Country %s missing alpha3Code", country.ID))
			} else if len(country.Alpha3Code) != 3 {
				errors = append(errors, fmt.Sprintf("Country %s has invalid alpha3Code length: %s", country.ID, country.Alpha3Code))
			}

			if country.NativeName == "" {
				errors = append(errors, fmt.Sprintf("Country %s missing nativeName", country.ID))
			}

			// Validate references
			if !languageIDs[country.MainLanguage] {
				errors = append(errors, fmt.Sprintf("Country %s references unknown mainLanguage: %s", country.ID, country.MainLanguage))
			}

			for _, langID := range country.Languages {
				if !languageIDs[langID] {
					errors = append(errors, fmt.Sprintf("Country %s references unknown language: %s", country.ID, langID))
				}
			}

			if !continentIDs[country.Continent] {
				errors = append(errors, fmt.Sprintf("Country %s references unknown continent: %s", country.ID, country.Continent))
			}

			if !currencyIDs[country.Currency] {
				errors = append(errors, fmt.Sprintf("Country %s references unknown currency: %s", country.ID, country.Currency))
			}

			// Check for missing capital (warning)
			if country.Capital == "" {
				warnings = append(warnings, fmt.Sprintf("Country %s has no capital", country.ID))
			}
		}

		// Validate flags exist
		flagsDir := filepath.Join(validateDataDir, "assets", "flags")
		for _, country := range intlData.Countries {
			flagPath := filepath.Join(flagsDir, toLowerCase(country.Alpha2Code)+".svg")
			if _, err := os.Stat(flagPath); os.IsNotExist(err) {
				warnings = append(warnings, fmt.Sprintf("Missing flag SVG for country %s (%s)", country.ID, country.Alpha2Code))
			}
		}

		// Validate translations reference valid codes
		i18nDir := filepath.Join(validateDataDir, "i18n")
		validateTranslationDir(filepath.Join(i18nDir, "countries"), alpha2Codes, "country", &errors, &warnings)
		validateTranslationDir(filepath.Join(i18nDir, "continents"), continentCodeMap(intlData.Continents), "continent", &errors, &warnings)
		validateLanguageTranslationDir(filepath.Join(i18nDir, "languages"), intlData.Languages, &errors, &warnings)
		validateCurrencyTranslationDir(filepath.Join(i18nDir, "currencies"), intlData.Currencies, &errors, &warnings)

		// Print results
		fmt.Printf("\nValidation Results:\n")
		fmt.Printf("==================\n")
		fmt.Printf("Countries:  %d\n", len(intlData.Countries))
		fmt.Printf("Languages:  %d\n", len(intlData.Languages))
		fmt.Printf("Currencies: %d\n", len(intlData.Currencies))
		fmt.Printf("Continents: %d\n", len(intlData.Continents))
		fmt.Println()

		if len(warnings) > 0 {
			fmt.Printf("Warnings (%d):\n", len(warnings))
			for _, w := range warnings {
				fmt.Printf("  ⚠ %s\n", w)
			}
			fmt.Println()
		}

		if len(errors) > 0 {
			fmt.Printf("Errors (%d):\n", len(errors))
			for _, e := range errors {
				fmt.Printf("  ✗ %s\n", e)
			}
			os.Exit(1)
		}

		fmt.Println("✓ Validation passed!")
	},
}

func init() {
	validateCmd.Flags().StringVarP(&validateDataDir, "data", "d", "../data", "Path to data directory")
}

func toLowerCase(s string) string {
	result := make([]byte, len(s))
	for i := range len(s) {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32
		} else {
			result[i] = c
		}
	}

	return string(result)
}

func continentCodeMap(continents []data.Continent) map[string]bool {
	m := make(map[string]bool)
	for _, c := range continents {
		m[c.Code] = true
	}

	return m
}

func validateTranslationDir(dir string, validCodes map[string]bool, entityType string, errors, warnings *[]string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		*warnings = append(*warnings, fmt.Sprintf("Could not read %s translations directory: %v", entityType, err))

		return
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		var translations map[string]string

		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Could not read %s: %v", file.Name(), err))

			continue
		}

		if err := parseJSON(content, &translations); err != nil {
			*errors = append(*errors, fmt.Sprintf("Invalid JSON in %s/%s: %v", entityType, file.Name(), err))

			continue
		}

		locale := file.Name()[:len(file.Name())-5]

		for code := range translations {
			if !validCodes[code] {
				*warnings = append(*warnings, fmt.Sprintf("Translation %s/%s has unknown %s code: %s", entityType, locale, entityType, code))
			}
		}
	}
}

func validateLanguageTranslationDir(dir string, languages []data.Language, errors, warnings *[]string) {
	validCodes := make(map[string]bool)
	for _, lang := range languages {
		validCodes[lang.Code] = true
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		*warnings = append(*warnings, fmt.Sprintf("Could not read languages translations directory: %v", err))

		return
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		var translations map[string]string

		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Could not read %s: %v", file.Name(), err))

			continue
		}

		if err := parseJSON(content, &translations); err != nil {
			*errors = append(*errors, fmt.Sprintf("Invalid JSON in languages/%s: %v", file.Name(), err))

			continue
		}

		locale := file.Name()[:len(file.Name())-5]

		for code := range translations {
			if !validCodes[code] {
				*warnings = append(*warnings, fmt.Sprintf("Translation languages/%s has unknown language code: %s", locale, code))
			}
		}
	}
}

func validateCurrencyTranslationDir(dir string, currencies []data.Currency, errors, warnings *[]string) {
	validCodes := make(map[string]bool)
	for _, curr := range currencies {
		validCodes[curr.Code] = true
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		*warnings = append(*warnings, fmt.Sprintf("Could not read currencies translations directory: %v", err))

		return
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Could not read %s: %v", file.Name(), err))

			continue
		}

		// Try nested format first
		var nestedData map[string]struct {
			Name       string `json:"name"`
			PluralName string `json:"pluralName"`
		}
		if err := parseJSON(content, &nestedData); err == nil && len(nestedData) > 0 {
			locale := file.Name()[:len(file.Name())-5]

			for code := range nestedData {
				if !validCodes[code] {
					*warnings = append(*warnings, fmt.Sprintf("Translation currencies/%s has unknown currency code: %s", locale, code))
				}
			}
		} else {
			// Try flat format
			var flatData map[string]string
			if err := parseJSON(content, &flatData); err != nil {
				*errors = append(*errors, fmt.Sprintf("Invalid JSON in currencies/%s: %v", file.Name(), err))

				continue
			}

			locale := file.Name()[:len(file.Name())-5]

			for code := range flatData {
				if !validCodes[code] {
					*warnings = append(*warnings, fmt.Sprintf("Translation currencies/%s has unknown currency code: %s", locale, code))
				}
			}
		}
	}
}

func parseJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
