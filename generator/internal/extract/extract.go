package extract

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/infobits-io/infobits-intl/generator/internal/data"
)

// ExtractAll extracts all data from Dart source files.
func ExtractAll(srcDir, outputDir string) error {
	// Extract countries
	countries, err := ExtractCountries(filepath.Join(srcDir, "countries.dart"))
	if err != nil {
		return fmt.Errorf("extracting countries: %w", err)
	}

	// Extract languages
	languages, err := ExtractLanguages(filepath.Join(srcDir, "languages.dart"))
	if err != nil {
		return fmt.Errorf("extracting languages: %w", err)
	}

	// Extract currencies
	currencies, err := ExtractCurrencies(filepath.Join(srcDir, "currencies.dart"))
	if err != nil {
		return fmt.Errorf("extracting currencies: %w", err)
	}

	// Extract continents
	continents, err := ExtractContinents(filepath.Join(srcDir, "continents.dart"))
	if err != nil {
		return fmt.Errorf("extracting continents: %w", err)
	}

	// Write output files
	if err := writeJSON(filepath.Join(outputDir, "countries.json"), data.CountriesFile{Countries: countries}); err != nil {
		return err
	}

	if err := writeJSON(filepath.Join(outputDir, "languages.json"), data.LanguagesFile{Languages: languages}); err != nil {
		return err
	}

	if err := writeJSON(filepath.Join(outputDir, "currencies.json"), data.CurrenciesFile{Currencies: currencies}); err != nil {
		return err
	}

	if err := writeJSON(filepath.Join(outputDir, "continents.json"), data.ContinentsFile{Continents: continents}); err != nil {
		return err
	}

	fmt.Printf("Extracted %d countries, %d languages, %d currencies, %d continents\n",
		len(countries), len(languages), len(currencies), len(continents))

	return nil
}

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}

	fmt.Printf("Written: %s\n", path)

	return nil
}

// ExtractCountries extracts country data from countries.dart.
func ExtractCountries(path string) ([]data.Country, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Pattern to match enum entries like:
	// afghanistan(
	//   alpha2Code: "AF",
	//   ...
	// ),
	enumPattern := regexp.MustCompile(
		`(?s)(\w+)\(\s*alpha2Code:\s*"([^"]+)",\s*alpha3Code:\s*"([^"]+)",\s*numericCode:\s*(\d+),` +
			`\s*nativeName:\s*"([^"]*)",\s*capital:\s*"([^"]*)",\s*mainLanguage:\s*Language\.(\w+),` +
			`\s*languages:\s*\[([^\]]*)\],\s*tld:\s*"([^"]+)",\s*callingCode:\s*(\d+),` +
			`\s*continent:\s*Continent\.(\w+),\s*currency:\s*Currency\.(\w+),\s*\)`,
	)

	matches := enumPattern.FindAllStringSubmatch(string(content), -1)

	countries := make([]data.Country, 0, len(matches))

	for _, match := range matches {
		numericCode, err := strconv.Atoi(match[4])
		if err != nil {
			return nil, fmt.Errorf("parsing numeric code for %s: %w", match[1], err)
		}

		callingCode, err := strconv.Atoi(match[10])
		if err != nil {
			return nil, fmt.Errorf("parsing calling code for %s: %w", match[1], err)
		}

		// Parse languages list
		languagesStr := match[8]

		var languages []string

		langPattern := regexp.MustCompile(`Language\.(\w+)`)

		langMatches := langPattern.FindAllStringSubmatch(languagesStr, -1)
		for _, lm := range langMatches {
			languages = append(languages, lm[1])
		}

		countries = append(countries, data.Country{
			ID:           match[1],
			Alpha2Code:   match[2],
			Alpha3Code:   match[3],
			NumericCode:  numericCode,
			NativeName:   match[5],
			Capital:      match[6],
			MainLanguage: match[7],
			Languages:    languages,
			TLD:          match[9],
			CallingCode:  callingCode,
			Continent:    match[11],
			Currency:     match[12],
		})
	}

	return countries, nil
}

