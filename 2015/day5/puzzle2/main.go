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

	result := puzzle2()

	fmt.Println(result)

}

func puzzle2() int {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(file)

	count := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		isNice := IsNice(string(line))

		if isNice {
			count++
		}

	}

	return count
}

func IsNice(s string) bool {
	return hasPair(s) && hasRepeatedWithOneBetween(s)
}

func hasPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Count(s, pair) > 1 {
			return true
		}
	}
	return false
}

func hasRepeatedWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
