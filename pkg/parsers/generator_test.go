package parsers

import (
	"os"
	"testing"
	"text/template"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultGenerator(t *testing.T) {
	Convey("It should implements Generator interface", t, func() {
		So(new(FromFileGenerator), ShouldImplement, (*Generator)(nil))
	})
}

func TestDefaultGenerator_WriteTaskToString_FromFile(t *testing.T) {
	judger := new(Codeforces)
	task := &DefaultTask{
		title: "Hello World!",
	}
	tmplDir := "/Users/psucoder/projects/go/src/github.com/psucodervn/iir/templates"
	if _, err := os.Stat(tmplDir); err != nil {
		t.Skipf("templates folder [%s] not found", tmplDir)
	}
	g := NewFromFileGenerator(judger, tmplDir)
	Convey("It should return correct generated code", t, func() {
		So(task, ShouldNotBeNil)
		str, err := g.WriteTaskToString(task)
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})
}

func TestDefaultGenerator_WriteTaskToString_FromMemory(t *testing.T) {
	judger := new(Codeforces)
	task := &DefaultTask{
		title: "Hello World!",
	}
	tmpl, err := template.New("main").Parse(strMainCCTmpl)
	if err != nil {
		t.Skip("parse template failed")
	}

	g := NewFromMemoryGenerator(judger, tmpl)
	Convey("It should return correct generated code", t, func() {
		So(task, ShouldNotBeNil)
		str, err := g.WriteTaskToString(task)
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})
}

var strMainCCTmpl = `{{ define "main" }}
/*
  Task: {{ .Title }}
*/
#include <bits/stdc++.h>
using namespace std;

int main(argv int, argc* []char) {
  return 0;
}

{{ end }}
`
