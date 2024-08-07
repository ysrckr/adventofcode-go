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

func main() {

	result1 := puzzle1()

	fmt.Println(result1)

}

func puzzle1() int {
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

		command := &Command{}

		if strings.Split(string(line), " ")[0] == "toggle" {
			command.action = "toggle"
			command.beginingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[1], ",")[0])
			if err != nil {
				continue
			}
			command.beginingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[1], ",")[1])
			if err != nil {
				continue
			}
			command.endingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[3], ",")[0])
			if err != nil {
				continue
			}
			command.endingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[3], ",")[1])
			if err != nil {
				continue
			}
		} else {
			command.action = strings.Split(string(line), " ")[1]
			command.beginingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[2], ",")[0])
			if err != nil {
				continue
			}
			command.beginingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[2], ",")[1])
			if err != nil {
				continue
			}
			command.endingCoords.x, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[4], ",")[0])
			if err != nil {
				continue
			}
			command.endingCoords.y, err = strconv.Atoi(strings.Split(strings.Split(string(line), " ")[4], ",")[1])
			if err != nil {
				continue
			}

		}

		switchLight(command.beginingCoords.x, command.beginingCoords.y, command.endingCoords.x, command.endingCoords.y, &grid, command.action)

	}

	Count(&grid, &count)

	return count
}

func Count(grid *[1000][1000]int, count *int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 1 {
				*count++
			}
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

func switchLight(beginingX, beginingY, endingX, endingY int, grid *[1000][1000]int, action string) {
	for i := beginingX; i <= endingX; i++ {
		for j := beginingY; j <= endingY; j++ {
			if action == "on" {
				grid[i][j] = 1
			}

			if action == "off" {
				grid[i][j] = 0
			}

			if action == "toggle" {

				grid[i][j] ^= 1

			}
		}
	}

}
