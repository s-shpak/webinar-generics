package practice

import (
	"reflect"
	"sync"
	"testing"
)

func TestForEach(t *testing.T) {
	type testCase[S ~[]E, E comparable] struct {
		s        S
		fn       func(e E) E
		expected S
	}
	intTestCase := testCase[[]int, int]{
		s: []int{1, 2, 3, 4},
		fn: func(e int) int {
			return e * 2
		},
		expected: []int{2, 4, 6, 8},
	}

	actualInt := ForEach(intTestCase.s, intTestCase.fn)
	if !reflect.DeepEqual(intTestCase.expected, actualInt) {
		t.Fatalf("expected %v, got %v", intTestCase.expected, actualInt)
	}
	strTestCase := testCase[[]string, string]{
		s: []string{"hello", "world"},
		fn: func(e string) string {
			return e + e
		},
		expected: []string{"hellohello", "worldworld"},
	}
	actualStr := ForEach(strTestCase.s, strTestCase.fn)
	if !reflect.DeepEqual(strTestCase.expected, actualStr) {
		t.Fatalf("expected %v, got %v", strTestCase.expected, actualStr)
	}
}

func TestFanOut(t *testing.T) {
	els := []int{1, 2, 3}

	expectedOut1 := make([]int, len(els))
	copy(expectedOut1, els)
	expectedOut2 := make([]int, len(els))
	copy(expectedOut2, els)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	inp := make(chan int)
	fo := NewFanOut(inp)

	out1 := fo.Take()
	out2 := fo.Take()

	var actualOut1 []int
	var actualOut2 []int

	go func() {
		defer wg.Done()
		for el := range out1 {
			actualOut1 = append(actualOut1, el)
		}
	}()

	go func() {
		defer wg.Done()
		for el := range out2 {
			actualOut2 = append(actualOut2, el)
		}
	}()

	for _, el := range els {
		inp <- el
	}
	close(inp)

	wg.Wait()

	if !reflect.DeepEqual(expectedOut1, actualOut1) {
		t.Fatalf("expected first out %v, got %v", expectedOut1, actualOut1)
	}

	if !reflect.DeepEqual(expectedOut2, actualOut2) {
		t.Fatalf("expected first out %v, got %v", expectedOut2, actualOut2)
	}
}
