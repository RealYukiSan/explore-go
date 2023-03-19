package fundamental

import (
	"fmt"
	"sync"
)

func ImplementWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()
}

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}
