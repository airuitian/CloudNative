package unit

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	input := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(input); i++ {
		ret := square(input[i])
		if ret != expected[i] {
			t.Errorf("input is %d, expected is %d, actual %d", input[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("error")
	fmt.Println("end")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("error")
	fmt.Println("end")
}
