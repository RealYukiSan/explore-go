package fundamental

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	val int
}

func (c *counter) Value() (x int) {
	return c.val
}

func (c *counter) Add(x int) {
	c.val++
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

	wg.Done()
	fmt.Println(meter.Value())
}