// ExtractLanguages extracts language data from languages.dart.
func ExtractLanguages(path string) ([]data.Language, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Pattern for simple language entries (no dialects, no defaultFlagBytes)
	simplePattern := regexp.MustCompile(`(?s)(\w+)\(\s*code:\s*"([^"]+)",\s*nativeName:\s*"([^"]*)",\s*dialects:\s*\[\],?\s*\)`)

	// Pattern for language with defaultFlagBytes
	flagPattern := regexp.MustCompile(
		`(?s)(\w+)\(\s*code:\s*"([^"]+)",\s*nativeName:\s*"([^"]*)",\s*dialects:\s*\[\],?\s*defaultFlagBytes:\s*AssetBytesLoader\("assets/country_flags/(\w+)\.svg"\),?\s*\)`,
	)

	// Pattern for language with dialects
	dialectPattern := regexp.MustCompile(
		`(?s)(\w+)\(\s*code:\s*"([^"]+)",\s*nativeName:\s*"([^"]*)",\s*dialects:\s*\[([\s\S]*?)\],(?:\s*defaultFlagBytes:\s*AssetBytesLoader\("assets/country_flags/(\w+)\.svg"\),?)?\s*\)`,
	)

	languages := make([]data.Language, 0)

	processedIDs := make(map[string]bool)

	// First, find languages with dialects
	dialectMatches := dialectPattern.FindAllStringSubmatch(string(content), -1)
	for _, match := range dialectMatches {
		if match[4] != "" && strings.Contains(match[4], "LanguageDialect") {
			id := match[1]
			if processedIDs[id] {
				continue
			}

			processedIDs[id] = true

			dialects := parseDialects(match[4])

			lang := data.Language{
				ID:         id,
				Code:       match[2],
				NativeName: match[3],
				Dialects:   dialects,
			}
			if match[5] != "" {
				lang.DefaultFlagCode = match[5]
			}

			languages = append(languages, lang)
		}
	}

	// Find languages with defaultFlagBytes but no dialects
	flagMatches := flagPattern.FindAllStringSubmatch(string(content), -1)
	for _, match := range flagMatches {
		id := match[1]
		if processedIDs[id] {
			continue
		}

		processedIDs[id] = true

		languages = append(languages, data.Language{
			ID:              id,
			Code:            match[2],
			NativeName:      match[3],
			DefaultFlagCode: match[4],
		})
	}

	// Find simple languages
	simpleMatches := simplePattern.FindAllStringSubmatch(string(content), -1)
	for _, match := range simpleMatches {
		id := match[1]
		if processedIDs[id] {
			continue
		}

		processedIDs[id] = true

		languages = append(languages, data.Language{
			ID:         id,
			Code:       match[2],
			NativeName: match[3],
		})
	}

	return languages, nil
}

func parseDialects(dialectsStr string) []data.LanguageDialect {
	// Pattern for LanguageDialect entries
	pattern := regexp.MustCompile(
		`LanguageDialect\(\s*id:\s*"([^"]+)",\s*code:\s*"([^"]+)",\s*nativeName:\s*"([^"]*)",\s*flagCode:\s*"([^"]+)",?\s*\)`,
	)

	matches := pattern.FindAllStringSubmatch(dialectsStr, -1)
	dialects := make([]data.LanguageDialect, 0, len(matches))

	for _, match := range matches {
		dialects = append(dialects, data.LanguageDialect{
			ID:         match[1],
			Code:       match[2],
			NativeName: match[3],
			FlagCode:   match[4],
		})
	}

	return dialects
}

// ExtractCurrencies extracts currency data from currencies.dart.
func ExtractCurrencies(path string) ([]data.Currency, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Pattern to match currency entries
	pattern := regexp.MustCompile(`(?s)(\w+)\(\s*code:\s*"([^"]+)",\s*nativeName:\s*"([^"]*)",\s*nativeNamePlural:\s*"([^"]*)",\s*symbol:\s*"([^"]*)",?\s*\)`)

	matches := pattern.FindAllStringSubmatch(string(content), -1)
	currencies := make([]data.Currency, 0, len(matches))

	for _, match := range matches {
		currencies = append(currencies, data.Currency{
			ID:               match[1],
			Code:             match[2],
			NativeName:       match[3],
			NativeNamePlural: match[4],
			Symbol:           match[5],
		})
	}

	return currencies, nil
}

// ExtractContinents extracts continent data from continents.dart.
func ExtractContinents(path string) ([]data.Continent, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Pattern to match continent entries
	pattern := regexp.MustCompile(`(?s)(\w+)\(\s*code:\s*"([^"]+)",\s*name:\s*"([^"]+)",?\s*\)`)

	matches := pattern.FindAllStringSubmatch(string(content), -1)
	continents := make([]data.Continent, 0, len(matches))

	for _, match := range matches {
		continents = append(continents, data.Continent{
			ID:   match[1],
			Code: match[2],
			Name: match[3],
		})
	}

	return continents, nil
}
