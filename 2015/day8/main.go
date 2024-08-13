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
	file := openFile("./input.txt")

	reader := createReader(file)

	lines := readLines(reader)

	tcl := totalCodeLength(&lines)

	tsl := totalStringLength(&lines)

	tel := totalEncodedLength(&lines)

	result := solution(tcl, tsl)
	result2 := solutionTwo(tel, tcl)

	fmt.Println(result)
	fmt.Println(result2)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
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
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read the line error is: %v", err)
		}

		lines = append(lines, string(line))

	}
	return lines
}

func totalCodeLength(lines *[]string) int {
	total := 0

	for _, line := range *lines {
		total += len(line)
	}

	return total
}

func totalStringLength(lines *[]string) int {
	total := 0

	for _, line := range *lines {

		newString, err := strconv.Unquote(line)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		total += len(newString)
	}

	return total
}

func totalEncodedLength(lines *[]string) int {
	total := 0

	for _, line := range *lines {

		newString := strconv.Quote(line)

		fmt.Println(newString)

		total += len(newString)
	}

	return total
}

func solution(totalCodeLength, totalStringLength int) int {
	return totalCodeLength - totalStringLength
}

func solutionTwo(totalEncodedLength, totalCodeLength int) int {
	return totalEncodedLength - totalCodeLength
}
