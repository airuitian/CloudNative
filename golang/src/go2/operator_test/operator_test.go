package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	b := [4]int{1, 3, 2, 4}
	c := [4]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	a = a &^ Writable

	t.Log(a&Readable == Readable, a&Writable == Writable)
}
