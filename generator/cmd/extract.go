package cmd

import (
	"github.com/infobits-io/infobits-intl/generator/internal/extract"
	"github.com/spf13/cobra"
)

var (
	extractSourceDir string
	extractOutputDir string
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract data from existing Dart enum files",
	Long:  `Parses existing Dart enum files and extracts data to JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := extract.ExtractAll(extractSourceDir, extractOutputDir); err != nil {
			exitWithError("Extraction failed", err)
		}
	},
}

func init() {
	extractCmd.Flags().StringVarP(&extractSourceDir, "source", "s", "../lib/src", "Source directory containing Dart files")
	extractCmd.Flags().StringVarP(&extractOutputDir, "output", "o", "../data/core", "Output directory for JSON files")
}
