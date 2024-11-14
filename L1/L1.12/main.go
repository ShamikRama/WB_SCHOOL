// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	mapa := make(map[string]struct{})

	// добавляем ключи в мапу, т.к все ключи уникальные,
	//они не повторяются
	for _, val := range arr {
		mapa[val] = struct{}{}
	}

	for key := range mapa {
		fmt.Println(key)
	}
}
