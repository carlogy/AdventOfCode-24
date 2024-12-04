package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	filereader "github.com/carlogy/AdventOfCode-24/days/fileReader"
)

func SolvePart1(path string) (int, error) {

	fmt.Println("Starting day 1")

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Experienced %w, while attempting to open the file", err)
	}

	defer file.Close()

	scanner, err := filereader.NewScanner(file)
	if err != nil {
		return 0, err
	}

	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "   ")
		i, err := strconv.Atoi(splitString[0])
		if err == nil {
			list1 = append(list1, i)
		}

		j, err := strconv.Atoi(splitString[1])
		if err == nil {
			list2 = append(list2, j)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	diffSlice := make([]int, len(list1))
	total := 0

	for i := 0; i < len(list1); i++ {
		diffInt := int(math.Abs(float64(list1[i] - list2[i])))
		diffSlice[i] = diffInt
		total += diffInt
	}

	return total, nil
}

func SolvePart2(path string) (int, error) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	set := make(map[int]bool)
	repeatedNums := make(map[int]int)

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Experienced %w, while attempting to open the file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "   ")
		i, err := strconv.Atoi(splitString[0])
		if err == nil {
			list1 = append(list1, i)
			set[i] = true
		}

		j, err := strconv.Atoi(splitString[1])
		if err == nil {
			list2 = append(list2, j)
		}
	}

	for _, num := range list2 {
		if set[num] {
			repeatedNums[num]++
		}
	}

	total := 0

	for k, v := range repeatedNums {
		total += k * v
	}

	return total, nil
}
