package parsers

import (
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultGenerator(t *testing.T) {
	Convey("It should implements Generator interface", t, func() {
		So(new(DefaultGenerator), ShouldImplement, (*Generator)(nil))
	})
}

func TestDefaultGenerator_WriteTaskToString(t *testing.T) {
	judger := new(Codeforces)
	task := &DefaultTask{
		title: "Hello World!",
	}
	wd := "/Users/psucoder/projects/go/src/github.com/psucodervn/iir"
	g := NewDefaultGenerator(judger, path.Join(wd, "templates"))
	Convey("It should return correct generated code", t, func() {
		So(task, ShouldNotBeNil)
		str, err := g.WriteTaskToString(task)
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})
}
