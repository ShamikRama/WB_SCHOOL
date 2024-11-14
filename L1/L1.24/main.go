// 24. Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(a, b float64) Point {
	return Point{
		x: a,
		y: b,
	}
}

func Distance(g, h Point) float64 {
	num1 := math.Pow(g.x-h.x, 2)
	num2 := math.Pow(g.y-h.y, 2)
	res := math.Sqrt(num1 + num2)
	return res
}

func main() {
	a := New(45.54, 78)
	b := New(67.36, 89.56)
	fmt.Println(Distance(a, b))

}
