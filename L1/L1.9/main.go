// // Разработать конвейер чисел.
// // Даны два канала:
// // в первый пишутся числа (x) из массива,
// // во второй — результат операции x*2,
// // после чего данные из второго канала должны выводиться в stdout.

// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	//wg := sync.WaitGroup{}
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("Горутина завершена по таймауту")
// 				close(ch1)
// 				return
// 			default:
// 				ch1 <- rand.Int()
// 				time.Sleep(time.Second * 2)
// 			}
// 		}
// 	}()

// 	go func(ctx context.Context) {
// 		for val := range ch1 {
// 			ch2 <- val * val
// 		}
// 		close(ch2)
// 	}(ctx)

// 	for val := range ch2 {
// 		fmt.Println(val)
// 	}

// }

package main

/*

1) Поочередно выполните http запросы по ссылкам, в случае ответа на запрос
"200 ОК" печатаем - "адрес - ок", в случае другого кода или ошибки -
"адрес - не ок"

2) Модифицируйте программу таким образом, чтобы использовался канал для коммуникации основного
потока с горутинами. Пример: запросы выполняются в горутинах, печать результатов в основном
потоке

*/

import (
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Err        error
	StatusCode int
	Url        string
}

func main() {
	urls := []string{
		"http://vk.com",
		"http://google.com",
		"http://mail.ru",
	}

	channel := make(chan Response)

	wg := sync.WaitGroup{}

	for _, val := range urls {
		wg.Add(1)
		val := val
		go func() {
			defer wg.Done()
			resp, err := http.Get(val)
			if err != nil {
				channel <- Response{
					Err: err,
					Url: val,
				}
			} else {
				channel <- Response{
					StatusCode: resp.StatusCode,
					Url:        val,
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(channel)
	}()

	for val := range channel {
		if val.Err != nil || val.StatusCode != http.StatusOK {
			fmt.Printf("%s - не ок\n", val.Url)
		} else {
			fmt.Printf("%s - ок\n", val.Url)
		}
	}

}
