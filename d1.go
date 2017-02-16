package main

import (
	"container/list"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Constants for direction
const (
	NORTH = iota
	EAST  = iota
	SOUTH = iota
	WEST  = iota
	LEFT  = 'L'
	RIGHT = 'R'
)

var (
	direction = NORTH
	long      = 0
	lat       = 0
)

func processToken(tok string) {
	log.Printf("Length of token %v is %v\n", tok, len(tok))
	dir := tok[0]
	log.Printf("Parsed direction %v\n", dir)
	rotate(dir)
	steps := tok[1:]
	log.Printf("Parsed %v steps to go \n", steps)
	moveSteps(steps)
}

func moveSteps(st string) {
	numberOfSteps, _ := strconv.Atoi(st)
	switch direction {
	case NORTH:
		long = long + numberOfSteps
	case SOUTH:
		long = long - numberOfSteps
	case WEST:
		lat = lat + numberOfSteps
	case EAST:
		lat = lat - numberOfSteps
	}
}

func rotate(direction byte) {
	switch direction {
	case RIGHT:
		turnRight()
	case LEFT:
		turnLeft()
	default:
		log.Fatalf("Direction %v is not supported\n", direction)
	}
}

func turnRight() {
	log.Print("Turn right")
	switch direction {
	case NORTH:
		direction = EAST
	case EAST:
		direction = SOUTH
	case SOUTH:
		direction = WEST
	case WEST:
		direction = NORTH
	default:
		log.Fatalf("Unknown direction right state: %v\n", direction)
	}
}

func turnLeft() {
	log.Print("Turn Left")
	switch direction {
	case NORTH:
		direction = WEST
	case EAST:
		direction = NORTH
	case SOUTH:
		direction = EAST
	case WEST:
		direction = SOUTH
	default:
		log.Fatalf("Unknown direction left state: %v\n", direction)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type pair struct {
	long int
	lat  int
}

func d1Part1() {

	log.Printf("Day 1\n")
	dat, err := ioutil.ReadFile("input/d1.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	commandstring := string(dat)

	mylist := list.New()
	commands := strings.Split(commandstring, ",")
	for _, element := range commands {
		processToken(strings.Trim(element, " \n\r"))
		mylist.PushBack(pair{long: long, lat: lat})
	}

	log.Printf("Direction is now %v\n", direction)
	log.Printf("Long is %v, Lat is %v\n", long, lat)
	res := abs(long) + abs(lat)
	log.Printf("Day1 Puzzle Part 1 result is  %v\n", res)
}
