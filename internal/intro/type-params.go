package intro

import (
	"golang.org/x/exp/constraints"
)

type Tree[T comparable] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	if t.value == x {
		return t
	}
	if t.left != nil {
		if f := t.left.Lookup(x); f != nil {
			return f
		}
	}
	if t.right != nil {
		if f := t.right.Lookup(x); f != nil {
			return f
		}
	}
	return nil
}

func GMin[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}
