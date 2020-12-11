package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE_NAME = "input"

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	numbers := getNumbersFromLines(lines)
	brokenNumber := exercise1(numbers)
	exercise2(numbers, brokenNumber)
}

func exercise1(numbers []int) int {
	wrongN := calcualteFirstWrongNumber(numbers, 25)
	fmt.Println("Exercise 1 - ", wrongN)
	return wrongN
}

func exercise2(numbers []int, broken int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers)+1; j++ {
			numberRange := numbers[i:j]
			sumOfRange := 0
			for _, n := range numberRange {
				sumOfRange = sumOfRange + n
			}
			if sumOfRange == broken {
				min := getMinFromSlice(numberRange)
				max := getMaxFromSlice(numberRange)
				fmt.Println("Exercise 2 - ", min+max)
				return min + max
			}
		}
	}
	fmt.Println("Not found")
	return -1
}

func getMaxFromSlice(v []int) int {
	m := 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func getMinFromSlice(v []int) int {
	m := 99999999999999999
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func calcualteFirstWrongNumber(numbers []int, preamble int) int {
	for index, number := range numbers[preamble:len(numbers)] {
		canBeSumOf := numbers[index : index+preamble]
		isSumOf := isSumOf(canBeSumOf, number)
		if !isSumOf {
			return number
		}
	}
	return -1
}

func isSumOf(numbers []int, n int) bool {
	for _, number1 := range numbers {
		for _, number2 := range numbers {
			if number1 == number2 {
				continue
			}
			if number1+number2 == n {
				return true
			}
		}
	}
	return false
}

func getNumbersFromLines(lines []string) []int {
	numbers := make([]int, 0)
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	return numbers
}

func readLinesFromLine(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Can't open the file. Does it exists?")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}
