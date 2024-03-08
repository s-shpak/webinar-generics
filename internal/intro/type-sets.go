package intro

import "strconv"

type Number interface {
	int | int16 | int32 | int64
}

func PrintNumber[T Number](n T) string {
	return strconv.FormatInt(int64(n), 10)
}

type NewNumber interface {
	~int | ~int16 | ~int32 | ~int64
}

func PrintNewNumber[T NewNumber](n T) string {
	return strconv.FormatInt(int64(n), 10)
}

type MyInt int

func PrintMyInt(i MyInt) string {
	return PrintNewNumber(i)
}
