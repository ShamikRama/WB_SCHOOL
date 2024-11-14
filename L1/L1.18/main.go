// Реализовать структуру счетчик которая бцдет инкремитироваться
// в конкурентной среде . По заверешению программа должна
// выводить итоговое значение счетяика

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancel()
	wg := sync.WaitGroup{}
	c := NewDefault()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.mu.Lock()
			c.val++
			c.mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(c.val)
}

func NewDefault() *Counter {
	return &Counter{
		val: 0,
	}
}
