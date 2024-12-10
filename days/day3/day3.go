package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

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

func getMulEx(str string) [][]string {

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	mulExpr := r.FindAllStringSubmatch(str, -1)

	return mulExpr
}

func SolvePart2(path string) (int, error) {

	do := "do()"
	dont := "don't()"

	total := 0
	// reNo := regexp.MustCompile(`(don't\(\))`)
	// reDo := regexp.MustCompile(`(do\(\))`)

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("Error opening file %w", err)
	}
	defer file.Close()
	scanner, err := filereader.NewScanner(file)
	if err != nil {
		return 0, fmt.Errorf("Error instantiating new scanner: %w", err)
	}

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		str := scanner.Text()

		if strings.Contains(str, dont) {
			fmt.Println("Dont Str check: ", str)
			dontSplit := splitByDont(str)

			for i, v := range dontSplit {
				if i == 0 {
					mulExpr := getMulEx(v)

					for j := 0; j < len(mulExpr); j++ {
						result := getResult(mulExpr[j][0])
						total += result
					}
				}

				if strings.Contains(v, do) {
					fmt.Println("Str:\t\t", str)
					fmt.Println("SubStr dont has do:\n", v)
					doSplit := splitByDo(v)
					fmt.Println("Do split:\t", doSplit)

					for i, v := range doSplit {
						if i > 0 && !strings.Contains(v, dont) {
							mulExpr := getMulEx(v)

							for j := 0; j < len(mulExpr); j++ {
								result := getResult(mulExpr[j][0])
								total += result
							}
						}

					}

				}

			}

		}

		// if dont not present

		mulExpr := getMulEx(str)

		for i := 0; i < len(mulExpr); i++ {
			result := getResult(mulExpr[i][0])
			total += result
		}

		// if strings.Count(str, do) == 0 {
		// 	// fmt.Println("No do, checking dont next...")

		// 	dontSplit := splitByDont(str)
		// 	// fmt.Println(dontSplit)

		// 	for i, v := range dontSplit {
		// 		if i == 0 {
		// 			mulExpr := getMulEx(v)

		// 			for j := 0; j < len(mulExpr); j++ {
		// 				result := getResult(mulExpr[j][0])
		// 				total += result
		// 				// fmt.Println(result, mulExpr[j][0])
		// 			}
		// 		}
		// 	}

		// 	if strings.Count(str, "do()") > 0 {
		// 		fmt.Println("Do is in string, check dont's")
		// 		fmt.Println(str)
		// 		break
		// 	}

		// }

	}

	return total, nil
}

func splitByDont(str string) []string {

	return strings.Split(str, "don't()")

}

func splitByDo(str string) []string {
	return strings.Split(str, "do()")
}

func SolvePartTwo(str string) (int, error) {
	total := 0
	do := "do()"
	dont := "don't()"

	dontCount := strings.Count(str, dont)
	doCount := strings.Count(str, do)

	if doCount == 0 && dontCount == 0 {
		mulExpr := getMulEx(str)

		for i := 0; i < len(mulExpr); i++ {
			result := getResult(mulExpr[i][0])
			total += result
		}
	}
	//to do solve for strings that have dont and do in them

	// switch {
	// case !strings.Contains(str, dont) && !strings.Contains(str, do):

	// 	mulExpr := getMulEx(str)

	// 	for i := 0; i < len(mulExpr); i++ {
	// 		result := getResult(mulExpr[i][0])
	// 		total += result
	// 	}

	// case strings.Contains(str, dont) && strings.Contains(str, do):
	// 	dontSplit := splitByDont(str)

	// 	for _, v := range dontSplit {
	// 		if strings.Contains(v, do) {
	// 			doSplit := splitByDo(v)

	// 			if len(doSplit) > 1 {
	// 				for j := 0; j < len(doSplit); j++ {
	// 					mulExpr := getMulEx(doSplit[j])

	// 					result := getResult(mulExpr[j][0])
	// 					total += result
	// 				}

	// 			}

	// 		}
	// 	}

	// case !strings.Contains(str, dont):
	// 	mulExpr := getMulEx(str)

	// 	for i := 0; i < len(mulExpr); i++ {
	// 		result := getResult(mulExpr[i][0])
	// 		total += result
	// 	}

	// }

	return total, nil
}
