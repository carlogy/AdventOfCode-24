package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type levelReports struct {
	levels  []int
	isValid bool
}

type ReportsCount struct {
	Passed int
	Failed int
	Total  int
}

func NewlevelReports(reports []int, isvalid bool) *levelReports {
	return &levelReports{
		levels:  reports,
		isValid: isvalid,
	}
}

var reportSlice = make([]levelReports, 0)

func SolvePart1(path string) (int, error) {

	fmt.Println("Starting day 2")

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Error opening file for reading: %w", err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")

		intLevels := converToIntSlice(str)

		reverseSlice := copySlice(intLevels)
		slices.Reverse(reverseSlice)

		compactLevel := copySlice(intLevels)
		compactLevel = slices.Compact(compactLevel)

		if len(compactLevel) == len(intLevels) && (slices.IsSorted(intLevels) || slices.IsSorted(reverseSlice)) {

			isSafe := true
			for i := 0; i < len(intLevels)-1; i++ {
				num := intLevels[i]
				next := intLevels[i+1]
				diff := int64(math.Abs(float64(num - next)))

				if !(diff >= 1 && diff <= 3) {
					isSafe = false
					break
				}
			}
			if isSafe {
				reportSlice = append(reportSlice, *NewlevelReports(intLevels, true))
				total++
				continue
			}
		}
		reportSlice = append(reportSlice, *NewlevelReports(intLevels, false))
	}

	return total, nil
}

func converToIntSlice(slice []string) []int {
	levels := make([]int, len(slice))

	for i, v := range slice {
		num, err := strconv.Atoi(v)
		if err == nil {
			levels[i] = num
		}
	}
	return levels
}

func removeBadLevels(slice []int) []int {

	newSlice := copySlice(slice)

	for i := 1; i < len(newSlice); i++ {
		currentNum := newSlice[i]
		prevNum := newSlice[i-1]
		absDiff := int(math.Abs(float64(currentNum) - float64(prevNum)))

		if absDiff == 0 {
			newSlice = slices.Delete(newSlice, i-1, i)
		}

		if absDiff > 3 {
			newSlice = slices.Delete(newSlice, i-1, i)
		}

	}

	scrubbedSlice := make([]int, 0)
	for _, v := range newSlice {
		if v != 0 {
			scrubbedSlice = append(scrubbedSlice, v)
		}
	}
	noDupSlice := removeDuplicates(scrubbedSlice)

	return noDupSlice
}

func removeDuplicates(slice []int) []int {

	newSlice := copySlice(slice)

	freqMap := make(map[int]int)

	for _, v := range newSlice {
		freqMap[v]++
	}

	scrubbedSlice := make([]int, 0)
	for i, v := range freqMap {
		if v <= 1 {
			scrubbedSlice = append(scrubbedSlice, i)
		}
	}
	return scrubbedSlice
}

func copySlice(slice []int) []int {
	newslice := make([]int, 0)
	newslice = append(newslice, slice...)

	return newslice
}

// figure out final sort
func sortFinalCount(slice []int) ([]int, int) {

	copy := copySlice(slice)

	swapCount := 0

	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(copy); i++ {
			if copy[i] < copy[i-1] {
				swapped = true
				swapCount++
				// copySlice[i-1], copySlice[i] = copySlice[i], copySlice[i-1]
				copy = slices.Delete(copy, i, i+1)
				// fmt.Println("Swap:\t", copySlice)
			}
		}
	}

	if len(copy) <= 1 {

		reverseSlice := copySlice(slice)
		slices.Reverse(reverseSlice)

		newSlice, swappedCount := sortFinalCount(reverseSlice)

		return newSlice, swappedCount
	}

	return copy, swapCount
}

func SplitSlice(slice []int) ([]int, []int) {
	pivot := int((len(slice) / 2))

	half1 := make([]int, 0)
	half1 = append(half1, slice[:pivot]...)

	half2 := make([]int, 0)
	half2 = append(half2, slice[pivot:]...)

	return half1, half2
}

func SolvePart2(path string) (int, error) {

	rp := ReportsCount{}

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %w", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		intSlice := converToIntSlice(strings.Split(scanner.Text(), " "))

		sortedSlice, swapCount := sortFinalCount(intSlice)

		newSlice := make([]int, 0)

		if swapCount <= 1 {
			dupScrub := removeDuplicates(sortedSlice)
			newSlice = append(newSlice, dupScrub...)
			// fmt.Println(sortedSlice, newSlice, intSlice)

		}

		lenDiff := int(math.Abs(float64(len(intSlice) - len(newSlice))))

		if lenDiff == 1 {

			rp.Passed++
			continue
		}
		rp.Failed++

	}
	rp.Total = rp.Passed + rp.Failed
	fmt.Println("Passed:\t", rp.Passed, "\nFailed:\t", rp.Failed, "\nTotal:\t", rp.Total)
	return 0, nil
}
