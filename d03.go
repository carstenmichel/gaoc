package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	impossibleTriangles = 0
	impossibleTriples   = 0
)

func processTriangles(line string) {
	s := strings.Fields(line)
	var numbers = []int{}

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}

	sort.Ints(numbers)
	if (numbers[0] + numbers[1]) > numbers[2] {
		impossibleTriangles = impossibleTriangles + 1
	}

}
func d03() int {
	log.Printf("Day 3\n")

	file, err := os.Open("input/d03.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		textLine := scanner.Text()
		processTriangles(textLine)

	}
	log.Printf("Number of Impossible Triangles %v\n", impossibleTriangles)
	return impossibleTriangles
}

func processTriples(sc *bufio.Scanner) {
	lineOne := sc.Text()
	sc.Scan()
	lineTwo := sc.Text()
	sc.Scan()
	lineThree := sc.Text()
	ones := strings.Fields(lineOne)
	twos := strings.Fields(lineTwo)
	threes := strings.Fields(lineThree)

	var numones = []int{}
	var numtwos = []int{}
	var numthrees = []int{}

	for index := 0; index < 3; index++ {
		o, _ := strconv.Atoi(ones[index])
		t, _ := strconv.Atoi(twos[index])
		th, _ := strconv.Atoi(threes[index])
		switch index {
		case 0:
			numones = append(numones, o)
			numones = append(numones, t)
			numones = append(numones, th)
		case 1:
			numtwos = append(numtwos, o)
			numtwos = append(numtwos, t)
			numtwos = append(numtwos, th)
		case 2:
			numthrees = append(numthrees, o)
			numthrees = append(numthrees, t)
			numthrees = append(numthrees, th)
		}
	}
	sort.Ints(numones)
	sort.Ints(numtwos)
	sort.Ints(numthrees)
	if (numones[0] + numones[1]) > numones[2] {
		impossibleTriples = impossibleTriples + 1
	}
	if (numtwos[0] + numtwos[1]) > numtwos[2] {
		impossibleTriples = impossibleTriples + 1
	}
	if (numthrees[0] + numthrees[1]) > numthrees[2] {
		impossibleTriples = impossibleTriples + 1
	}

}

func d03Part2() int {
	log.Printf("Day3 Part 2\n")
	file, err := os.Open("input/d03.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		processTriples(scanner)

	}
	log.Printf("Number of Impossible Triples %v\n", impossibleTriples)
	return impossibleTriples

}
