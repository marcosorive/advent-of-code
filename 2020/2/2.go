package main

import (
	"bufio"
    "fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

const INPUT_FILE_NAME = "2_input" 

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	exercise_1(lines)
	exercise_2(lines)
}

func exercise_1(lines []string){
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	var result int = 0
	for i := range lines {
		regexMatches := re.FindAllStringSubmatch(lines[i], -1)
		minTimes, err := strconv.Atoi(regexMatches[0][1])
		if(err != nil){
			panic("Converting min time failed!")
		}
		maxTimes, err := strconv.Atoi(regexMatches[0][2])
		if(err != nil){
			panic("Converting max time failed!")
		}
		letterToMatch := regexMatches[0][3]
		stringToCheck := regexMatches[0][4]
		times := strings.Count(stringToCheck, letterToMatch)
		if ((minTimes <= times) && (times <= maxTimes)) {
			result++
		}
	}
	fmt.Println("Exercise 1 - Number of correct passwords is:", result)
}

func exercise_2(lines []string){
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	var result int = 0
	for i := range lines {
		regexMatches := re.FindAllStringSubmatch(lines[i], -1)
		firstPos, err := strconv.Atoi(regexMatches[0][1])
		if(err != nil){
			panic("Converting firstPos failed!")
		}
		lastPos, err := strconv.Atoi(regexMatches[0][2])
		if(err != nil){
			panic("Converting lasPost failed!")
		}
		letterToMatch := regexMatches[0][3]
		stringToCheck := regexMatches[0][4]
		firstLetterMatches := string(stringToCheck[firstPos-1]) == letterToMatch
		lastLetterMatches :=  string(stringToCheck[lastPos-1]) == letterToMatch
		if ((lastPos-1 < len(stringToCheck)) && (firstLetterMatches || lastLetterMatches) && !(firstLetterMatches && lastLetterMatches)){	
			result++
		}
	}
	fmt.Println("Exercise 2 - Number of correct passwords: ", result)
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