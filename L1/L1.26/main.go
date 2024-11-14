// 26. Разработать программу, которая проверяет,
//что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func unic(s string) bool {
	strings.ToLower(s)
	mapa := make(map[rune]struct{})
	for _, elem := range s {
		if elem != ' ' {
			if _, ok := mapa[elem]; ok {
				return false
			} else {
				mapa[elem] = struct{}{}
			}
		}
	}
	return true
}

func main() {
	//a := "adafCVWSc"
	//b := "ascvfgL "
	c := "ascv  fgL  "
	fmt.Println(unic(c))

}
