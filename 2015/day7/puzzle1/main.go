package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	result := puzzle1()

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

func findValues(lines *[]string) map[string]int {
	values := map[string]int{}
	re := regexp.MustCompile(`(\d+) -> (\w+)`)

	for _, line := range *lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}

		_, ok := values[match[2]]
		if ok {
			continue
		}

		matchInInt, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}

		values[match[2]] = matchInInt
	}

	return values
}

func filterCircuit(lines *[]string) {
	re := regexp.MustCompile(`(\d+) -> (\w+)`)
	new := []string{}

	for i, line := range *lines {
		match := re.MatchString(line)
		if match == false {
			continue
		}

		new = slices.Delete(*lines, i, i+1)
		*lines = new

	}

}

func notOperation(a int) int {
	return ^a
}

func orOperation(a, b int) int {
	return a | b
}

func andOperation(a, b int) int {
	return a & b
}

func leftShiftOperation(a, b int) int {
	return a << b
}

func rightShiftOperation(a, b int) int {
	return a >> b
}

func operate(line *string, values map[string]int) {
	re := regexp.MustCompile(`(\w{1,2}) (LSHIFT|RSHIFT|AND|OR) (\w{1,2}|\d) -> (\w{1,2})`)
	renot := regexp.MustCompile(`(NOT) (\w{1,2}|\d) -> (\w{1,2})`)
	reg := regexp.MustCompile(`(\w{1,2}) -> (\w{1,2})`)

	matchSwap := reg.MatchString(*line)
	if matchSwap == true {
		match := reg.FindStringSubmatch(*line)

		values[match[2]] = values[match[1]]
	}

	match := re.MatchString(*line)
	if match == true {
		match := re.FindStringSubmatch(*line)
		switch match[2] {
		case "AND":
			a := values[match[1]]
			b := values[match[3]]

			result := andOperation(a, b)

			values[match[4]] = result
		case "OR":
			a := values[match[1]]
			b := values[match[3]]

			result := orOperation(a, b)

			values[match[4]] = result
		case "LSHIFT":
			a := values[match[1]]
			b := values[match[3]]

			result := leftShiftOperation(a, b)

			values[match[4]] = result

		case "RSHIFT":
			a := values[match[1]]
			b := values[match[3]]

			result := rightShiftOperation(a, b)

			values[match[4]] = result
		}
	}
	matchnot := renot.MatchString(*line)
	if matchnot == true {
		match := renot.FindStringSubmatch(*line)
		a := values[match[2]]
		result := notOperation(a)
		values[match[3]] = result
	}

}

func puzzle1() int {
	file := openFile()

	reader := createReader(file)

	lines := readLines(reader)
	values := findValues(&lines)
	filterCircuit(&lines)

	fmt.Println("before: ", values)

	for _, line := range lines {
		operate(&line, values)
	}

	fmt.Println("after: ", values)

	return values["a"]
}
