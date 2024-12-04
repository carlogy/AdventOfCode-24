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
		intLevels := make([]int, len(str))

		for i, v := range str {

			if num, err := strconv.Atoi(v); err == nil {
				intLevels[i] = num
			}
		}

		reverseSlice := make([]int, len(intLevels))
		_ = copy(reverseSlice, intLevels)
		slices.Reverse(reverseSlice)

		compactLevel := make([]int, len(intLevels))
		_ = copy(compactLevel, intLevels)
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

func SolvePart2(path string) (int, error) {
	rp := ReportsCount{}
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Error opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		levels := converToIntSlice(str)

		scrubbedLevels := removeBadLevels(levels)
		sliceLenDiff := len(levels) - len(scrubbedLevels)

		switch {
		case slices.IsSorted(scrubbedLevels) && sliceLenDiff == 1:
			rp.Passed++
			continue
		case !slices.IsSorted(scrubbedLevels):
			reverseSlice := make([]int, 0)
			reverseSlice = append(reverseSlice, scrubbedLevels...)
			slices.Reverse(reverseSlice)
			if slices.IsSorted(reverseSlice) {
				rp.Passed++
				continue
			}

			fmt.Println("Hey I'm not sorted! ", reverseSlice, "OG Slice: ", scrubbedLevels)
			break
		}

	}
	fmt.Println(rp.Failed)
	return rp.Passed, nil
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

	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice...)

	for i := 1; i < len(newSlice); i++ {
		currentNum := newSlice[i]
		prevNum := newSlice[i-1]
		absDiff := int(math.Abs(float64(currentNum) - float64(prevNum)))

		if absDiff == 0 {
			slices.Delete(newSlice, i-1, i)
		}

		if absDiff > 3 {
			slices.Delete(newSlice, i-1, i)
		}
	}

	scrubbedSlice := make([]int, 0)
	for _, v := range newSlice {
		if v != 0 {
			scrubbedSlice = append(scrubbedSlice, v)
		}
	}
	return scrubbedSlice
}

// figure out final scrubb
func sortFinalCount(slice []int) ([]int, int, bool) {
	copySlice := make([]int, 0)
	copySlice = append(copySlice, slice...)
	swappedCount := 0
	dupExists := false
	freqMap := make(map[int]int)

	return copySlice, swappedCount, dupExists
}
