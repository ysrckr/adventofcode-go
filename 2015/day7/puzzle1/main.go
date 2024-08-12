package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	command     string
	arguments   []string
	destination string
}

func newInstruction(line *string) Instruction {
	instruction := Instruction{}

	commandRe := regexp.MustCompile(`[A-Z]+`)
	argumentsRe := regexp.MustCompile(`[a-z0-9]+`)

	instruction.command = commandRe.FindString(*line)
	instruction.arguments = argumentsRe.FindAllString(*line, -1)
	instruction.destination = instruction.arguments[len(instruction.arguments)-1]
	instruction.arguments = instruction.arguments[:len(instruction.arguments)-1]

	return instruction
}

type Operation interface {
	not(uint16) uint16
	or(uint16, uint16) uint16
	and(uint16, uint16) uint16
	leftShift(uint16, uint16) uint16
	rightShift(uint16, uint16) uint16
}

func not(a uint16) uint16 {
	return ^a
}

func or(a, b uint16) uint16 {
	return a | b
}

func and(a, b uint16) uint16 {
	return a & b
}

func leftShift(a, b uint16) uint16 {
	return a << b
}

func rightShift(a, b uint16) uint16 {
	return a >> b
}

var wires map[string]uint64

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

	for _, line := range lines {
		instruction := newInstruction(&line)

		wires[instruction.destination] = 0
	}

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

func isInt(possibleInt string) bool {
	if _, err := strconv.Atoi(possibleInt); err != nil {
		return false
	}

	return true
}
