package benmark

import "testing"

const numOfElmes = 100000

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < numOfElmes; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElmes)
		for j := 0; j < numOfElmes; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOversizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElmes*8)
		for j := 0; j < numOfElmes; j++ {
			s = append(s, j)
		}
	}
}
