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
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(file)

	niceStringCount := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		if strings.Contains(string(line), "ab") || strings.Contains(string(line), "cd") || strings.Contains(string(line), "pq") || strings.Contains(string(line), "xy") {
			continue
		}

		vowelCount := 0
		hasDouble := false

		for i, c := range line {
			if i > 0 {
				if line[i-1] == c {
					hasDouble = line[i-1] == c
				}
			}
			switch string(c) {
			case "a", "e", "i", "o", "u":
				vowelCount++

			default:
				continue
			}
		}

		if vowelCount >= 3 && hasDouble {
			niceStringCount++
		} else {
			continue
		}

	}

	fmt.Println(niceStringCount)
}
