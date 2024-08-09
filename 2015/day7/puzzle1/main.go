package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func readLines(reader *bufio.Reader) map[string]string {

	circuitMap := map[string]string{}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read the line error is: %v", err)
		}

		circuitMap[getMapKey(line)] = getMapValue(line)

	}
	return circuitMap
}

func getMapKey(text string) string {
	return strings.Split(text, " -> ")[1]
}

func getMapValue(text string) string {
	return strings.Split(text, " -> ")[0]
}

func puzzle1() int {
	file := openFile()

	reader := createReader(file)

	circuitMap := readLines(reader)
	fmt.Println(circuitMap)
	return 0
}
