package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"sort"
)

// D06LookupTable is a lookup table
type D06LookupTable map[string]int

var allD06Chars [8]D06LookupTable

func processD06Line(line string) {
	for index, char := range line {
		ac := allD06Chars[index]
		s := string(char)
		if value, ok := ac[s]; ok {
			ac[s] = value + 1
		} else {
			ac[s] = 1
		}
	}
}
func d06() (string, string) {
	log.Printf("Day 6\n")

	for index := 0; index < 8; index++ {
		allD06Chars[index] = make(map[string]int)
	}
	file, err := os.Open("input/d06.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		textLine := scanner.Text()
		processD06Line(textLine)
	}

	var result bytes.Buffer
	for index := 0; index < 8; index++ {
		ac := allD06Chars[index]
		var allChars = make(AllCharacters, len(ac))

		for k, v := range ac {
			allChars = append(allChars, KeyValuePair{Character: k, Occurence: v})
		}
		sort.Sort(sort.Reverse(allChars))
		result.WriteString(allChars[0].Character)
	}
	log.Printf("Result is %v\n", result.String())
	firstResult := result.String()

	// second Part of Day 06

	result.Reset()
	for index := 0; index < 8; index++ {
		ac := allD06Chars[index]
		var allChars = make(AllCharacters, len(ac))

		for k, v := range ac {
			allChars = append(allChars, KeyValuePair{Character: k, Occurence: v})
		}
		sort.Sort(allChars)
		for indexx := 0; indexx < len(allChars); indexx++ {
			if allChars[indexx].Occurence > 0 {
				result.WriteString(allChars[indexx].Character)
				break
			}
		}

	}
	log.Printf("2nd Result is %v\n", result.String())
	secondResult := result.String()
	return firstResult, secondResult
}
