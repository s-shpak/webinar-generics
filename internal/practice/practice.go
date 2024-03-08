package practice

func ForEach[S ~[]E, E any](s S, fn func(e E) E) S {
	ns := make([]E, len(s))
	for i, e := range s {
		ne := fn(e)
		ns[i] = ne
	}
	return ns
}

type FanOut[T any] struct {
	out   []chan T
	input <-chan T
}

func NewFanOut[T any](input <-chan T) *FanOut[T] {
	ld := &FanOut[T]{
		input: input,
	}
	go ld.run()
	return ld
}

func (l *FanOut[T]) run() {
	for links := range l.input {
		for _, c := range l.out {
			c <- links
		}
	}
	for _, c := range l.out {
		close(c)
	}
}

func (l *FanOut[T]) Take() <-chan T {
	c := make(chan T)
	l.out = append(l.out, c)
	return c
}
