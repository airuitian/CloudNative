package csp_test

import (
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 2000):
		t.Error("timeout")
	}
}
