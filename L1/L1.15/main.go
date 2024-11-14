// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить?
// Приведите корректный пример реализации.

package main

import "fmt"

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

// вот так надо исправить :
package main

import (
	"fmt"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// Создаем новый срез байтов и копируем в него первые 100 байтов из v
	bytes := make([]byte, 100)
	copy(bytes, v[:100])
	justString = string(bytes)
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	// Создаем срез байтов длиной size
	bytes := make([]byte, size)
	// Заполняем срез символом 'A'
	for i := range bytes {
		bytes[i] = 'A'
	}
	// Преобразуем срез байтов в строку
	return string(bytes)
}




// можно еще использовать string.Builder
package main

import (
	"bytes"
	"fmt"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// Используем bytes.Buffer для создания новой строки
	var buffer bytes.Buffer
	buffer.WriteString(v[:100])
	justString = buffer.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	// Создаем срез байтов длиной size
	bytes := make([]byte, size)
	// Заполняем срез символом 'A'
	for i := range bytes {
		bytes[i] = 'A'
	}
	// Преобразуем срез байтов в строку
	return string(bytes)
}