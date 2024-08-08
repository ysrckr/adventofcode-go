package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func readLines(reader *bufio.Reader) {

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read the line error is: %v", err)
		}

		fmt.Println(line)

	}

}

func puzzle1() int {
	file := openFile()

	reader := createReader(file)

	readLines(reader)
	return 0
}
