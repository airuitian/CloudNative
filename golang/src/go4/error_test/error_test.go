package error_test

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"testing"
)

const delta = 1e-6

var ErrNegativeSqrt = errors.New("Sqrt: negative number")

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt
	}
	z := x / 2
	fmt.Println(z)
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z, nil
}

func TestSqrt(t *testing.T) {

	if v, err := Sqrt(-1); err != nil {
		if err == ErrNegativeSqrt {
			fmt.Println(".....")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}

}

func TestSqrt2(t *testing.T) {
	GetSqrt("2")
}

func GetSqrt(str string) {
	var (
		i   int
		err error
		r   float64
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return

	}
	if r, err = Sqrt(float64(i)); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(r)

}
