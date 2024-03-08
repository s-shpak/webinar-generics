package main

import "testing"

func BenchmarkIface(b *testing.B) {
	b.Run("ifaceByPointer", func(b *testing.B) {
		ctr := &counterImpl{}
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				addsubmul(ctr)
			}
		}
	})
}

func BenchmarkGenericPointer(b *testing.B) {
	b.Run("genericByPointer", func(b *testing.B) {
		ctr := &counterImpl{}
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				addsubmulGeneric(ctr)
			}
		}
	})
}

func BenchmarkGenericValue(b *testing.B) {
	b.Run("genericByValue", func(b *testing.B) {
		ctr := counterImpl{}
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				addsubmulGeneric(ctr)
			}
		}
	})
}
