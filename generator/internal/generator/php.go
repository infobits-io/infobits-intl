package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/infobits-io/infobits-intl/generator/internal/template"
)

// PHPGenerator generates PHP code.
type PHPGenerator struct {
	templateDir string
	engine      *template.Engine
}

// NewPHPGenerator creates a new PHP generator.
func NewPHPGenerator(templateDir string) *PHPGenerator {
	return &PHPGenerator{
		templateDir: templateDir,
		engine:      template.NewEngine(templateDir),
	}
}

// Name returns the generator name.
func (g *PHPGenerator) Name() string {
	return "php"
}

// Generate generates PHP code.
func (g *PHPGenerator) Generate(intlData *data.IntlData, flags map[string]string, translations Translations, cfg *config.TargetConfig) error {
	// Load templates
	if err := g.engine.LoadTemplates("php"); err != nil {
		return fmt.Errorf("loading php templates: %w", err)
	}

	outputDir := cfg.Output

	// Create output directories
	if err := os.MkdirAll(filepath.Join(outputDir, "I18n"), 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Generate Country.php
	if err := g.generateCountries(intlData, outputDir); err != nil {
		return fmt.Errorf("generating countries: %w", err)
	}

	// Generate Language.php
	if err := g.generateLanguages(intlData, outputDir); err != nil {
		return fmt.Errorf("generating languages: %w", err)
	}

	// Generate Currency.php
	if err := g.generateCurrencies(intlData, outputDir); err != nil {
		return fmt.Errorf("generating currencies: %w", err)
	}

	// Generate Continent.php
	if err := g.generateContinents(intlData, outputDir); err != nil {
		return fmt.Errorf("generating continents: %w", err)
	}

	// Generate Flags.php
	if err := g.generateFlags(flags, outputDir); err != nil {
		return fmt.Errorf("generating flags: %w", err)
	}

	// Generate translations
	if err := g.generateTranslations(translations, outputDir); err != nil {
		return fmt.Errorf("generating translations: %w", err)
	}

	fmt.Printf("PHP: Generated code in %s\n", outputDir)

	return nil
}

func (g *PHPGenerator) generateCountries(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("php", "country.php.tmpl"),
		filepath.Join(outputDir, "Country.php"),
		intlData.Countries,
	)
}

func (g *PHPGenerator) generateLanguages(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("php", "language.php.tmpl"),
		filepath.Join(outputDir, "Language.php"),
		intlData.Languages,
	)
}

func (g *PHPGenerator) generateCurrencies(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("php", "currency.php.tmpl"),
		filepath.Join(outputDir, "Currency.php"),
		intlData.Currencies,
	)
}

func (g *PHPGenerator) generateContinents(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("php", "continent.php.tmpl"),
		filepath.Join(outputDir, "Continent.php"),
		intlData.Continents,
	)
}

func (g *PHPGenerator) generateFlags(flags map[string]string, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("php", "flags.php.tmpl"),
		filepath.Join(outputDir, "Flags.php"),
		flags,
	)
}

func (g *PHPGenerator) generateTranslations(translations Translations, outputDir string) error {
	i18nDir := filepath.Join(outputDir, "I18n")

	if err := g.generateTranslationFile("Countries", "countries", translations.Countries, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("Languages", "languages", translations.Languages, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("Currencies", "currencies", translations.Currencies, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("Continents", "continents", translations.Continents, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("Capitals", "capitals", translations.Capitals, i18nDir); err != nil {
		return err
	}

	return nil
}

func (g *PHPGenerator) generateTranslationFile(className, name string, translations map[string]map[string]string, outputDir string) error {
	tmplData := struct {
		ClassName    string
		Name         string
		Translations map[string]map[string]string
	}{
		ClassName:    className,
		Name:         name,
		Translations: translations,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("php", "translations.php.tmpl"),
		filepath.Join(outputDir, className+"Translations.php"),
		tmplData,
	)
}
