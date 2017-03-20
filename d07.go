package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	tls int
	ssl int
)

func hasAnna(anna string) bool {
	for index := 0; index < (len(anna) - 3); index++ {
		if anna[index] == anna[index+3] && anna[index+1] == anna[index+2] && anna[index] != anna[index+1] {
			return true
		}
	}
	return false
}
func checkd07SSL(inner []string, outer []string) {
	thisLineHasSSL := false
	for index := 0; index < len(outer); index++ {
		outerString := outer[index]
		for outCounter := 0; outCounter < (len(outerString) - 2); outCounter++ {
			if outerString[outCounter] == outerString[outCounter+2] && outerString[outCounter] != outerString[outCounter+1] {
				aa := outerString[outCounter]
				bb := outerString[outCounter+1]
				for i2 := 0; i2 < len(inner); i2++ {
					innerString := inner[i2]
					for inCounter := 0; inCounter < (len(innerString) - 2); inCounter++ {
						if innerString[inCounter] == bb && innerString[inCounter+1] == aa && innerString[inCounter+2] == bb {
							thisLineHasSSL = true
						}
					}
				}
			}
		}
	}
	if thisLineHasSSL {
		ssl = ssl + 1
	}
}

func processD07Line(in string) {
	work := in
	outer := []string{}
	inner := []string{}
	laenge := len(work)
	isOutside := true
	for index := 0; index < laenge; index++ {
		delimiter := 0
		if isOutside {
			delimiter = strings.Index(work, "[")
			if delimiter == -1 {
				outer = append(outer, work)
				break
			}
			outer = append(outer, work[:delimiter])
			isOutside = false
		} else {
			delimiter = strings.Index(work, "]")
			inner = append(inner, work[:delimiter])
			isOutside = true
		}
		work = work[delimiter+1:]
		index = delimiter + 1

	}
	hasInnerAbba := false
	hasOuterAbba := false
	for index := 0; index < len(inner); index++ {
		if hasAnna(inner[index]) {
			hasInnerAbba = true
		}
	}
	if hasInnerAbba == false {
		for index := 0; index < len(outer); index++ {
			if hasAnna(outer[index]) {
				hasOuterAbba = true
			}
		}
	}
	if hasOuterAbba {
		tls = tls + 1
	}
	checkd07SSL(inner, outer)
}

func d07() (int, int) {
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
	log.Printf("TLS is %v\n", tls)
	log.Printf("SSL is %v\n", ssl)

	return tls, ssl
}
