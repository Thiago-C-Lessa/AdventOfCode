package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	_PATH_DAY2_ = "../inputs/day2.txt"
)

func readLine2() []string {
	var lines []string
	file, err := os.Open(_PATH_DAY2_)
	if err != nil {
		panic(fmt.Errorf("Error opeing file: %w", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func verifyCrescent(l string) bool {
	splited_line := strings.Split(l, " ")

	for i := 0; i < (len(splited_line) - 1); i++ {
		value_1, err := strconv.ParseInt(splited_line[i], 10, 64)
		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		value_2, err := strconv.ParseInt(splited_line[i+1], 10, 64)

		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		if value_1 > value_2 || ((value_2 - value_1) > 3) {
			return false
		}
	}

	return true
}

func verifyDecrescent(l string) bool {
	splited_line := strings.Split(l, " ")

	for i := 0; i < (len(splited_line) - 1); i++ {
		value_1, err := strconv.ParseInt(splited_line[i], 10, 64)
		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		value_2, err := strconv.ParseInt(splited_line[i+1], 10, 64)

		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		if value_1 < value_2 || ((value_1 - value_2) > 3) {
			return false
		}
	}

	return true
}

func verifyEquality(l string) bool {
	splited_line := strings.Split(l, " ")

	for i := 0; i < (len(splited_line) - 1); i++ {
		value_1, err := strconv.ParseInt(splited_line[i], 10, 64)
		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		value_2, err := strconv.ParseInt(splited_line[i+1], 10, 64)

		if err != nil {
			panic(fmt.Errorf("Error parsing integer :%w", err))
		}

		if value_1 == value_2 {
			return false
		}
	}

	return true
}

func Day2() {
	var safe_count int64
	lines := readLine2()
	for _, line := range lines {
		if (verifyCrescent(line) || verifyDecrescent(line)) && verifyEquality(line) {
			safe_count++
		}
	}

	fmt.Printf("Number of safe lines: %d\n", safe_count)

}
