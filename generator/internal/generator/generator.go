package generator

import (
	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
)

// Generator is the interface for code generators.
type Generator interface {
	// Name returns the name of the generator
	Name() string

	// Generate generates code for the target
	Generate(intlData *data.IntlData, flags map[string]string, translations Translations, cfg *config.TargetConfig) error
}

// Translations holds all translation data.
type Translations struct {
	Countries  map[string]map[string]string // locale -> code -> name
	Languages  map[string]map[string]string
	Currencies map[string]map[string]string
	Continents map[string]map[string]string
	Capitals   map[string]map[string]string
}

// LoadTranslations loads all translations from the i18n directory.
func LoadTranslations(dataDir string) (Translations, error) {
	var (
		t   Translations
		err error
	)

	t.Countries, err = data.LoadTranslations(dataDir + "/i18n/countries")
	if err != nil {
		return t, err
	}

	t.Languages, err = data.LoadTranslations(dataDir + "/i18n/languages")
	if err != nil {
		return t, err
	}

	// Currencies have nested structure
	t.Currencies, err = data.LoadCurrencyTranslations(dataDir + "/i18n/currencies")
	if err != nil {
		return t, err
	}

	t.Continents, err = data.LoadTranslations(dataDir + "/i18n/continents")
	if err != nil {
		return t, err
	}

	t.Capitals, err = data.LoadTranslations(dataDir + "/i18n/capitals")
	if err != nil {
		return t, err
	}

	return t, nil
}
