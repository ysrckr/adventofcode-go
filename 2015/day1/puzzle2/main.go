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
	position := 0

	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			break
		}

		if count == -1 {
			break
		}

		position++

		switch string(r) {
		case "(":
			count += 1
		case ")":
			count -= 1
		default:
			continue
		}
	}

	fmt.Println(position)

}
