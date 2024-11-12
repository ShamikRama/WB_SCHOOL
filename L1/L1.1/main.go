// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action 
// от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name string
	Age int
}

func (h *Human) Say(){
	fmt.Printf("Всем привет меня зовут %s и мне %d лет", h.Name, h.Age)
}

type Action struct {
	Human
}

func main(){
	me := Action{
		Human : Human{
			Name : "Иван",
			Age : 45,
		},
	}
	me.Say()
}