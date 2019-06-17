package unsafe_test

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	t.Log(float64(i))

	t.Log(*(*float64)(unsafe.Pointer(&i)))
}

type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}
