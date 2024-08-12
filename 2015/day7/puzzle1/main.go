package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

func findValues(lines *[]string) map[string]uint16 {
	values := map[string]uint16{}
	re := regexp.MustCompile(`(\d+) -> (\w+)`)

	for _, line := range *lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}

		matchInInt, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}

		values[match[2]] = uint16(matchInInt)
	}

	fmt.Println(values)

	return values
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

func operate(line *string, values map[string]uint16) {
	re := regexp.MustCompile(`(\w{1,2}) (RSHIFT|AND|OR|LSHIFT) (\w{1,2}|\d) -> (\w{1,2})`)
	renot := regexp.MustCompile(`(NOT) (\w{1,2}|\d) -> (\w{1,2})`)
	reg := regexp.MustCompile(`(\w{1,2}) -> (\w{1,2})`)

	matchreg := reg.MatchString(*line)
	if matchreg == true {
		match := reg.FindStringSubmatch(*line)
		var v uint16
		r, err := strconv.Atoi(match[1])
		if err != nil {
			v = values[match[1]]
		}

		v = uint16(r)

		values[match[2]] = v
	}

	matchnot := renot.MatchString(*line)
	if matchnot == true {
		match := renot.FindStringSubmatch(*line)
		a := values[match[2]]
		result := notOperation(a)

		values[match[3]] = result
	}

	match := re.MatchString(*line)
	if match == true {
		match := re.FindStringSubmatch(*line)
		switch match[2] {
		case "AND":
			a := values[match[1]]

			var b uint16
			r, err := strconv.Atoi(match[3])
			if err != nil {
				b = values[match[3]]
			}

			b = uint16(r)

			result := andOperation(a, b)

			values[match[4]] = result
		case "OR":
			a := values[match[1]]
			var b uint16
			r, err := strconv.Atoi(match[3])
			if err != nil {
				b = values[match[3]]
			}

			b = uint16(r)

			result := orOperation(a, b)

			values[match[4]] = result
		case "LSHIFT":
			a := values[match[1]]
			var b uint16
			r, err := strconv.Atoi(match[3])
			if err != nil {
				b = values[match[3]]
			}

			b = uint16(r)

			result := leftShiftOperation(a, b)

			values[match[4]] = result

		case "RSHIFT":
			a := values[match[1]]
			var b uint16
			r, err := strconv.Atoi(match[3])
			if err != nil {
				b = values[match[3]]
			}

			b = uint16(r)

			result := rightShiftOperation(a, b)

			values[match[4]] = result
		}
	}

}

func puzzle1(lines []string) uint16 {

	values := findValues(&lines)

	for _, line := range lines {
		operate(&line, values)
	}

	return values["a"]
}
