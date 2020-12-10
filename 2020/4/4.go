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

var NEEDED_FIELDS = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	lines := readLinesFromLine(INPUT_FILE_NAME)
	passports := separatePassports(lines)
	passportWithFields := separatePassportFields(passports)
	exercise1(passportWithFields)
	exercise2(passportWithFields)
}

func exercise1(passports []map[string]string) {
	validPassports := 0
	for _, passport := range passports {
		if processPassport(passport) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func exercise2(passports []map[string]string) {
	validPassports := 0
	for _, passport := range passports {
		if processPassport(passport) && validatePassportFields(passport) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func processPassport(passport map[string]string) bool {
	for _, mandatoryField := range NEEDED_FIELDS {
		_, exists := passport[mandatoryField]
		if !exists {
			return false
		}
	}
	return true
}

func validatePassportFields(passport map[string]string) bool {
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || iyr > 2030 {
		return false
	}
	if !isHeightValid(passport["hgt"]) {
		return false
	}
	if !isHairColorValid(passport["hcl"]) {
		return false
	}
	if !isEyeColorValid(passport["ecl"]) {
		return false
	}
	if len(passport["pid"]) != 9 {
		return false
	}
	return true
}

func isHeightValid(height string) bool {
	re := regexp.MustCompile(`(\d*)(cm|in)`)
	regexResults := re.FindAllStringSubmatch(height, -1)
	if len(regexResults) < 1 {
		return false
	}
	heightRegex := regexResults[0]
	number, err := strconv.Atoi(heightRegex[1])
	if err != nil {
		return false
	}
	unit := heightRegex[2]
	if unit == "cm" {
		if number < 150 || number > 193 {
			return false
		}
	} else if unit == "in" {
		if number < 59 || number > 76 {
			return false
		}
	} else {
		return false
	}
	return true
}

func isHairColorValid(color string) bool {
	re := regexp.MustCompile(`^#([a-fA-F0-9]{6})$`)
	regexResult := re.FindString(color)
	return color == regexResult
}

func isEyeColorValid(color string) bool {
	validColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, exists := validColors[color]
	return exists
}

func separatePassportFields(passports []string) []map[string]string {
	var passportsMap []map[string]string
	re := regexp.MustCompile(`[a-z]{3}:\S+`)
	for _, passport := range passports {
		regexMatches := re.FindAllString(passport, -1)
		onePassportMap := make(map[string]string, len(regexMatches))
		for _, match := range regexMatches {
			key := match[0:3]
			value := match[4:len(match)]
			onePassportMap[key] = value
		}
		passportsMap = append(passportsMap, onePassportMap)
	}
	return passportsMap
}

func separatePassports(lines []string) []string {
	passports := make([]string, 0)
	var processingPassport string
	for _, s := range lines {
		if s == string("") {
			passports = append(passports, processingPassport)
			processingPassport = ""
		} else {
			processingPassport = processingPassport + " " + s
		}
	}
	if processingPassport != "" {
		passports = append(passports, processingPassport)
	}
	return passports
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
