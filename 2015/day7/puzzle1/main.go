package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var wires map[string]uint64

func main() {
	file := openFile()

	reader := createReader(file)

	lines := readLines(reader)



	result := puzzle1(&lines)

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

func isInt(possibleInt string) bool {
	if _, err := strconv.Atoi(possibleInt); err != nil {
		return false
	}

	return true
}

func toInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

func dfs(wires map[string]string, entry string, memo map[string]int) int {
	if memoVal, ok := memo[entry]; ok {
		return memoVal
	}


	if regexp.MustCompile("[0-9]").MatchString(entry) {
		return toInt(entry)
	}

	sourceRule := wires[entry]
	parts := strings.Split(sourceRule, " ")

	var result int
	switch {
	case len(parts) == 1:
		result = dfs(wires, parts[0], memo)
	case parts[0] == "NOT":
		start := dfs(wires, parts[1], memo)
		result = (math.MaxUint16) ^ start
	case parts[1] == "AND":
		result = dfs(wires, parts[0], memo) & dfs(wires, parts[2], memo)
	case parts[1] == "OR":
		result = dfs(wires, parts[0], memo) | dfs(wires, parts[2], memo)
	case parts[1] == "LSHIFT":
		result = dfs(wires, parts[0], memo) << dfs(wires, parts[2], memo)
	case parts[1] == "RSHIFT":
		result = dfs(wires, parts[0], memo) >> dfs(wires, parts[2], memo)
	}

  fmt.Println(result)

	memo[entry] = result
	return result
}

func puzzle1(lines *[]string) int {
	wires := map[string]string{}

	for _, instruction := range *lines {
		parts := strings.Split(instruction, " -> ")
		wires[parts[1]] = parts[0]
	}

	aSignal := dfs(wires, "a", map[string]int{})

	return aSignal

}
