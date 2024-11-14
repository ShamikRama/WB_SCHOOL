package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"sync"
)

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
	go func(){
		defer wg.Done()
		for val := range channel {
		fmt.Println(val)
		}
	}()

	wg.Wait()

}

// Разработать программу,
// которая будет последовательно отправлять значения в канал,
// а с другой сторонекунд программа должна завершаться.
// ы канала — читать.
// По истечению N с