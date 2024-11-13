package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Запрашиваем у пользователя количество горутин
	fmt.Print("Введите количество горутин: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Преобразуем введенное значение в целое число
	numGoroutines, _ := strconv.Atoi(input)
	wg := sync.WaitGroup{}
	channel := make(chan int)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := 0
			for {
				select {
				case channel <- val:
					val++
					fmt.Printf("Записали в канал %d\n", val)
					time.Sleep(time.Second * 2)
				default:
					//fmt.Println("Канал закрыт")
					return
				}
			}
		}()
	}

	for val := range channel {
		fmt.Printf("Считали из канала %d\n", val)
	}

	wg.Wait()
}
