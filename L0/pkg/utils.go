package pkg

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
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf(" Error reading the file %s", err)
	}
	return bytes, err
}
