package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	countd09         = 0
	countd09Version2 = 0
)

func processCompressed09(r *bufio.Reader) {
	var b bytes.Buffer
	for {
		rune, _, err := r.ReadRune()
		if err != nil {
			log.Fatalf("Failed to read rune %v\n", err)
		}
		if rune == 0x29 {
			break
		} else {
			b.WriteRune(rune)
		}
	}
	headerString := b.String()
	//log.Printf("Header is %v\n", headerString)
	elements := strings.Split(headerString, "x")
	amount, _ := strconv.Atoi(elements[1])
	length, _ := strconv.Atoi(elements[0])
	// read amount number of runes
	countd09 = countd09 + (amount * length)

	var otherBuffer bytes.Buffer
	for index := 0; index < length; index++ {
		rr, _, _ := r.ReadRune()
		otherBuffer.WriteRune(rr)
	}
}

func d09() int {
	file, err := os.Open("input/d09.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		rune, _, err := reader.ReadRune()
		if err == io.EOF {
			log.Printf("Reached end of file")
			break
		} else if err != nil {
			log.Fatalf("Rune could not be read %v\n", err)
		}

		if rune == 0x28 {
			processCompressed09(reader)
		} else {
			if rune == 0x0a {
				break
			}
			countd09++
		}
	}
	log.Printf("Decompressed Version 1 %v \n", countd09)
	return countd09
}

// Part 2
/**
  Called for starting of a new Block
*/

func decompress(str string, recursionLevel int) (length int) {
	var index int
	theMultiplierPattern := regexp.MustCompile("\\(([0-9x]*)\\)")

	for index < len(str) {
		// Search for start Character
		if str[index] == '(' {
			//	log.Printf("Recursion Level %v\n", recursionLevel)

			// extract Header info
			theMultiplier := theMultiplierPattern.FindStringSubmatch(str[index:])[1]
			nums := strings.Split(theMultiplier, "x")

			count, _ := strconv.Atoi(nums[0])
			repetitions, _ := strconv.Atoi(nums[1])
			//log.Printf("Multiply %v chars exactly %v times\n", count, repetitions)

			// find substring to process
			recursionStart := index + len(theMultiplier) + 2
			recursionString := str[recursionStart : recursionStart+count]
			//log.Printf("Now recurse substring %v\n", recursionString)
			length += decompress(recursionString, recursionLevel+1) * repetitions
			index += count + len(theMultiplier) + 2
		} else {
			length++
			index++
		}
	}
	countd09Version2 = length

	return length
}

func d09part2() int {
	input, _ := ioutil.ReadFile("input/d09.txt")
	rc := decompress(string(input), 1)
	rc = rc - 1
	log.Printf("d09 count2 = %v\n", rc)
	return rc
}

func d09Main() {
	log.Printf("Day 9\n")
	d09()
	d09part2()
}
