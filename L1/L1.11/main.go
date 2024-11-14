// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 5}
	arr2 := []int{6, 66, 245, 4, 5}

	mapa := make(map[int]int)
	res := make([]int, 0)
	for _, val := range arr1 {
		if mapa[val] == 0 { // если элемента в мапе нет
			mapa[val] = 1
		} else {
			mapa[val] += 1
		}
	}
	for _, val := range arr2 {
		if mapa[val] > 0 { // проверяем есть ли этот элемент в мапе
			mapa[val] -= 1
			res = append(res, val)
		}
	}
	fmt.Println(res)
}
