package anyone_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	ch := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("after:", runtime.NumGoroutine())
}
