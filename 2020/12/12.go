package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
)

const INPUT_FILE_NAME = "input"
const EAST = 'E'
const WEST = 'W'
const SOUTH = 'S'
const NORTH = 'N'

func main() {
	lines := readLinesFromFile(INPUT_FILE_NAME)
	exercise1(lines)
	exercise2(lines)
}

func exercise1(lines []string) {
	facingDirection := EAST
	movements := make(map[rune]int)
	for _, line := range lines {
		command := rune(line[0])
		quantity,_ := strconv.Atoi(line[1:len(line)])
		if command == 'R' || command == 'L' {
			newDirection := calculateFacingDirection(facingDirection, line)
			fmt.Println("Command is",line, "From" , string(facingDirection), "to", string(newDirection))
			facingDirection = rune(newDirection)
		} else if command == 'F' {
			movements[facingDirection] = movements[facingDirection] + quantity
		} else {
			movements[command] = movements[command] + quantity
		}
	}
	manhattanDistance := calculateManhattanDistance(movements)
	fmt.Println("Exercise 1 - ", manhattanDistance)
}

func calculateFacingDirection(nowFacing rune, instruction string) rune {
	command := rune(instruction[0])
	quantity,_ := strconv.Atoi(instruction[1:len(instruction)])
	cardinals := [4]rune{NORTH, EAST, SOUTH, WEST}
	numberOfTurns := quantity/90
	cardinalIndex := getCardinalIndex(nowFacing, cardinals)
	indexOfNewDirection := 0
	if command == 'R' {
		indexOfNewDirection = cardinalIndex + numberOfTurns
		if (indexOfNewDirection > 3) {
			indexOfNewDirection = indexOfNewDirection - 4
		}
	} else if command == 'L' {
		indexOfNewDirection = cardinalIndex - numberOfTurns
		if indexOfNewDirection < 0 {
			indexOfNewDirection = indexOfNewDirection + 4
		}
	}
	return cardinals[indexOfNewDirection]
}

func getCardinalIndex(cardinal rune, allCardinals [4]rune) int{
	for i, c := range allCardinals {
		if c == cardinal {
			return i
		}
	}
	return -1
}

func exercise2(lones []string) {
	waypoint := map[string]int{"E":10, 	"N":1}
	ship := map[rune]int
	for _, line := lines {
		command := rune(line[0])
		quantity,_ := strconv.Atoi(line[1:len(line)])
		if command == 'F' {
			ship := shipMovementWithWaypoint(quantity, ship, waypoint)
		} else if command == 'R' | command == 'L' {
			waypoint = rotateWayPoint(quantity, waypoint)
		} else {
			waypoint = moveWaypoint(quantity, waypoint)
		}
	}
	manhattanDistance := calculateManhattanDistance(ship)
	fmt.Println("Exercise 2 - ", manhattanDistance)
}

func shipMovementWithWaypoint(quantity int,ship map[string]int, waypoint map[rune]int) map[rune]int {
	waypointQuantity := make(map[rune]int)
	for key, element := range waypoint {
		waypointQuantity[key] = element*quantity
	}
	for 
}

func calculateManhattanDistance(movements map[rune]int) int {
	verticalMov := int(math.Abs(float64(movements[NORTH] - movements[SOUTH])))
	horizontalMov := int(math.Abs(float64(movements[EAST] - movements[WEST])))
	return verticalMov + horizontalMov
}

func readLinesFromFile(fileName string) []string {
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
