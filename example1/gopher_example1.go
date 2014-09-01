package example1

import (
	"fmt"
)

type Runner interface {
	Run(string) (string, error)
}

type Gopher struct {
	name string
	age  int
}

func (g Gopher) Run(desc string) (string, error) {
	return fmt.Sprintf("I am %s and I run for %s", g.name, desc), nil
}

// func (g *Gopher) Name() string {
// 	return g.name
// }
