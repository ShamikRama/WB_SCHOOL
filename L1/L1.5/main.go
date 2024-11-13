package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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

	for val := range channel {
		fmt.Println(val)
	}

}

// Разработать программу,
// которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
