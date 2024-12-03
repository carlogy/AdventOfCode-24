package day

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
	readCount := 0
	for scanner.Scan() {
		readCount++
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
	fmt.Println("Read Count:", readCount)
	return total, nil
}

// func SolvePart2(path string) (int, error) {
// 	total := 0
// 	originalTotal := 0
// 	failedTotal := 0
// 	fmt.Println("report Slice: ", len(reportSlice))
// 	for _, v := range reportSlice {
// 		// count := 0
// 		if v.isValid {
// 			originalTotal++
// 		}
// 		if !v.isValid {
// 			dampnerSlice := make([]int, 0)
// 			reverseSlice := make([]int, 0)

// 			for i := 0; i < len(v.levels)-1; i++ {
// 				currentNum := v.levels[i]
// 				nextNum := v.levels[i+1]
// 				diff := int(math.Abs(float64(currentNum - nextNum)))

// 				if diff <= 3 && diff >= 1 {
// 					// fmt.Println("Diff: ", diff, "\tNum: ", currentNum, "\tNext: ", nextNum, "\tReport: ", v.levels)
// 					dampnerSlice = append(dampnerSlice, v.levels[i])
// 				}

// 			}

// 			reverseSlice = append(reverseSlice, dampnerSlice...)
// 			slices.Reverse(reverseSlice)

// 			if !(len(v.levels)-len(dampnerSlice) > 1) && (slices.IsSorted(dampnerSlice)) || slices.IsSorted(reverseSlice) {
// 				total++
// 				continue
// 			}
// 			failedTotal++
// 		}
// 	}

// 	fmt.Println("Original Total: ", originalTotal, "\nSecond Pass Total: ", total, "\nRemaining failed: ", failedTotal)
// 	fmt.Println(len(reportSlice) - failedTotal - total - originalTotal)

// 	return originalTotal + total, nil

// }
