package example3

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Test isolation with Convey context
func TestConveyContextIsWorking(t *testing.T) {
	Convey("Given a slice of two gophers", t, func() {

		adam := Gopher{"Adam", 12}
		jaime := Gopher{"Jaime", 14}

		gophers := []Gopher{adam, jaime}

		Convey("append a new gopher in this context", func() {
			tom := Gopher{"Tom", 1}
			gophers = append(gophers, tom)
			Convey("then the slice should contain the gopher Tom", func() {
				So(gophers[2].name, ShouldEqual, "Tom")
			})
		})

		So(len(gophers), ShouldEqual, 2)
	})
}
