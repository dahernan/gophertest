package example4

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mocking is easy when you use interfaces
type RasingErrorRunner struct{}

func (r RasingErrorRunner) Run(string) (string, error) {
	return "", fmt.Errorf("Error here")
}

func TestRunAllReturnsError(t *testing.T) {
	Convey("Run all returns an error if some of the runners, return a error", t, func() {
		g := Gopher{"Adam", 12}
		err := RunAll(g, RasingErrorRunner{})
		So(err, ShouldNotBeNil)
	})
}

// but is not easy  if you want to validate the arguments
type SomeRunner struct {
	mock.Mock
}

func (s *SomeRunner) Run(desc string) (string, error) {
	args := s.Mock.Called(desc)
	return args.String(0), args.Error(1)
}

func TestRunAllUsesTheRightArguments(t *testing.T) {
	Convey("Run all returns an error if some of the runners, return a error", t, func() {

		runner := &SomeRunner{}
		arg1 := "run all the things!"
		runner.On("Run", arg1).Return("OK", nil)

		err := RunAll(runner)

		runner.Mock.AssertExpectations(t)
		So(err, ShouldBeNil)
	})
}
