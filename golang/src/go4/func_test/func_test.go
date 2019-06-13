package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(2 * time.Second)
	return op
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsF := timeSpent(slowFun)
	t.Log(tsF(10))

	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

	pos := adder()
	for i := 0; i < 5; i++ {
		t.Log(pos(i))
	}
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
}

func clear() {
	fmt.Println("clear")
}

func TestDefer(t *testing.T) {
	defer clear()

	panic("err")

	fmt.Println("start")

}
