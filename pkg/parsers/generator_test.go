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
	task := Task{
		Name: "Hello World!",
	}
	tmplDir := "/Users/psucoder/projects/go/src/github.com/psucodervn/iir/templates"
	if _, err := os.Stat(tmplDir); err != nil {
		t.Skipf("templates folder [%s] not found", tmplDir)
	}
	g := NewFromFileGenerator(judger, tmplDir)
	Convey("It should return correct generated code", t, func() {
		So(task, ShouldNotBeNil)
		str, err := g.WriteTaskToString(task, "main.cc")
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})
}

func TestDefaultGenerator_WriteTaskToString_FromMemory(t *testing.T) {
	task := Task{
		Name: "Hello World!",
	}
	tmpl, err := template.New("main").Parse(strMainCCTmpl)
	if err != nil {
		t.Skip("parse template failed")
	}

	g := NewFromMemoryGenerator(tmpl)
	Convey("It should return correct generated code", t, func() {
		So(task, ShouldNotBeNil)
		str, err := g.WriteTaskToString(task, "main.cc")
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})
}

var strMainCCTmpl = `{{ define "main.cc" }}
/*
  Task: {{ .Name }}
*/
#include <bits/stdc++.h>
using namespace std;

int main(argv int, argc* []char) {
  ios::sync_with_stdio(false); cin.tie(nullptr);
  return 0;
}

{{ end }}
`
