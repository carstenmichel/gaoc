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
	direction    = NORTH
	long         = 0
	lat          = 0
	mylist       = list.New()
	alreadyFound = false
)

func processToken(tok string) {
	dir := tok[0]
	rotate(dir)
	steps := tok[1:]
	moveSteps(steps)
}

func moveSteps(st string) {
	numberOfSteps, _ := strconv.Atoi(st)
	for index := 0; index < numberOfSteps; index++ {

		switch direction {
		case NORTH:
			long = long + 1
		case SOUTH:
			long = long - 1
		case WEST:
			lat = lat + 1
		case EAST:
			lat = lat - 1
		}

		//	log.Printf("I am here: long %v lat %v\n", long, lat)
		checkIfIWasAlreadyHere(long, lat)

		mylist.PushBack(pair{long: long, lat: lat})
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

func checkIfIWasAlreadyHere(long, lat int) {
	if alreadyFound == false {
		for element := mylist.Front(); element != nil; element = element.Next() {
			p := element.Value
			if (p.(pair).lat == lat) && (p.(pair).long == long) {
				away := abs(long) + abs(lat)
				log.Printf("Shortpath (2nd part of puzzle), %v steps away\n", away)
				alreadyFound = true
			}
		}
	}
}

func d1Part1() {
	mylist.PushBack(pair{long: 0, lat: 0})
	log.Printf("Day 1\n")
	dat, err := ioutil.ReadFile("input/d1.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	commandstring := string(dat)

	commands := strings.Split(commandstring, ",")
	for _, element := range commands {
		processToken(strings.Trim(element, " \n\r"))
	}

	res := abs(long) + abs(lat)
	log.Printf("Day 1 Puzzle Part 1 result is  %v\n", res)
}
