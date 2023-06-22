package mulprec

import (
	"fmt"
	"math"
	"strings"
)

type NUMBER_TYPE int64

// const NUMBER_TYPE_MAX = NUMBER_TYPE(^uint64(0) >> 1)
const NUMBER_TYPE_MAX = 10
const CAPACITY_SIZE = 1000

type NUMBER struct {
	N    []NUMBER_TYPE
	Sign bool
	Dp   int
}

type NUMBER_DETAILS struct {
	N                   []NUMBER_TYPE
	Sign                bool
	Dp                  int
	RealPartLength      int
	ImaginaryPartLength int
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

func (a *NUMBER) Set(bp int, n NUMBER_TYPE) {
	// TODO: Assume enought capacity slice
	p := a.GetPos(bp)

	if p < 0 {
		newN := make([]NUMBER_TYPE, len(a.N)-p)
		copy(newN[-p:], a.N)
		newN[0] = n
		a.N = newN
		a.Dp = bp + 1

	} else if p >= len(a.N) {
		newN := make([]NUMBER_TYPE, p)
		copy(newN, a.N)
		newN[len(newN)-1] = n
		a.N = newN
	} else {
		a.N[p] = n
	}
}

// Both GetPos and GetRealPos are avaliable with this function
func (a *NUMBER) GetPos(p int) int {
	return a.Dp - p - 1
}

func (a *NUMBER) GetDigit(p int) NUMBER_TYPE {
	realPos := a.GetPos(p)
	if realPos < 0 || realPos >= len(a.N) {
		return 0
	}
	return a.N[realPos]
}

func (a *NUMBER) FixDigits() {
	for bi := range a.N {
		i := len(a.N) - bi - 1
		v := a.N[i]
		if v >= NUMBER_TYPE_MAX {
			pos := a.GetPos(i)
			a.Set(pos+1, a.GetDigit(pos+1)+NUMBER_TYPE(math.Floor(float64(v)/NUMBER_TYPE_MAX)))
			a.Set(pos, v%NUMBER_TYPE_MAX)
		}
		if v < 0 {
			pos := a.GetPos(i)
			if i == 0 {
				fmt.Println()
				fmt.Println(a)
				panic("This NUMBER is corrupted.")
			}
			a.Set(pos+1, a.GetDigit(pos+1)+NUMBER_TYPE(math.Floor(float64(v)/NUMBER_TYPE_MAX)))
			a.Set(pos, NUMBER_TYPE_MAX-(v%NUMBER_TYPE_MAX))
		}
	}

}

func (a *NUMBER) Analyze() NUMBER_DETAILS {
	return NUMBER_DETAILS{
		N:                   a.N,
		Sign:                a.Sign,
		Dp:                  a.Dp,
		RealPartLength:      a.Dp,
		ImaginaryPartLength: len(a.N) - a.Dp,
	}
}

func (a *NUMBER) Add(b *NUMBER) {
	if a.Sign != b.Sign {
		a.Sub(b)
		return
	}
	for i, v := range b.N {
		pos := b.GetPos(i)
		a.Set(pos, a.GetDigit(pos)+v)
	}
	a.FixDigits()
}

func (a *NUMBER) compare(b *NUMBER) int {
	if a.Sign != b.Sign {
		if a.Sign {
			return 1
		} else {
			return -1
		}
	}

	aPos, bPos := a.GetPos(0), b.GetPos(0)
	if aPos != bPos {
		if aPos > bPos {
			return 1
		} else {
			return -1
		}
	}

	for i, v := range a.N {
		if v != b.N[i] {
			if v > b.N[i] {
				return 1
			} else {
				return -1
			}
		}
	}

	return 0
}

func (a *NUMBER) Sub(b *NUMBER) {
	if a.Sign != b.Sign {
		a.Add(b)
		return
	}

	var dst, src *NUMBER
	aIsSmaller := false
	if a.compare(b) == 1 {
		dst, src = a, b
	} else {
		dst, src = b, a
		aIsSmaller = true
	}

	var min, max int
	aDetails, bDetails := a.Analyze(), b.Analyze()
	if aDetails.RealPartLength > bDetails.RealPartLength {
		max = aDetails.RealPartLength - 1
	} else {
		max = bDetails.RealPartLength - 1
	}

	if aDetails.ImaginaryPartLength > bDetails.ImaginaryPartLength {
		min = -aDetails.ImaginaryPartLength
	} else {
		min = -bDetails.ImaginaryPartLength
	}

	for i := min; i <= max; i++ {
		a.Set(i, dst.GetDigit(i)-src.GetDigit(i))
	}

	if aIsSmaller {
		a.Sign = !a.Sign
	}

	a.FixDigits()
}

func (a *NUMBER) Mul(b *NUMBER) {
	c := NUMBER{
		N:    make([]NUMBER_TYPE, len(a.N)+len(b.N)),
		Sign: a.Sign != b.Sign,
		Dp:   0, //a.Dp + b.Dp,
	}
	for ai, av := range a.N {
		for bi, bv := range b.N {
			c.Set(a.GetPos(ai)+b.GetPos(bi), c.GetDigit(a.GetPos(ai)+b.GetPos(bi))+av*bv)
		}
	}

	c.FixDigits()
	*a = c
}
