package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	filereader "github.com/carlogy/AdventOfCode-24/days/fileReader"
)

func SolvePart1(path string) (int, error) {

	total := 0

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

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {

		mulEx := re.FindAllStringSubmatch(scanner.Text(), -1)

		for _, v := range mulEx {

			result := getResult(v[0])

			total += result
		}

	}
	return total, nil

}

func mul(num1, num2 int) int {
	return num1 * num2
}

func getResult(mulExpression string) int {
	var result int
	r := regexp.MustCompile(`(\d+)`)

	nums := r.FindAllString(mulExpression, -1)

	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	result = mul(num1, num2)
	return result
}
