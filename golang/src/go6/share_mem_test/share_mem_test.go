package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 50000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 50000; i++ {
		go func() {
			mut.Lock()
			defer func() {
				mut.Unlock()
			}()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			mut.Lock()
			defer func() {
				mut.Unlock()
			}()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
