// Разработать программу которая переворачивает строку
// sun dog cat - cat dog sun

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "sun cat"
	words := strings.Fields(s) // Разбиваем строку на слова

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку
	res := strings.Join(words, " ")
	fmt.Println(res) // Вывод: cat dog sun
}
