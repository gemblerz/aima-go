package aima

import (
	"fmt"
	"strings"
)

type Expr struct {
	Op   string
	Args []*Expr
}

func NewExpr(x string) *Expr {
	fmt.Println(x)
	a := expr_handle_infix_ops(x)
	fmt.Println(a)
	return &Expr{}
}

func (e *Expr) Add(rhs *Expr) *Expr {
	return &Expr{
		Op:   "+",
		Args: []*Expr{e, rhs},
	}
}

func (e *Expr) Print() string {
	if len(e.Args) == 0 {
		return e.Op
	}
	op := " " + e.Op + " "
	var result []string
	for _, e := range e.Args {
		result = append(result, e.Print())
	}
	return strings.Join(result, op)
}

func expr_handle_infix_ops(x string) string {
	infix_ops := []string{"==>", "<==", "<=>"}
	for _, op := range infix_ops {
		x = strings.Replace(x, op, fmt.Sprintf("|%s|", op), -1)
	}
	return x
}

// func eval(x string) []*Expr {

// }
