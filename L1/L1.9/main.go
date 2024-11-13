// Разработать конвейер чисел.
// Даны два канала:
// в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//wg := sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена по таймауту")
				close(ch1)
				return
			default:
				ch1 <- rand.Int()
				time.Sleep(time.Second * 2)
			}
		}
	}()

	go func(ctx context.Context) {
		for val := range ch1 {
			ch2 <- val * val
		}
		close(ch2)
	}(ctx)

	for val := range ch2 {
		fmt.Println(val)
	}

}
