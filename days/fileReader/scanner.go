package filereader

import (
	"bufio"
	"os"
)

func NewScanner(file *os.File) (*bufio.Scanner, error) {

	// file, err := os.Open(path)
	// if err != nil {
	// 	return nil, fmt.Errorf("Experienced %w, while attempting to open the file", err)
	// }

	// defer file.Close()
	newscanner := bufio.NewScanner(file)
	return newscanner, nil
}
