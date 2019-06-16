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
	//retCh := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func TestAsyncService(t *testing.T) {
	ret := AsyncService()
	otherTask()
	fmt.Println(<-ret)
}
