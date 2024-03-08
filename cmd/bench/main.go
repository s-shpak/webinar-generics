package main

func main() {
	ctr := &counterImpl{}

	addsubmulGeneric(ctr)
	addsubmulGeneric(*ctr)

	ctr2 := &counterImpl2{}

	addsubmulGeneric(ctr2)
	addsubmulGeneric(*ctr2)
}

type counter interface {
	Add()
	Sub()
	Multiply()
}

type counterImpl struct {
	n int
}

func (p counterImpl) Add() {
	p.n++
}

func (p counterImpl) Sub() {
	p.n--
}

func (p counterImpl) Multiply() {
	_ = p.n * 2
}

type counterImpl2 struct {
	n int
}

func (p counterImpl2) Add() {
	p.n++
}

func (p counterImpl2) Sub() {
	p.n--
}

func (p counterImpl2) Multiply() {
	_ = p.n * 2
}

func addsubmul(p counter) {
	p.Add()
	p.Multiply()
	p.Sub()
}

func addsubmulGeneric[T counter](p T) {
	p.Add()
	p.Multiply()
	p.Sub()
}
