package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the root configuration structure.
type Config struct {
	Version string        `yaml:"version"`
	Data    DataConfig    `yaml:"data"`
	Flags   FlagsConfig   `yaml:"flags"`
	Targets TargetsConfig `yaml:"targets"`
}

// DataConfig specifies where source data is located.
type DataConfig struct {
	Source string `yaml:"source"`
}

// FlagsConfig specifies flag handling.
type FlagsConfig struct {
	Source string `yaml:"source"`
	Inline bool   `yaml:"inline"`
}

// TargetsConfig contains all target configurations.
type TargetsConfig struct {
	Dart       *TargetConfig `yaml:"dart,omitempty"`
	TypeScript *TargetConfig `yaml:"typescript,omitempty"`
	Go         *TargetConfig `yaml:"go,omitempty"`
}

// TargetConfig is the configuration for a single target.
type TargetConfig struct {
	Enabled     bool   `yaml:"enabled"`
	Output      string `yaml:"output"`
	PackageName string `yaml:"packageName,omitempty"`
	ModuleName  string `yaml:"moduleName,omitempty"`
}

// Load reads and parses the configuration file.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return &cfg, nil
}

// Default returns a default configuration.
func Default() *Config {
	return &Config{
		Version: "1.0",
		Data: DataConfig{
			Source: "../data",
		},
		Flags: FlagsConfig{
			Source: "../data/assets/flags",
			Inline: true,
		},
		Targets: TargetsConfig{
			Dart: &TargetConfig{
				Enabled:     true,
				Output:      "./packages/dart/lib/src",
				PackageName: "infobits_intl",
			},
			TypeScript: &TargetConfig{
				Enabled:     true,
				Output:      "./packages/typescript/src",
				PackageName: "@infobits/intl",
			},
			Go: &TargetConfig{
				Enabled:    true,
				Output:     "./packages/go",
				ModuleName: "github.com/infobits-io/infobits-intl-go",
			},
		},
	}
}
