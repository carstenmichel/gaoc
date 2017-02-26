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
