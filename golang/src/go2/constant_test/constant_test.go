package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstant2(t *testing.T) {
	a := 4
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
