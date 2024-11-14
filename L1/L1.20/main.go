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

// надо сделать еще :
// 21. Реализовать паттерн «адаптер» на любом примере.
// 23. Удалить i-ый элемент из слайса.
// 24. Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
// 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
