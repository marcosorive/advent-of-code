package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const INPUT_FILE_NAME = "input"

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	exercise1(lines)
	exercise2(lines)
}

func exercise1(lines []string) {
	totalLines := 0
	allDeclarations := convertDeclarationToString(lines)
	for _, line := range allDeclarations {
		totalLines = totalLines + processDeclarationByAnyone(line)
	}
	fmt.Println("Total lines", totalLines)
}

func processDeclarationByAnyone(declaration string) int {
	questionMap := make(map[string]bool)
	for _, char := range declaration {
		if !questionMap[string(char)] {
			questionMap[string(char)] = true
		}
	}
	return len(questionMap)
}

func exercise2(lines []string) {
	totalLines := 0
	allDeclarations := convertDeclarationToArrayOfStrings(lines)
	for _, line := range allDeclarations {
		totalLines = totalLines + processDeclarationByEveryone(line)
	}
	fmt.Println("Total lines", totalLines)
}

func processDeclarationByEveryone(declaration []string) int {
	answers := 0
	// Iterating a single anwsers from the first person
	for _, singleAnswer := range declaration[0] {
		// Looking if that answer is in the other person's declaration.
		everyoneSaysYes := true
		for _, otherAnswer := range declaration[1:] {
			if !strings.Contains(otherAnswer, string(singleAnswer)) {
				everyoneSaysYes = false
			}
		}
		if everyoneSaysYes {
			answers++
		}
	}
	return answers
}

func deleteSpacesFromString(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func convertDeclarationToString(lines []string) []string {
	declarations := make([]string, 0)
	var processingDeclaration string
	for i, s := range lines {
		if s == string("") || i == len(lines) {
			noSpaceDeclaration := deleteSpacesFromString(processingDeclaration)
			declarations = append(declarations, noSpaceDeclaration)
			processingDeclaration = ""
		} else {
			processingDeclaration = processingDeclaration + " " + s
		}
	}
	if processingDeclaration != "" {
		noSpaceDeclaration := deleteSpacesFromString(processingDeclaration)
		declarations = append(declarations, noSpaceDeclaration)
	}
	return declarations
}

func convertDeclarationToArrayOfStrings(lines []string) [][]string {
	declarations := make([][]string, 0)
	var processingDeclaration []string
	for i, s := range lines {
		if s == string("") || i == len(lines) {
			declarations = append(declarations, processingDeclaration)
			processingDeclaration = nil
		} else {
			processingDeclaration = append(processingDeclaration, s)
		}
	}
	if processingDeclaration != nil {
		declarations = append(declarations, processingDeclaration)
	}
	return declarations
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
