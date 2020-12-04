package main

import (
	"bufio"
    "fmt"
	"log"
	"os"
// 	"regexp"
// 	"strings"
// 	"strconv"
 )

const INPUT_FILE_NAME = "input" 
const TREE = "#"

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	exercise_1(lines)
	exercise_2(lines)
}

func exercise_1(lines []string){
	fmt.Println("Exercise 1 - Encountered trees: ", calculateTrees(3,1,lines))
}

func exercise_2(lines []string){
	oneOne := calculateTrees(1,1,lines)
	threeOne := calculateTrees(3,1,lines)
	fiveOne := calculateTrees(5,1,lines)
	sevenOne := calculateTrees(7,1,lines)
	oneTwo := calculateTrees(1,2,lines)
	result := oneOne * threeOne * fiveOne * sevenOne * oneTwo
	fmt.Println("Exercise 2: ", result)
}

func calculateTrees(horizontalMov int, verticalMovement int, lines []string) int{
	var trees int = 0
	var horizontalPos int = 0
	var verticalPos int = 0
	for verticalPos < len(lines)-1{

		horizontalPos = horizontalPos + horizontalMov
		verticalPos = verticalPos + verticalMovement

		if horizontalPos > 30 {
			horizontalPos = horizontalPos - 31
		}

		if string(lines[verticalPos][horizontalPos]) == TREE {
			trees++
		}
	}
	return trees
}

func readLinesFromLine(fileName string) []string {
	file, err := os.Open(fileName) 
  
    if err != nil { 
        log.Fatalf("Can't open the file. Does it exists?")   
	} 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) 

	var lines[] string
	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}
	file.Close()
	return lines
}	