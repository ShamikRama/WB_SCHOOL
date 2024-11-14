// Разработать программу,
// которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import "fmt"

func main() {
	str := "adkfww аппыаиыаи"
	fmt.Println(string(reverse(str)))
}

func reverse(some string) []rune {
	runes := []rune(some)
	reversestr := make([]rune, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		reversestr = append(reversestr, runes[i])
	}

	return reversestr
}
