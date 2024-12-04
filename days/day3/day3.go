package day3

import (
	"fmt"
	"os"

	filereader "github.com/carlogy/AdventOfCode-24/days/fileReader"
)

func SolvePart1(path string) (int, error) {

	fmt.Println("Starting day 3")

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Error opening file %w", err)
	}
	defer file.Close()

	scanner, err := filereader.NewScanner(file)
	if err != nil {
		return 0, fmt.Errorf("Error instantiating new scanner: %w", err)
	}

	// re, err := regexp.Compile(`/(mul)\((\d*\,\d*\))/`)

	for scanner.Scan() {

		for i, v := range scanner.Text() {
			switch {
			case string(v) == "m":
				fmt.Println("Match: ", i, string(v))
			case string(v) == "u":
				fmt.Println("Match: ", i, string(v))
			case string(v) == "l":
				fmt.Println("Match: ", i, string(v))
			}
		}
		break
	}

	return 0, nil
}
