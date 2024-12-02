package main

import (
	"fmt"

	d "github.com/carlogy/AdventOfCode-24/days/day1"
	filereader "github.com/carlogy/AdventOfCode-24/days/fileReader"
)

func main() {
	fmt.Println("It's that time of the year again!\nAdvent of Code 2024")
	day := 1
	fullPath := fmt.Sprintf("%sday%dInput.txt", filereader.InputFilePath, day)
	solution1, err := d.SolvePart1(fullPath)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Part 1: %d\n", solution1)

	solution2, err := d.SolvePart2(fullPath)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Printf("Part 2: %d\n", solution2)
}
