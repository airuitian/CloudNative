package variable_test

import "testing"

func TestFibList(t *testing.T) {
	a := 1
	b := 0

	for i := 0; i < 10; i++ {
		t.Log(b)
		a, b = b, a+b
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
