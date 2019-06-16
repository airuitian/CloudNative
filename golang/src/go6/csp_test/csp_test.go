package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 1000)
	return "done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
	}()

	return retCh
}

func TestAsyncService(t *testing.T) {
	ret := AsyncService()
	otherTask()
	fmt.Println(<-ret)
}
