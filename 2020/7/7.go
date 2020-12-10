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

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	bags := calculateBags(lines)
	fmt.Println(bags)
	exercise1(bags)
	exercise2(bags)
}

func exercise1(bags map[string]map[string]int) {
	bagsContainingShiny := 0
	for bag := range bags {
		if bagContainsShinyGold(bag, bags) {
			bagsContainingShiny++
		}
	}
	fmt.Println("Exercise 1 -", bagsContainingShiny)
}

func exercise2(bags map[string]map[string]int) {
	fmt.Println("Exercise 2 -", calculateContainingBags("shiny gold", bags))
}

func bagContainsShinyGold(bagName string, bagRules map[string]map[string]int) bool {
	bagContent, exists := bagRules[bagName]
	if !exists {
		panic("This bag does not exist :(")
	}
	_, hasShinyGold := bagContent["shiny gold"]
	if hasShinyGold {
		return true
	}
	for bag := range bagContent {
		hasShinyGold = hasShinyGold || bagContainsShinyGold(bag, bagRules)
	}
	return hasShinyGold

}

func calculateContainingBags(bagName string, bagRules map[string]map[string]int) int {
	if len(bagRules[bagName]) == 0 {
		return 0
	}
	bagsContained := 0
	for bag, bagQuantity := range bagRules[bagName] {
		bagsContained = bagsContained + bagQuantity + (bagQuantity * calculateContainingBags(bag, bagRules))
	}
	return bagsContained
}

func calculateBags(rules []string) map[string]map[string]int {
	bagsMap := make(map[string]map[string]int)
	for _, rule := range rules {
		ruleRegex := regexp.MustCompile(`(\w*\s\w*) bags contain (.*)`)
		regexMatches := ruleRegex.FindAllStringSubmatch(string(rule), -1)
		bag := regexMatches[0][1]
		bagsContainedArray := regexMatches[0][2:len(regexMatches[0])]
		bagsContainedRegex := regexp.MustCompile(`(\d) (\w* \w*)`)
		regexMatchesBagsContained := bagsContainedRegex.FindAllStringSubmatch(bagsContainedArray[0], -1)
		bagsContainedMap := make(map[string]int)
		for _, containedBag := range regexMatchesBagsContained {
			numberOfBags, _ := strconv.Atoi(containedBag[1])
			bagsContainedMap[containedBag[2]] = numberOfBags
		}
		bagsMap[bag] = bagsContainedMap
	}
	return bagsMap
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
