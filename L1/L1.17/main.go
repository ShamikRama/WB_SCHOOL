// Реализовать бинарный поиск

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 4, 3, 2, 5, 6, 8, 7, 10, 9}
	fmt.Println(binarysearch(arr, 5))
}

func binarysearch(some []int, tar int) int {
	sort.Ints(some)
	left, right := 0, len(some)-1

	for {
		mid := (left + right) / 2
		if some[mid] == tar {
			return mid
		} else if tar > some[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
