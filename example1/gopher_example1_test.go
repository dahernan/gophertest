package example1

import (
	"testing"
)

// Simple test using the standard tools
// We're going to use GoConvey to run it
// https://github.com/smartystreets/goconvey
// 		$ go get github.com/smartystreets/goconvey
// 		$ $GOPATH/bin/goconvey
func TestGopherCanRun(t *testing.T) {

	g := Gopher{"Adam", 12}

	result, err := g.Run("fun!")

	if err != nil {
		t.Errorf("Error %v when a gopher run", err)
	}

	expected := "I am Adam and I run for fun!"
	if expected != result {
		t.Errorf("actual = '%v', expected '%v'", result, expected)
	}

}
