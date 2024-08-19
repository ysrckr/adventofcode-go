package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	createRoutes()
}

type Route struct {
	src      string
	dest     string
	distance int
}

var allRoutes []Route

func createRoutes() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			break
		}

		splitByTo := strings.Join(strings.Split(string(line), " to "), " ")
		structToBe := strings.Split(splitByTo, " ")
		src, dest, distanceAsStr := structToBe[0], structToBe[1], structToBe[3]

		distance, err := strconv.Atoi(distanceAsStr)
		if err != nil {
			log.Fatalln(err)
			break
		}

		route := Route{
			src:      src,
			dest:     dest,
			distance: distance,
		}

		allRoutes = append(allRoutes, route)
	}

	log.Println(allRoutes)
}
