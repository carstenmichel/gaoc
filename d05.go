package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
)

func d05() string {
	log.Printf("Day 05\n")
	puzzleInput := []byte("wtnhxymk")
	var buffer bytes.Buffer
	var sum [16]byte
	var done = 0
	for index := 0; true; index++ {
		bb := append(puzzleInput, []byte(strconv.Itoa(index))...)
		sum = md5.Sum(bb)
		if (sum[0] == 0x00) && (sum[1] == 0x0) && (sum[2]|0x0f == 0x0f) {
			buffer.WriteString(fmt.Sprintf("%x", sum[2]))
			done = done + 1
			if done == 8 {
				break
			}
		}
	}
	s := buffer.String()
	log.Printf("password is %v\n", s)
	return s
}

//D0502BitField contains the bitfield
type D0502BitField struct {
	Byte0 bool
	Byte1 bool
	Byte2 bool
	Byte3 bool
	Byte4 bool
	Byte5 bool
	Byte6 bool
	Byte7 bool
}

func (bf D0502BitField) allSet() bool {
	if bf.Byte0 && bf.Byte1 && bf.Byte2 && bf.Byte3 && bf.Byte4 && bf.Byte5 && bf.Byte6 && bf.Byte7 {
		return true
	}
	return false
}

func (bf D0502BitField) setField(index byte) D0502BitField {
	switch index {
	case 0:
		bf.Byte0 = true
	case 1:
		bf.Byte1 = true
	case 2:
		bf.Byte2 = true
	case 3:
		bf.Byte3 = true
	case 4:
		bf.Byte4 = true
	case 5:
		bf.Byte5 = true
	case 6:
		bf.Byte6 = true
	case 7:
		bf.Byte7 = true
	}
	return bf
}

func (bf D0502BitField) getField(index byte) bool {
	switch index {
	case 0:
		return bf.Byte0
	case 1:
		return bf.Byte1
	case 2:
		return bf.Byte2
	case 3:
		return bf.Byte3
	case 4:
		return bf.Byte4
	case 5:
		return bf.Byte5
	case 6:
		return bf.Byte6
	case 7:
		return bf.Byte7
	}
	return false
}

func d05Part2() string {
	log.Printf("Day 05 Part 2\n")
	puzzleInput := []byte("wtnhxymk")
	resultPassword := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	var sum [16]byte

	flagsField := D0502BitField{false, false, false, false, false, false, false, false}
	for index := 0; true; index++ {
		bb := append(puzzleInput, []byte(strconv.Itoa(index))...)
		sum = md5.Sum(bb)
		if (sum[0] == 0x00) && (sum[1] == 0x0) && (sum[2]|0x0f == 0x0f) {
			position := sum[2] & 0x0f
			if position < 8 {
				if !flagsField.getField(position) {
					valum := (sum[3] & 0xf0) / 16
					resultPassword[position] = valum
					flagsField = flagsField.setField(position)
					if flagsField.allSet() {
						break
					}
				}
			}
		}
	}
	s := fmt.Sprintf("%x", resultPassword)
	var rs bytes.Buffer
	for index, cha := range s {
		if index%2 == 1 {
			rs.WriteRune(cha)
		}
	}
	log.Printf("password is %v\n", rs.String())
	return rs.String()
}
