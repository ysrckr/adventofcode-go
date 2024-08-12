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

func puzzle1() int {
	file := openFile()

	reader := createReader(file)

	lines := readLines(reader)
	values := findValues(&lines)
	fmt.Println(values)
	return 0
}
