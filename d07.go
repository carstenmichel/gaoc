package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	tls int
)

func hasAnna(anna string) bool {
	for index := 0; index < (len(anna) - 3); index++ {
		if anna[index] == anna[index+3] && anna[index+1] == anna[index+2] && anna[index] != anna[index+1] {
			return true
		}
	}
	return false
}

func processD07Line(in string) {
	log.Printf("------------------------")
	log.Printf("Inspect: %v\n", in)
	work := in
	outer := []string{}
	inner := []string{}
	laenge := len(work)
	isOutside := true
	for index := 0; index < laenge; index++ {
		delimiter := 0
		if isOutside {
			delimiter = strings.Index(work, "[")
			log.Printf("Delimiter for outside is %v\n", delimiter)
			if delimiter == -1 {
				outer = append(outer, work)
				break
			}
			outer = append(outer, work[:delimiter])
			isOutside = false
		} else {
			delimiter = strings.Index(work, "]")
			log.Printf("Delimiter for inside is %v\n", delimiter)
			inner = append(inner, work[:delimiter])
			isOutside = true
		}
		work = work[delimiter+1:]
		log.Printf("remains: %v\n", work)
		index = delimiter + 1

	}
	log.Printf("Outside %v\n", outer)
	log.Printf("Inner %v\n", inner)
	// if inner has annagram then fail
	for index := 0; index < len(inner); index++ {
		if hasAnna(inner[index]) {
			return
		}
	}
	for index := 0; index < len(outer); index++ {
		if hasAnna(outer[index]) {
			tls = tls + 1
			return
		}
	}

}

func d07() {
	log.Printf("Day 7\n")

	file, err := os.Open("input/d07.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		textLine := scanner.Text()
		processD07Line(textLine)

	}
	log.Printf("TLS  %v\n", tls)
	//return sumOfSectors, secretSector
}
