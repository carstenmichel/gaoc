package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	matrix [][]bool
)

func initMatrix() {
	for i := 0; i < 6; i++ {
		var row []bool
		for j := 0; j < 50; j++ {
			row = append(row, false)
		}
		matrix = append(matrix, row)
	}
}
func printMatrix() {
	for rows := 0; rows < len(matrix); rows++ {
		row := matrix[rows]
		var b bytes.Buffer
		for col := 0; col < len(row); col++ {
			field := row[col]
			if field {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		log.Printf("%v\n", b.String())
	}
}
func processD08Line(in string) {
	if strings.HasPrefix(in, "rect") {
		processD08Rect(in)
	}
	if strings.HasPrefix(in, "rotate column") {
		processD08Col(in)
	}
	if strings.HasPrefix(in, "rotate row") {
		processD08Row(in)
	}

}

func processD08Rect(in string) {
	commandString := in[5:]
	xPos := strings.Index(commandString, "x")
	xString := commandString[:xPos]
	yString := commandString[xPos+1:]
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	for a := 0; a < y; a++ {
		for b := 0; b < x; b++ {
			matrix[a][b] = true
		}
	}
}
func processD08Row(in string) {
	commandString := in[13:]
	xPos := strings.Index(commandString, " by")
	xString := commandString[:xPos]
	yString := commandString[xPos+4:]
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	for shift := 0; shift < y; shift++ {
		carry := matrix[x][49]
		for row := 49; row > 0; row-- {
			matrix[x][row] = matrix[x][row-1]
		}
		matrix[x][0] = carry
	}
}

func processD08Col(in string) {
	commandString := in[16:]
	xPos := strings.Index(commandString, " by")
	xString := commandString[:xPos]
	yString := commandString[xPos+4:]
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	for shift := 0; shift < y; shift++ {
		carry := matrix[5][x]
		for col := 5; col > 0; col-- {
			matrix[col][x] = matrix[col-1][x]
		}
		matrix[0][x] = carry
	}
}
func countPixelInMatrix() int {
	counter := 0
	for rows := 0; rows < len(matrix); rows++ {
		row := matrix[rows]
		for col := 0; col < len(row); col++ {
			field := row[col]
			if field {
				counter = counter + 1
			}
		}
	}
	return counter
}

func d08() int {
	log.Printf("Day 8\n")

	file, err := os.Open("input/d08.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	initMatrix()
	for scanner.Scan() {
		textLine := scanner.Text()
		processD08Line(textLine)

	}

	ledsLit := countPixelInMatrix()
	log.Printf("Leds turned on %v\nCode is", ledsLit)
	printMatrix()
	return ledsLit
}
