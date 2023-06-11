package mulprec

import (
	"fmt"
	"strings"
)

type NUMBER_TYPE int64

type NUMBER struct {
	N    []NUMBER_TYPE
	Sign bool
	Dp   int
}

func (a NUMBER) String() string {
	var sign string
	if a.Sign {
		sign = "-"
	} else {
		sign = "+"
	}

	n := make([]string, len(a.N))
	for i, num := range a.N {
		n[i] = fmt.Sprint(num)
	}

	return fmt.Sprintf("%v %v . %v", sign, strings.Join(n[:a.Dp], " "), strings.Join(n[a.Dp:], " "))
}

func (a *NUMBER) Set(p int, n NUMBER_TYPE) {
	// TODO: Assume enought capacity slice

	if p >= a.Dp {
		newN := make([]NUMBER_TYPE, len(a.N)+p-a.Dp+1)
		copy(newN[(p-a.Dp+1):], a.N)
		newN[0] = n
		a.N = newN
		a.Dp = p + 1
	} else if a.Dp-p > len(a.N) {
		newN := make([]NUMBER_TYPE, a.Dp-p)
		copy(newN, a.N)
		newN[len(newN)-1] = n
		a.N = newN
	} else {
		a.N[a.Dp-p-1] = n
	}
}

func (a *NUMBER) Add(b *NUMBER) {

}
