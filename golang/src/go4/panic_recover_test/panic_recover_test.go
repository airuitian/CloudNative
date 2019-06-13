package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverd from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something wrong"))
	fmt.Println("end")
}
