package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/infobits-io/infobits-intl/generator/internal/template"
)

// GoGenerator generates Go code.
type GoGenerator struct {
	templateDir string
	engine      *template.Engine
}

// NewGoGenerator creates a new Go generator.
func NewGoGenerator(templateDir string) *GoGenerator {
	return &GoGenerator{
		templateDir: templateDir,
		engine:      template.NewEngine(templateDir),
	}
}

// Name returns the generator name.
func (g *GoGenerator) Name() string {
	return "go"
}

// Generate generates Go code.
func (g *GoGenerator) Generate(intlData *data.IntlData, flags map[string]string, translations Translations, cfg *config.TargetConfig) error {
	// Load templates
	if err := g.engine.LoadTemplates("go"); err != nil {
		return fmt.Errorf("loading go templates: %w", err)
	}

	outputDir := cfg.Output

	// Create output directories
	if err := os.MkdirAll(filepath.Join(outputDir, "i18n"), 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Generate country.go
	if err := g.generateCountries(intlData, outputDir); err != nil {
		return fmt.Errorf("generating countries: %w", err)
	}

	// Generate language.go
	if err := g.generateLanguages(intlData, outputDir); err != nil {
		return fmt.Errorf("generating languages: %w", err)
	}

	// Generate currency.go
	if err := g.generateCurrencies(intlData, outputDir); err != nil {
		return fmt.Errorf("generating currencies: %w", err)
	}

	// Generate continent.go
	if err := g.generateContinents(intlData, outputDir); err != nil {
		return fmt.Errorf("generating continents: %w", err)
	}

	// Generate flags.go
	if err := g.generateFlags(flags, outputDir); err != nil {
		return fmt.Errorf("generating flags: %w", err)
	}

	// Generate translations
	if err := g.generateTranslations(translations, outputDir); err != nil {
		return fmt.Errorf("generating translations: %w", err)
	}

	// Generate go.mod
	if err := g.generateGoMod(cfg.ModuleName, outputDir); err != nil {
		return fmt.Errorf("generating go.mod: %w", err)
	}

	fmt.Printf("Go: Generated code in %s\n", outputDir)

	return nil
}

func (g *GoGenerator) generateCountries(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("go", "country.go.tmpl"),
		filepath.Join(outputDir, "country.go"),
		intlData.Countries,
	)
}

func (g *GoGenerator) generateLanguages(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("go", "language.go.tmpl"),
		filepath.Join(outputDir, "language.go"),
		intlData.Languages,
	)
}

func (g *GoGenerator) generateCurrencies(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("go", "currency.go.tmpl"),
		filepath.Join(outputDir, "currency.go"),
		intlData.Currencies,
	)
}

func (g *GoGenerator) generateContinents(intlData *data.IntlData, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("go", "continent.go.tmpl"),
		filepath.Join(outputDir, "continent.go"),
		intlData.Continents,
	)
}

func (g *GoGenerator) generateFlags(flags map[string]string, outputDir string) error {
	return g.engine.ExecuteToFile(
		filepath.Join("go", "flags.go.tmpl"),
		filepath.Join(outputDir, "flags.go"),
		flags,
	)
}

func (g *GoGenerator) generateTranslations(translations Translations, outputDir string) error {
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

func (g *GoGenerator) generateTranslationFile(name string, translations map[string]map[string]string, outputDir string) error {
	tmplData := struct {
		Name         string
		Translations map[string]map[string]string
	}{
		Name:         name,
		Translations: translations,
	}

	return g.engine.ExecuteToFile(
		filepath.Join("go", "translations.go.tmpl"),
		filepath.Join(outputDir, name+".go"),
		tmplData,
	)
}

func (g *GoGenerator) generateGoMod(moduleName, outputDir string) error {
	content := fmt.Sprintf("module %s\n\ngo 1.21\n", moduleName)

	return os.WriteFile(filepath.Join(outputDir, "go.mod"), []byte(content), 0o644)
}
