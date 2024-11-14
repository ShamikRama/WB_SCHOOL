// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	mapa := make(map[string]struct{})

	for _, val := range arr {
		mapa[val] = struct{}{}
	}

	for key := range mapa {
		fmt.Println(key)
	}
}
