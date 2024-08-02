package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	count := 0

	reader := bufio.NewReader(file)

	for {
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		} else {
			switch string(r) {
			case "(":
				count += 1
			case ")":
				count -= 1
			default:
				continue
			}
		}

	}

	fmt.Println(count)

}
