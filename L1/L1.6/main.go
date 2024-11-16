// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"math/rand"
	"fmt"
	"sync"
	"time"
)

// // с помощью дополнительного канала
func main() {
	channel := make(chan int, 5)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			j, ok := <-channel
			if ok {
				fmt.Printf("Значение равно %d\n", j)
			} else {
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		channel <- i
	}

	close(channel)
	<-done
}

// // с помощю контекста с таймаутом
func main() {
	channel := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена по таймауту")
				close(channel)
				return
			default:
				channel <- rand.Int()
				time.Sleep(time.Second * 2)
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range channel {
			fmt.Println(val)
		}
	}()

	wg.Wait()

}

// с помощью wg.Wait()
func main() {
	channel := make(chan int, 5)
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(<-channel)
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			channel <- i
		}()
	}

	wg.Wait()
}
