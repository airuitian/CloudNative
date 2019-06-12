package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr1[1])
	t.Log(arr, arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySelection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec := arr[1:4]
	t.Log(arr_sec)
}
