package main

import (
	"bufio"
	"bytes"
	"container/list"
	"log"
	"os"
	"strconv"
)

var (
	keypad    [][]int
	bigKeypad [][]string
	xPosition = 1
	yPosition = 1
)

func decX() {
	if xPosition > 0 {
		xPosition = xPosition - 1
	}
}
func incX() {
	if xPosition < 2 {
		xPosition = xPosition + 1
	}
}

func decY() {
	if yPosition > 0 {
		yPosition = yPosition - 1
	}
}
func incY() {
	if yPosition < 2 {
		yPosition = yPosition + 1
	}
}

/* Here comes the special logic for the extended bigKeypad
    1
  2 3 4
5 6 7 8 9
  A B C
    D
*/
func minPosition(pos int) int {
	var retc = 0
	switch pos {
	case 0, 4:
		retc = 2
	case 1, 3:
		retc = 1
	case 2:
		retc = 0
	}
	return retc
}
func maxPosition(pos int) int {
	var retc = 0
	switch pos {
	case 0, 4:
		retc = 2
	case 1, 3:
		retc = 3
	case 2:
		retc = 4
	}
	return retc
}
func decBigX() {
	minXPosition := minPosition(yPosition)

	if xPosition > minXPosition {
		xPosition = xPosition - 1
	}
}
func incBigX() {
	maxXPosition := maxPosition(yPosition)
	if xPosition < maxXPosition {
		xPosition = xPosition + 1
	}
}

func decBigY() {
	minYPos := minPosition(xPosition)
	if yPosition > minYPos {
		yPosition = yPosition - 1
	}
}
func incBigY() {
	maxYPoisition := maxPosition(xPosition)
	if yPosition < maxYPoisition {
		yPosition = yPosition + 1
	}
}
func processCharacter(character string) {
	switch character {
	case "U":
		decX()
	case "D":
		incX()
	case "L":
		decY()
	case "R":
		incY()
	default:
		log.Fatalf("something went wrong with character %v /n", character)
	}
}

func processCharacterBig(character string) {
	switch character {
	case "U":
		decBigX()
	case "D":
		incBigX()
	case "L":
		decBigY()
	case "R":
		incBigY()
	default:
		log.Fatalf("something went wrong with character %v /n", character)
	}
}

func processLine(line string) int {
	for index := 0; index < len(line); index++ {
		character := string(line[index])
		processCharacter(character)
	}
	return keypad[xPosition][yPosition]
}

func processLineBig(line string) string {
	for index := 0; index < len(line); index++ {
		character := string(line[index])
		processCharacterBig(character)
	}
	return bigKeypad[xPosition][yPosition]
}

func d2() string {
	log.Printf("Day 2\n")
	// Init KeyPad
	keypad = append(keypad, []int{1, 2, 3})
	keypad = append(keypad, []int{4, 5, 6})
	keypad = append(keypad, []int{7, 8, 9})

	file, err := os.Open("input/d2.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	resultList := list.New()
	for scanner.Scan() {
		textLine := scanner.Text()
		result := processLine(textLine)
		resultList.PushBack(result)
	}
	var bu bytes.Buffer
	for element := resultList.Front(); element != nil; element = element.Next() {
		bu.WriteString(strconv.Itoa(element.Value.(int)))
	}
	log.Printf("%v is the result of Part 1\n", bu.String())
	return bu.String()
}

func d2Part2() string {
	log.Printf("Day 2 Part 2 \n")
	bigKeypad = append(bigKeypad, []string{"X", "X", "1", "X", "X"})
	bigKeypad = append(bigKeypad, []string{"X", "2", "3", "4", "X"})
	bigKeypad = append(bigKeypad, []string{"5", "6", "7", "8", "9"})
	bigKeypad = append(bigKeypad, []string{"X", "A", "B", "C", "X"})
	bigKeypad = append(bigKeypad, []string{"X", "X", "D", "X", "X"})

	// Starting at virtual position 5

	xPosition = 0
	yPosition = 2
	file, err := os.Open("input/d2.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	resultList := list.New()
	for scanner.Scan() {
		textLine := scanner.Text()
		result := processLineBig(textLine)
		resultList.PushBack(result)
	}
	var bu bytes.Buffer
	for element := resultList.Front(); element != nil; element = element.Next() {
		bu.WriteString(element.Value.(string))
	}
	log.Printf("%v is the result of Part 1\n", bu.String())
	return bu.String()
}
