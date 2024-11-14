// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

// func main() {
// 	arr := []int{1, 2}
// 	arr[0], arr[1] = arr[1], arr[0]
// 	fmt.Println(arr)
// }

func main() {
	a := 10
	b := 20

	// Меняем местами значения переменных a и b
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)

}
