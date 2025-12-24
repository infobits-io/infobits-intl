package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/infobits-io/infobits-intl/generator/internal/template"
)

// TypeScriptGenerator generates TypeScript code.
type TypeScriptGenerator struct {
	templateDir string
	engine      *template.Engine
}

// NewTypeScriptGenerator creates a new TypeScript generator.
func NewTypeScriptGenerator(templateDir string) *TypeScriptGenerator {
	return &TypeScriptGenerator{
		templateDir: templateDir,
		engine:      template.NewEngine(templateDir),
	}
}

// Name returns the generator name.
func (g *TypeScriptGenerator) Name() string {
	return "typescript"
}

// Generate generates TypeScript code.
func (g *TypeScriptGenerator) Generate(intlData *data.IntlData, flags map[string]string, translations Translations, cfg *config.TargetConfig) error {
	// Load templates
	if err := g.engine.LoadTemplates("typescript"); err != nil {
		return fmt.Errorf("loading typescript templates: %w", err)
	}

	outputDir := cfg.Output

	// Create output directories
	if err := os.MkdirAll(filepath.Join(outputDir, "i18n"), 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Generate types.ts
	if err := g.generateTypes(outputDir); err != nil {
		return fmt.Errorf("generating types: %w", err)
	}

	// Generate countries.ts
	if err := g.generateCountries(intlData, outputDir); err != nil {
		return fmt.Errorf("generating countries: %w", err)
	}

	// Generate languages.ts
	if err := g.generateLanguages(intlData, outputDir); err != nil {
		return fmt.Errorf("generating languages: %w", err)
	}

	// Generate currencies.ts
	if err := g.generateCurrencies(intlData, outputDir); err != nil {
		return fmt.Errorf("generating currencies: %w", err)
	}

	// Generate continents.ts
	if err := g.generateContinents(intlData, outputDir); err != nil {
		return fmt.Errorf("generating continents: %w", err)
	}

	// Generate flags.ts
	if err := g.generateFlags(flags, outputDir); err != nil {
		return fmt.Errorf("generating flags: %w", err)
	}

	// Generate translations
	if err := g.generateTranslations(translations, outputDir); err != nil {
		return fmt.Errorf("generating translations: %w", err)
	}

	// Generate index.ts
	if err := g.generateIndex(outputDir); err != nil {
		return fmt.Errorf("generating index: %w", err)
	}

	fmt.Printf("TypeScript: Generated code in %s\n", outputDir)

	return nil
}

func (g *TypeScriptGenerator) generateTypes(outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "types.ts.tmpl"),
		filepath.Join(outputDir, "types.ts"),
		nil,
	)
}

func (g *TypeScriptGenerator) generateCountries(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "countries.ts.tmpl"),
		filepath.Join(outputDir, "countries.ts"),
		intlData.Countries,
	)
}

func (g *TypeScriptGenerator) generateLanguages(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "languages.ts.tmpl"),
		filepath.Join(outputDir, "languages.ts"),
		intlData.Languages,
	)
}

func (g *TypeScriptGenerator) generateCurrencies(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "currencies.ts.tmpl"),
		filepath.Join(outputDir, "currencies.ts"),
		intlData.Currencies,
	)
}

func (g *TypeScriptGenerator) generateContinents(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "continents.ts.tmpl"),
		filepath.Join(outputDir, "continents.ts"),
		intlData.Continents,
	)
}

func (g *TypeScriptGenerator) generateFlags(flags map[string]string, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "flags.ts.tmpl"),
		filepath.Join(outputDir, "flags.ts"),
		flags,
	)
}

func (g *TypeScriptGenerator) generateTranslations(translations Translations, outputDir string) error {
	i18nDir := filepath.Join(outputDir, "i18n")

	// Generate each translation file
	if err := g.generateTranslationFile("countries", translations.Countries, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("languages", translations.Languages, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("currencies", translations.Currencies, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("continents", translations.Continents, i18nDir); err != nil {
		return err
	}

	if err := g.generateTranslationFile("capitals", translations.Capitals, i18nDir); err != nil {
		return err
	}

	return nil
}

func (g *TypeScriptGenerator) generateTranslationFile(name string, translations map[string]map[string]string, outputDir string) error {
	tmplData := struct {
		Name         string
		Translations map[string]map[string]string
	}{
		Name:         name,
		Translations: translations,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "translations.ts.tmpl"),
		filepath.Join(outputDir, name+".ts"),
		tmplData,
	)
}

func (g *TypeScriptGenerator) generateIndex(outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("typescript", "index.ts.tmpl"),
		filepath.Join(outputDir, "index.ts"),
		nil,
	)
}
