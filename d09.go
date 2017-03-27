package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	countd09          = 0
	completeStringd09 bytes.Buffer
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
	log.Printf("Repeat the following %v characters %v times\n", length, amount)
	// read amount number of runes
	countd09 = countd09 + (amount * length)

	var otherBuffer bytes.Buffer
	for index := 0; index < length; index++ {
		rr, _, _ := r.ReadRune()
		otherBuffer.WriteRune(rr)
	}
	log.Printf("Other buffer is %v\n", otherBuffer.String())

	for index := 0; index < amount; index++ {
		completeStringd09.Write(otherBuffer.Bytes())
	}
}

func d09() {
	log.Printf("Day 9\n")

	file, err := os.Open("input/d09.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	log.Printf("%v", file.Name())
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
		//log.Printf("%v byte avail\n", string(rune))

		if rune == 0x28 {
			//process processed
			processCompressed09(reader)

		} else {
			//process uncompressed
			//	log.Printf("%v", string(rune))
			if rune == 0x0a {
				//ignore linefeed
				break
			}
			countd09++
			completeStringd09.WriteRune(rune)
		}
	}
	log.Printf("CompleteString is \n%v", completeStringd09.String())
	log.Printf("Total number of characters is %v \n", countd09)
	log.Printf("Completelength is %v \n", completeStringd09.Len())
}
