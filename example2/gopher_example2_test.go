package example2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGopherCanRun(t *testing.T) {
	Convey("Given a gopher called Adam", t, func() {
		g := Gopher{"Adam", 12}
		Convey("When Adam Run", func() {
			result, err := g.Run("fun!")
			Convey("then Adam has fun", func() {
				So(result, ShouldEqual, "I am Adam and I run for fun!")
				So(err, ShouldBeNil)
			})
		})
	})
}
