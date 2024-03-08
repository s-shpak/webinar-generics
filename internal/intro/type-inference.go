package intro

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func UseGMin() {
	x, y := 42, -42

	firstMin := GMin[int](x, y)
	fmt.Printf("first min: %d\n", firstMin)

	var x32, y32 int32 = 42, -42
	secondMin := GMin(x32, y32)
	fmt.Printf("second min: %d\n", secondMin)
}

func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	elems := make([]string, len(p))
	for i, el := range p {
		elems[i] = strconv.FormatInt(int64(el), 10)
	}
	return strings.Join(elems, ", ")
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r)
	// fmt.Println(r.String())
}

func NewScale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func NewScaleAndPrint(p Point) {
	r := NewScale(p, 2)
	fmt.Println(r.String())
}
