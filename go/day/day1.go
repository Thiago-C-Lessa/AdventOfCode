package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	_PATH_DAY1_ = "../inputs/day1.txt"
)

func readLine1() []string {
	file, err := os.Open(_PATH_DAY1_)
	defer file.Close()
	if err != nil {
		panic(fmt.Errorf("Error openng file: %v", err))
	}
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}

func Day1() {
	var lines []string = readLine1()
	var distances int64 = 0
	var iterations int64 = 0
	column1 := make([]int32, 1000)
	column2 := make([]int32, 1000)

	for _, line := range lines {
		iterations++

		splited := strings.Split(line, "   ")

		val1, err1 := strconv.ParseInt(splited[0], 10, 32)
		val2, err2 := strconv.ParseInt(splited[1], 10, 32)

		if err1 != nil || err2 != nil {
			continue // skip lines with non-integer values
		}

		column1 = append(column1, int32(val1))
		column2 = append(column2, int32(val2))

		//            if i<5{   number  := int8(char - '0') // rune é a tabela asc, 0 é 48. -'0' é o número da asc - 48
		/*
			for i, char := range line {

				if i < 5 {
					number := int8(char - '0') // rune é a tabela asc, 0 é 48. -'0' é o número da asc - 48

					column1 = append(column1, number)
				} else if i < 8 {
					continue
				} else {
					number := int8(char - '0')

					column2 = append(column2, number)
				}


			}
		*/
	}

	sort.Slice(column1, func(i, j int) bool {
		return column1[i] < column1[j]
	})

	sort.Slice(column2, func(i, j int) bool {
		return column2[i] < column2[j]
	})

	for i, _ := range column1 {
		distance := int64(column1[i] - column2[i])

		if distance < 0 {
			distance *= -1
		}
		distances += int64(distance)
	}

	fmt.Printf("column1:%v column2:%v\n", column1, column2)
	fmt.Printf("it:%v, distance:%v\n", iterations, distances)

	fmt.Printf("Distances sum: %v\n", distances)
}
