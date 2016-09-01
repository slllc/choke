package choke

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// Test001 - run this test with a width of > 1 in the choke
// and it will randomly fail.
func Test001(t *testing.T) {
	fmt.Printf("mp = %d\n", runtime.GOMAXPROCS(-1))

	ch := New(1).Start(context.Background())

	const w = 2048
	var total int

	var wg sync.WaitGroup

	wg.Add(w)

	for i := 0; i < w; i++ {
		go func(n int) {
			ch.Do(func() error {
				total++
				runtime.Gosched()
				total--
				runtime.Gosched()
				return nil
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("total = %d\n", total)
	if total != 0 {
		t.Fail()
	}
}
