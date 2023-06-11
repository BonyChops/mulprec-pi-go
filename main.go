package main

import (
	"fmt"

	"github.com/BonyChops/mulprec-pi-go/pkg/mulprec"
)

func main() {
	fmt.Println("Hello, playground")
	var a mulprec.NUMBER
	a.N = []mulprec.NUMBER_TYPE{1, 2, 3}
	fmt.Println(a)
	a.Set(-4, 9)
	fmt.Println(a)
	a.Set(0, 9)
	fmt.Println(a)
	a.Set(5, 9)
	fmt.Println(a)
}
