package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Command struct {
	action         string
	beginingCoords Coords
	endingCoords   Coords
}

func (c *Command) switchLight(grid *[1000][1000]int) {
	for i := c.beginingCoords.x; i <= c.endingCoords.x; i++ {
		for j := c.beginingCoords.y; j <= c.endingCoords.y; j++ {
			if c.action == "on" {
				grid[i][j] += 1
			}

			if c.action == "off" {
				if grid[i][j] > 0 {
					grid[i][j] -= 1
				}
			}

			if c.action == "toggle" {

				grid[i][j] += 2

			}
		}
	}

}

func main() {

	result2 := puzzle2()

	fmt.Println(result2)

}

func puzzle2() int {
	reader := openFile()

	count := 0

	grid := createGrid(0)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		command, err := newCommand(string(line))
		if err != nil {
			continue
		}

		command.switchLight(&grid)

	}

	measureBrightness(&grid, &count)

	return count
}

func newCommand(line string) (*Command, error) {
	command := &Command{}
	var err error
	if strings.Split(line, " ")[0] == "toggle" {
		command.action = "toggle"
		command.beginingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[1], ",")[0])
		if err != nil {
			return nil, err
		}
		command.beginingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[1], ",")[1])
		if err != nil {
			return nil, err
		}
		command.endingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[3], ",")[0])
		if err != nil {
			return nil, err
		}
		command.endingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[3], ",")[1])
		if err != nil {
			return nil, err
		}
	}

	if strings.Split(line, " ")[0] != "toggle" {
		command.action = strings.Split(line, " ")[1]
		command.beginingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[2], ",")[0])
		if err != nil {
			return nil, err
		}
		command.beginingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[2], ",")[1])
		if err != nil {
			return nil, err
		}
		command.endingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[4], ",")[0])
		if err != nil {
			return nil, err
		}
		command.endingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(line, " ")[4], ",")[1])
		if err != nil {
			return nil, err
		}

	}

	return command, nil
}

func measureBrightness(grid *[1000][1000]int, count *int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			*count += grid[i][j]
		}
	}
}

func createGrid(fill int) [1000][1000]int {
	grid := [1000][1000]int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = fill
		}
	}

	return grid
}

func openFile() *bufio.Reader {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(file)
	return reader
}
