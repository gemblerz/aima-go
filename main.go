package main

import (
	"fmt"

	aima "github.com/gemblerz/aima-go/pkg/aima"
)

func main() {
	// exp := aima.NewExpr("Main(new) ==> Run(now)")
	// fmt.Print(exp)
	a := aima.Expr{Op: "x"}
	c := a.Add(&aima.Expr{Op: "y"})
	fmt.Print(c.Print())
}
