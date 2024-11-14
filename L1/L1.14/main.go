// Разработать программу,
// в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

package main

import "fmt"

func main() {
	var v1 interface{} = 42
	var v2 interface{} = "string"
	var v3 interface{} = true
	var v4 interface{} = make(chan int)

	ggg(v1)
	ggg(v2)
	ggg(v3)
	ggg(v4)
}

func ggg(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Переменная имеет тип int: %d\n", v.(int))
	case string:
		fmt.Printf("Переменная имеет тип string: %s\n", v.(string))
	case bool:
		fmt.Printf("Переменная имеет тип bool: %v\n", v.(bool))
	case chan int:
		fmt.Printf("Переменная имеет тип chan int: %v\n", v.(chan int))
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}
