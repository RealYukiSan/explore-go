package fundamental

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Value() (x int) {
	return c.val
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func ImplementMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
