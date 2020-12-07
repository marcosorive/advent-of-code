package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const INPUT_FILE_NAME = "input"
const PLANE_ROWS = 127
const PLANE_COLS = 7

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	exercise1(lines)
	exercise2(lines)
}

func exercise1(lines []string) {
	var biggestID int = 0
	for _, s := range lines {
		id := processLine(s)
		if id > biggestID {
			biggestID = id
		}
	}
	fmt.Println("Biggest ID: ", biggestID)
}

func exercise2(lines []string) {
	allIDs := make([]int, len(lines))
	for _, line := range lines {
		allIDs = append(allIDs, processLine(line))
	}
	sort.Ints(allIDs)
	for index, id := range allIDs {
		if index+1 < len(allIDs) && allIDs[index+1] != id+1 && id != 0 {
			fmt.Println("Your seat is ", id+1)
		}
	}
}

func processLine(line string) int {
	row := getRow(line[0:7])
	col := getCol(line[7:len(line)])
	return row*8 + col
}

func getRow(lines string) int {
	rowRange := [2]int{0, PLANE_ROWS}
	for _, char := range lines[0:6] {
		nextRow := float64((rowRange[1] - rowRange[0]) / 2)
		if char == 'F' {
			value := int(math.Floor(nextRow))
			if value < rowRange[0] {
				value = rowRange[0] + value
			}
			rowRange = [2]int{rowRange[0], value}
		} else if char == 'B' {
			value := (int(math.Floor(nextRow)) + 1) + rowRange[0]
			rowRange = [2]int{value, rowRange[1]}
		}
	}

	if lines[6] == 'F' {
		return rowRange[0]
	}
	return rowRange[1]

}

func getCol(chars string) int {
	colRange := [2]int{0, PLANE_COLS}
	for _, char := range chars {
		nextCol := float64((colRange[1] - colRange[0]) / 2)
		if char == 'L' {
			value := int(math.Floor(nextCol))
			if value < colRange[0] {
				value = colRange[0] + value
			}
			colRange = [2]int{colRange[0], value}
		} else if char == 'R' {
			value := (int(math.Ceil(nextCol)) + 1) + colRange[0]
			colRange = [2]int{value, colRange[1]}
		}
	}

	if chars[2] == 'R' {
		return colRange[0]
	}
	return colRange[1]
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
