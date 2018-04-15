package parsers

import (
	"io"
	"path"
	"strings"
	"text/template"
)

// Generator is the code generator interface
type Generator interface {
	WriteTask(writer io.Writer, task Task) error
	WriteTaskToString(task Task) (string, error)
}

// DefaultGenerator is the default implement of Generator interface
type DefaultGenerator struct {
	judger Judger
	tplDir string
}

// WriteTask writes the generated code to writer
func (g *DefaultGenerator) WriteTask(writer io.Writer, task Task) error {
	tmpl, err := template.ParseFiles(path.Join(g.tplDir, "main.cc.tmpl"))
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(writer, "main", task)
	if err != nil {
		return err
	}
	return nil
}

// WriteTaskToString returns generated code in string format
func (g *DefaultGenerator) WriteTaskToString(task Task) (string, error) {
	sb := &strings.Builder{}
	if err := g.WriteTask(sb, task); err != nil {
		return "", err
	}
	return sb.String(), nil
}

// NewDefaultGenerator returns new instance of DefaultGenerator
func NewDefaultGenerator(judger Judger, tplDir string) Generator {
	return &DefaultGenerator{
		judger: judger,
		tplDir: tplDir,
	}
}
