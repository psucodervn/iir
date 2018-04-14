package cmd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicServer(t *testing.T) {
	Convey("It should implements Server interface", t, func() {
		server := new(HTMLServer)
		So(server, ShouldImplement, (*Server)(nil))
	})
}
