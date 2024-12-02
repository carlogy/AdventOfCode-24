package filereader

import (
	"fmt"
	"os"
)

const InputFilePath = "./inputFiles/"

func ReadFile(path string) ([]byte, error) {

	if path == "" {
		return nil, fmt.Errorf("Empty string as path supplied: %s", path)
	}
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Error attempting to read file. Error: %w\n", err)
	}

	return data, nil
}
