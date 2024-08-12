package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file := openFile()

	reader := createReader(file)

	lines := readLines(reader)

	// lines := []string{
	// 	"123 -> x",
	// 	"456 -> y",
	// 	"x AND y -> d",
	// 	"x OR y -> e",
	// 	"x LSHIFT 2 -> f",
	// 	"y RSHIFT 2 -> g",
	// 	"NOT x -> h",
	// 	"NOT y -> i",
	// }

	result := puzzle1(lines)

	fmt.Println(result)
}

func openFile() *os.File {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("could not open file error is: %v", err)
	}

	return file
}

func createReader(file *os.File) *bufio.Reader {
	reader := bufio.NewReader(file)

	return reader
}

func readLines(reader *bufio.Reader) []string {

	lines := []string{}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read the line error is: %v", err)
		}

		lines = append(lines, line)

	}
	return lines
}

func puzzle1(lines []string) uint16 {
	var wires map[string]string

	return 0
}

func notOperation(a uint16) uint16 {
	return ^a
}

func orOperation(a, b uint16) uint16 {
	return a | b
}

func andOperation(a, b uint16) uint16 {
	return a & b
}

func leftShiftOperation(a, b uint16) uint16 {
	return a << b
}

func rightShiftOperation(a, b uint16) uint16 {
	return a >> b
}

func isInt(possibleInt string) bool {
	if _, err := strconv.Atoi(possibleInt); err != nil {
		return false
	}

	return true
}
