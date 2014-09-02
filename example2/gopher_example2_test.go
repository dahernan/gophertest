package example2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Using goconvey to write test BDD style
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
