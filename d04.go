package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	sumOfSectors = 0
	secretSector = 0
	re           = regexp.MustCompile("(.*)-([0-9]+)\\[([a-z]*)\\]")
	northPole    = regexp.MustCompile(".*north.*")
)

// KeyValuePair hold a Character and its number of occurence
type KeyValuePair struct {
	Character string
	Occurence int
}

// AllCharacters holds a list of characters together
// with info about usage
type AllCharacters []KeyValuePair

// Len returns the length of the list
func (ac AllCharacters) Len() int {
	return len(ac)
}

// Less function is implemented here
func (ac AllCharacters) Less(i, j int) bool {
	if ac[i].Occurence == ac[j].Occurence {
		return ac[i].Character > ac[j].Character
	}
	return ac[i].Occurence < ac[j].Occurence
}

func (ac AllCharacters) Swap(i, j int) {
	ac[i], ac[j] = ac[j], ac[i]
}

func decipherd04(riddle string, shiftLevel int) {
	sl := shiftLevel % 26
	var buffer bytes.Buffer
	for _, c := range riddle {
		var ci = int(c)
		if c == 0x2d {
			ci = 0x20
		} else {
			ci = ci + sl
			if ci > 0x7a {
				ci = ci - 26
			}
		}
		buffer.WriteByte(byte(ci))
	}
	deciphered := buffer.String()
	if len(northPole.FindStringIndex(deciphered)) > 0 {
		secretSector = shiftLevel
	}

}

func processd04(line string) {
	elements := re.FindStringSubmatch(line)

	codeonly := strings.Replace(elements[1], "-", "", -1)
	singleCharacters := strings.Split(codeonly, "")
	//var m map[string]int
	m := make(map[string]int)
	for _, element := range singleCharacters {
		s := element
		m[s] = m[s] + 1
	}

	var allChars = make(AllCharacters, len(m))

	for k, v := range m {
		allChars = append(allChars, KeyValuePair{Character: k, Occurence: v})
	}
	sort.Sort(sort.Reverse(allChars))

	checksum := elements[3]
	var checkSumIsMatching = true
	for index := 0; index < 5; index++ {
		if checksum[index] != allChars[index].Character[0] {
			checkSumIsMatching = false
		}
	}
	if checkSumIsMatching {
		sectorid, err := strconv.Atoi(elements[2])
		if err != nil {
			log.Fatalf("Sectorid %v could not be parsed\n", elements[2])
		}
		decipherd04(elements[1], sectorid)
		sumOfSectors = sumOfSectors + sectorid
	}
}

func d04() (int, int) {
	log.Printf("Day 4\n")

	file, err := os.Open("input/d04.txt")
	if err != nil {
		log.Fatalf("Failed to read input file %v \n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		textLine := scanner.Text()
		processd04(textLine)

	}
	log.Printf("Number of Sectors  %v\n", sumOfSectors)
	log.Printf("Secret Sector  %v\n", secretSector)
	return sumOfSectors, secretSector
}
