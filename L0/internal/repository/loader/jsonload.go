package loader

import (
	"L0/internal/repository"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func LoadOrders(dirpath string) (orders []repository.Order, err error) {
	info, err := os.ReadDir(dirpath)
	if err != nil {
		return
	}
	for _, file := range info {
		var filePath string
		if filepath.Ext(file.Name()) == ".json" {
			filePath = filepath.Join(dirpath, file.Name())
		} else {
			fmt.Print("No json files in directory")
		}
		orderbyte, err := ReadJson(filePath)
		if err != nil {
			return
		}
		var order repository.Order
		err = json.Unmarshal(orderbyte, &order)
		if err != nil {
			fmt.Print("Error unmarshalling")
			return
		}
		orders = append(orders, order)
		return
	}
}

func ReadJson(dirpath string) (data []byte, err error) {
	file, err := os.Open(dirpath)
	if err != nil {
		fmt.Print("Error opening directory")
		return
	}
	defer file.Close()
	data, err = io.ReadAll(file)
	if err != nil {
		fmt.Print("Error reading json")
	}
	return
}
