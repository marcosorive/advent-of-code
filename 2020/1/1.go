package main

import (
	"bufio"
    "fmt"
	"log"
	"strconv"
    "os"
)

const FILE_NAME = "1_input"

func main(){
	numbers := readFileNumberByLine(FILE_NAME)
	exercise_1(numbers)
	exercise_2(numbers)
}

func exercise_1(numbers []int){
	var result int
	for i, s := range numbers{
		for j := i; j<len(numbers); j++{
				if s + numbers[j] == 2020 {
				result = s * numbers[j]
			}
		}
	}
	fmt.Println("Result is for exercise 1 is", result)
}

func exercise_2(numbers []int){
	var result int
	for i, s := range numbers{
		for j:= i+1; j<len(numbers); j++{
			for k:= j+1; k<len(numbers); k++{
				if s + numbers[j] + numbers[k] == 2020 {
					result = s * numbers[j] * numbers[k]
				}
			}
		}
	}
	fmt.Println("Result for exercise 2 is", result)
}

func readFileNumberByLine(fileName string) []int {
	file, err := os.Open(fileName) 
  
    if err != nil { 
        log.Fatalf("Can't open the file. Does it exists?")   
	} 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) 

	var numbers[] int
	for scanner.Scan() {
		i, error :=  strconv.Atoi(scanner.Text())
		if error == nil {
			numbers = append(numbers,i)
		}
	}
	file.Close()
	return numbers
}	