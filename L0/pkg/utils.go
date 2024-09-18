package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadJson(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf(" Error opening filepath %s", err)
	}
	byte, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf(" Error reading the file %s", err)
	}
	return byte, err
}
