package testbench_test

// https://philpearl.github.io/post/bad_go_slice_of_pointers/

import "testing"

type MyStruct struct {
	A int
	B int
}

const startCapacity = 10

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]*MyStruct, 0, startCapacity)
		for j := 0; j < 100; j++ {
			slice = append(slice, &MyStruct{A: j, B: j + 1})
		}
	}
}
func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]MyStruct, 0, startCapacity)
		for j := 0; j < 100; j++ {
			slice = append(slice, MyStruct{A: j, B: j + 1})
		}
	}
}
