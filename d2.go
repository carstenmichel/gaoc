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

func processLine(line string) int {
	for index := 0; index < len(line); index++ {
		character := string(line[index])
		processCharacter(character)
	}
	return keypad[xPosition][yPosition]
}
func d2() {
	log.Printf("Day 2\n")
	// Init KeyPad
	keypad = append(keypad, []int{1, 2, 3})
	keypad = append(keypad, []int{4, 5, 6})
	keypad = append(keypad, []int{7, 8, 9})

	file, err := os.Open("input/d2.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
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
}
