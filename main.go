package main

import (
	"fmt"

	"github.com/BonyChops/mulprec-pi-go/pkg/mulprec"
)

func main() {

	fmt.Println("Hello, playground")
	var a, b mulprec.NUMBER
	a.N = []mulprec.NUMBER_TYPE{1, 2, 3}
	b.N = []mulprec.NUMBER_TYPE{4, 5, 7}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	fmt.Println(a, "+", b)
	a.Add(&b)
	fmt.Println(a)
	fmt.Println()

	fmt.Println(a, "-", b)
	a.Sub(&b)
	fmt.Println(a)
	fmt.Println()

	fmt.Println(a, "*", b)
	a.Mul(&b)
	fmt.Println(a)
	fmt.Println()

}
