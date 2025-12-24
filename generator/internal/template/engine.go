package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Engine wraps Go's template engine with custom functions.
type Engine struct {
	templateDir string
	templates   map[string]*template.Template
}

// NewEngine creates a new template engine.
func NewEngine(templateDir string) *Engine {
	return &Engine{
		templateDir: templateDir,
		templates:   make(map[string]*template.Template),
	}
}

// LoadTemplates loads all templates from a subdirectory.
func (e *Engine) LoadTemplates(subdir string) error {
	dir := filepath.Join(e.templateDir, subdir)

	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("reading template directory %s: %w", dir, err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".tmpl" {
			continue
		}

		name := file.Name()
		path := filepath.Join(dir, name)

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading template %s: %w", name, err)
		}

		tmpl, err := template.New(name).Funcs(FuncMap()).Parse(string(content))
		if err != nil {
			return fmt.Errorf("parsing template %s: %w", name, err)
		}

		key := filepath.Join(subdir, name)
		e.templates[key] = tmpl
	}

	return nil
}

// Execute executes a template with the given data.
func (e *Engine) Execute(templatePath string, data any) (string, error) {
	tmpl, ok := e.templates[templatePath]
	if !ok {
		return "", fmt.Errorf("template not found: %s", templatePath)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing template %s: %w", templatePath, err)
	}

	return buf.String(), nil
}

// ExecuteToFile executes a template and writes the result to a file.
func (e *Engine) ExecuteToFile(templatePath, outputPath string, data any) error {
	content, err := e.Execute(templatePath, data)
	if err != nil {
		return err
	}

	// Create output directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", dir, err)
	}

	if err := os.WriteFile(outputPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("writing file %s: %w", outputPath, err)
	}

	return nil
}

// RenderString renders a template string with the given data.
func RenderString(templateStr string, data any) (string, error) {
	tmpl, err := template.New("inline").Funcs(FuncMap()).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return buf.String(), nil
}
