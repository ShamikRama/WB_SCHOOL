// Написать программу, конкурентно рассчитывает квадраты чисел из массива
// (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main 

import (
	"fmt"
	"sync"
)

func main(){
	array := [5]int{2,4,6,8,10}
	square := make(chan int, len(array))
	wg := sync.WaitGroup{}

	for _, val := range array {
		wg.Add(1)
		go func (val int) {
			defer wg.Done()
			square <- val * val
		} (val)
	}

	

	for i := 0; i < len(array); i ++ {
			fmt.Println(<-square)
	}

}









// package main

// import "fmt"

// func squares(c chan int, numbers []int) {
// 	for _, num := range numbers {
// 		c <- num * num
// 	}
// 	close(c)
// }

// func main() {
// 	arrayInt := []int{2, 4, 6, 8, 10}

// 	c := make(chan int)

// 	go squares(c, arrayInt)

// 	for val := range c {
// 		fmt.Println(val)
// 	}
// }
