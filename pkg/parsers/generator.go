package parsers

import (
	"errors"
	"io"
	"path"
	"strings"
	"text/template"
)

var (
	// ErrInvalidTemplate is invalid template error
	ErrInvalidTemplate = errors.New("invalid template")
)

// Generator is the code generator interface
type Generator interface {
	WriteTask(writer io.Writer, task Task, tmplName string) error
	WriteTaskToString(task Task, tmplName string) (string, error)
}

// FromMemoryGenerator generates code from parsed template
type FromMemoryGenerator struct {
	tmpl *template.Template
}

// NewFromMemoryGenerator returns new instance of Generator
func NewFromMemoryGenerator(tmpl *template.Template) Generator {
	return &FromMemoryGenerator{
		tmpl: tmpl,
	}
}

// WriteTask writes the generated code to writer
func (g *FromMemoryGenerator) WriteTask(writer io.Writer, task Task, tmplName string) error {
	if g.tmpl == nil {
		return ErrInvalidTemplate
	}
	return g.tmpl.ExecuteTemplate(writer, tmplName, task)
}

// WriteTaskToString returns generated code in string format
func (g *FromMemoryGenerator) WriteTaskToString(task Task, tmplName string) (string, error) {
	sb := &strings.Builder{}
	if err := g.WriteTask(sb, task, tmplName); err != nil {
		return "", err
	}
	return sb.String(), nil
}

// FromFileGenerator generates code from file templates
type FromFileGenerator struct {
	judger Judger
	tplDir string
}

// WriteTask writes the generated code to writer
func (g *FromFileGenerator) WriteTask(writer io.Writer, task Task, tmplName string) error {
	tmpl, err := template.ParseFiles(path.Join(g.tplDir, "main.cc.tmpl"))
	if err != nil {
		return err
	}
	var data = struct {
		Task
		OutDir string
	}{
		task,
		outDir,
	}
	return tmpl.ExecuteTemplate(writer, tmplName, data)
}

// WriteTaskToString returns generated code in string format
func (g *FromFileGenerator) WriteTaskToString(task Task, tmplName string) (string, error) {
	sb := &strings.Builder{}
	if err := g.WriteTask(sb, task, tmplName); err != nil {
		return "", err
	}
	return sb.String(), nil
}

// NewFromFileGenerator returns new instance of FromFileGenerator
func NewFromFileGenerator(judger Judger, tplDir string) Generator {
	return &FromFileGenerator{
		judger: judger,
		tplDir: tplDir,
	}
}
