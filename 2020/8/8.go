package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const INPUT_FILE_NAME = "input"


type instruction struct{
	operation string
	symbol string
	argument int
}

func main(){
	lines := readLinesFromLine(INPUT_FILE_NAME)
	instructions := getInstructionsFromLine(lines)
	exercise1(instructions)
}

func exercise1(instructions []instruction){
	pointer := 0
	accumulator := 0
	pointerMap := make(map[int]bool)
	for pointer < len(instructions){
		executingInstruction := instructions[pointer]
		_,instructionHasBeenExecuted := pointerMap[pointer]
		if(instructionHasBeenExecuted){
			fmt.Println("Part 1 - Acumulator is:",accumulator)
			return
		}
		pointerMap[pointer] = true
		switch executingInstruction.operation {
		case "acc":
			if executingInstruction.symbol == "+" {
				accumulator = accumulator + executingInstruction.argument
			}else if executingInstruction.symbol == "-" {
				accumulator = accumulator - executingInstruction.argument
			}
			pointer++
		case "jmp":
			if executingInstruction.symbol == "+" {
				pointer = pointer + executingInstruction.argument
			}else if executingInstruction.symbol == "-" {
				pointer = pointer - executingInstruction.argument
			}
		case "nop":
			pointer++
		}
	}
}

func getInstructionsFromLine(lines []string) []instruction{
	instructions := make([]instruction, 0)	
	for _, line := range lines {
		re := regexp.MustCompile(`(\w{3}) (\+|\-)(\d*)`)
		regexResult := re.FindAllStringSubmatch(line,-1)[0]
		number,_ := strconv.Atoi(regexResult[3])
		i := instruction{
			operation: regexResult[1],
			symbol: regexResult[2],
			argument: number}
		instructions = append(instructions, i)
	}
	return instructions
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