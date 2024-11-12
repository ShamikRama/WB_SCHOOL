// // Дана последовательность чисел: 2,4,6,8,10.
// // Найти сумму их квадратов

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main(){
// 	array := []int{2,4,6,8,10}
// 	sum := 0
// 	channel := make(chan int)
// 	wg := sync.WaitGroup{}
// 	for _, val := range array{
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			channel <- val * val
// 		}()
// 	}
// 	go func(){
// 		wg.Wait()
// 		close(channel)
// 	}()
	
// 	for val := range channel {
// 		sum +=val
// 	}
// 	fmt.Println(sum)
// }




package main

import (
	"fmt"
	"sync"
)

func main(){
	array := []int{2,4,6,8,10}
	sum := 0
	channel := make(chan int)
	wg := sync.WaitGroup{}

	go func(){
		for val := range channel {
			sum += val
		} 
	}()


	for _, val := range array{
		wg.Add(1)
		go func(){
			defer wg.Done()
			channel <- val * val
		}()
	}

	wg.Wait()
	close(channel)
	
	fmt.Println(sum)
}

