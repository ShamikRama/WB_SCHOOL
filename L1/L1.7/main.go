// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	mapa := make(map[int]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			mapa[i] = i * i
		}()
	}

	for key, val := range mapa {

		fmt.Printf("%d - ключб, %d - значение ", key, val)

	}

	wg.Wait()

	fmt.Println(mapa)
}
