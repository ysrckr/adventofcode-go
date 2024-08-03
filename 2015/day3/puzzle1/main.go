package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

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

	reader := bufio.NewReader(file)

	coordinates := &Coordinates{}

	history := []string{
		strconv.Itoa(coordinates.x) + ":" + strconv.Itoa(coordinates.y),
	}

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		switch r {
		case '^':
			coordinates.y += 1
			if slices.Contains(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y)) {
				break
			}
			history = append(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y))

		case 'v':
			coordinates.y -= 1
			if slices.Contains(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y)) {
				break
			}
			history = append(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y))

		case '>':
			coordinates.x += 1
			if slices.Contains(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y)) {
				break
			}
			history = append(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y))

		case '<':
			coordinates.x -= 1
			if slices.Contains(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y)) {
				break
			}
			history = append(history, strconv.Itoa(coordinates.x)+":"+strconv.Itoa(coordinates.y))
		default:
			break

		}

	}

	fmt.Println(coordinates)

	fmt.Println(len(history))
}
