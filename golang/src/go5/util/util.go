package util

import "fmt"

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func Squrae(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	ret := []int{0, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
