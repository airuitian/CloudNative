package benmark

import "testing"

type Content struct {
	Detail [10000]int
}

func withValue(arr [1000]Content) int {
	return 0
}

func withRef(arr *[1000]Content) int {
	return 0
}

func BenchmarkPassValue(b *testing.B) {
	var arr [1000]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassRef(b *testing.B) {
	var arr [1000]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withRef(&arr)
	}
	b.StopTimer()
}
