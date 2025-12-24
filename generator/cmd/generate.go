package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/infobits-io/infobits-intl/generator/internal/config"
	"github.com/infobits-io/infobits-intl/generator/internal/data"
	"github.com/infobits-io/infobits-intl/generator/internal/generator"
	"github.com/spf13/cobra"
)

var (
	generateTarget      string
	generateAll         bool
	generateConfigPath  string
	generateDataDir     string
	generateTemplateDir string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for target frameworks",
	Long:  `Generates Dart, TypeScript, and Go code from JSON data files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load or create config
		var cfg *config.Config
		var err error

		if generateConfigPath != "" {
			cfg, err = config.Load(generateConfigPath)
			if err != nil {
				exitWithError("Loading config", err)
			}
		} else {
			cfg = config.Default()
		}

		// Override data directory if specified
		if generateDataDir != "" {
			cfg.Data.Source = generateDataDir
			cfg.Flags.Source = generateDataDir + "/assets/flags"
		}

		// Load data
		intlData, err := data.LoadAll(cfg.Data.Source)
		if err != nil {
			exitWithError("Loading data", err)
		}

		// Load flags
		flags, err := data.LoadFlags(cfg.Flags.Source)
		if err != nil {
			exitWithError("Loading flags", err)
		}

		// Load translations
		translations, err := generator.LoadTranslations(cfg.Data.Source)
		if err != nil {
			exitWithError("Loading translations", err)
		}

		fmt.Printf("Loaded %d countries, %d languages, %d currencies, %d continents, %d flags\n",
			len(intlData.Countries), len(intlData.Languages), len(intlData.Currencies),
			len(intlData.Continents), len(flags))

		// Get template directory
		templateDir := generateTemplateDir
		if templateDir == "" {
			// Fall back to executable directory
			execPath, err := os.Executable()
			if err != nil {
				exitWithError("Getting executable path", err)
			}
			templateDir = filepath.Join(filepath.Dir(execPath), "templates")
		}

		// Generate for requested targets
		if generateAll || generateTarget == "" {
			// Generate all enabled targets
			if cfg.Targets.Dart != nil && cfg.Targets.Dart.Enabled {
				gen := generator.NewDartGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.Dart); err != nil {
					exitWithError("Generating Dart", err)
				}
			}

			if cfg.Targets.TypeScript != nil && cfg.Targets.TypeScript.Enabled {
				gen := generator.NewTypeScriptGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.TypeScript); err != nil {
					exitWithError("Generating TypeScript", err)
				}
			}

			if cfg.Targets.Go != nil && cfg.Targets.Go.Enabled {
				gen := generator.NewGoGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.Go); err != nil {
					exitWithError("Generating Go", err)
				}
			}
		} else {
			// Generate specific target
			switch generateTarget {
			case "dart":
				if cfg.Targets.Dart == nil {
					exitWithError("Dart target not configured", nil)
				}
				gen := generator.NewDartGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.Dart); err != nil {
					exitWithError("Generating Dart", err)
				}
			case "typescript", "ts":
				if cfg.Targets.TypeScript == nil {
					exitWithError("TypeScript target not configured", nil)
				}
				gen := generator.NewTypeScriptGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.TypeScript); err != nil {
					exitWithError("Generating TypeScript", err)
				}
			case "go":
				if cfg.Targets.Go == nil {
					exitWithError("Go target not configured", nil)
				}
				gen := generator.NewGoGenerator(templateDir)
				if err := gen.Generate(intlData, flags, translations, cfg.Targets.Go); err != nil {
					exitWithError("Generating Go", err)
				}
			default:
				exitWithError("Unknown target: "+generateTarget, nil)
			}
		}

		fmt.Println("Generation complete!")
	},
}

func init() {
	generateCmd.Flags().StringVarP(&generateTarget, "target", "t", "", "Target to generate (dart, typescript, go)")
	generateCmd.Flags().BoolVarP(&generateAll, "all", "a", false, "Generate all targets")
	generateCmd.Flags().StringVarP(&generateConfigPath, "config", "c", "", "Path to generator.yaml config file")
	generateCmd.Flags().StringVarP(&generateDataDir, "data", "d", "../data", "Path to data directory")
	generateCmd.Flags().StringVar(&generateTemplateDir, "templates", "", "Path to templates directory (defaults to executable directory)")
}
