package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/infobits-io/infobits-intl/generator/internal/template"
)

// DartGenerator generates Dart/Flutter code.
type DartGenerator struct {
	templateDir string
	engine      *template.Engine
}

// NewDartGenerator creates a new Dart generator.
func NewDartGenerator(templateDir string) *DartGenerator {
	return &DartGenerator{
		templateDir: templateDir,
		engine:      template.NewEngine(templateDir),
	}
}

// Name returns the generator name.
func (g *DartGenerator) Name() string {
	return "dart"
}

// Generate generates Dart code.
func (g *DartGenerator) Generate(intlData *data.IntlData, flags map[string]string, translations Translations, cfg *config.TargetConfig) error {
	// Load templates
	if err := g.engine.LoadTemplates("dart"); err != nil {
		return fmt.Errorf("loading dart templates: %w", err)
	}

	outputDir := cfg.Output

	// Create output directories
	if err := os.MkdirAll(filepath.Join(outputDir, "i18n"), 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Generate countries.dart
	if err := g.generateCountries(intlData, flags, outputDir); err != nil {
		return fmt.Errorf("generating countries: %w", err)
	}

	// Generate languages.dart
	if err := g.generateLanguages(intlData, flags, outputDir); err != nil {
		return fmt.Errorf("generating languages: %w", err)
	}

	// Generate currencies.dart
	if err := g.generateCurrencies(intlData, outputDir); err != nil {
		return fmt.Errorf("generating currencies: %w", err)
	}

	// Generate continents.dart
	if err := g.generateContinents(intlData, outputDir); err != nil {
		return fmt.Errorf("generating continents: %w", err)
	}

	// Generate translation delegates
	if err := g.generateTranslations(translations, outputDir); err != nil {
		return fmt.Errorf("generating translations: %w", err)
	}

	// Generate flags.g.dart
	if err := g.generateFlags(flags, outputDir); err != nil {
		return fmt.Errorf("generating flags: %w", err)
	}

	fmt.Printf("Dart: Generated code in %s\n", outputDir)

	return nil
}

func (g *DartGenerator) generateCountries(intlData *data.IntlData, flags map[string]string, outputDir string) error {
	tmplData := struct {
		Countries []data.Country
		Flags     map[string]string
	}{
		Countries: intlData.Countries,
		Flags:     flags,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("dart", "countries.dart.tmpl"),
		filepath.Join(outputDir, "countries.g.dart"),
		tmplData,
	)
}

func (g *DartGenerator) generateLanguages(intlData *data.IntlData, flags map[string]string, outputDir string) error {
	tmplData := struct {
		Languages []data.Language
		Flags     map[string]string
	}{
		Languages: intlData.Languages,
		Flags:     flags,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("dart", "languages.dart.tmpl"),
		filepath.Join(outputDir, "languages.g.dart"),
		tmplData,
	)
}

func (g *DartGenerator) generateCurrencies(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("dart", "currencies.dart.tmpl"),
		filepath.Join(outputDir, "currencies.g.dart"),
		intlData.Currencies,
	)
}

func (g *DartGenerator) generateContinents(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("dart", "continents.dart.tmpl"),
		filepath.Join(outputDir, "continents.g.dart"),
		intlData.Continents,
	)
}

func (g *DartGenerator) generateTranslations(translations Translations, outputDir string) error {
	i18nDir := filepath.Join(outputDir, "i18n")

	// Countries translations
	if err := g.generateTranslationDelegate("countries", translations.Countries, i18nDir); err != nil {
		return err
	}

	// Languages translations
	if err := g.generateTranslationDelegate("languages", translations.Languages, i18nDir); err != nil {
		return err
	}

	// Currencies translations
	if err := g.generateTranslationDelegate("currencies", translations.Currencies, i18nDir); err != nil {
		return err
	}

	// Continents translations
	if err := g.generateTranslationDelegate("continents", translations.Continents, i18nDir); err != nil {
		return err
	}

	// Capitals translations
	if err := g.generateTranslationDelegate("capitals", translations.Capitals, i18nDir); err != nil {
		return err
	}

	return nil
}

func (g *DartGenerator) generateTranslationDelegate(name string, translations map[string]map[string]string, outputDir string) error {
	caser := cases.Title(language.English)
	tmplData := struct {
		Name         string
		ClassName    string
		Translations map[string]map[string]string
	}{
		Name:         name,
		ClassName:    caser.String(name) + "TranslationsDelegate",
		Translations: translations,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("dart", "translations.dart.tmpl"),
		filepath.Join(outputDir, name+".g.dart"),
		tmplData,
	)
}

func (g *DartGenerator) generateFlags(flags map[string]string, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("dart", "flags.dart.tmpl"),
		filepath.Join(outputDir, "flags.g.dart"),
		flags,
	)
}
