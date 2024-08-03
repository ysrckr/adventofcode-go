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

type Box struct {
	l uint64
	w uint64
	h uint64
}

func newBox(str string) *Box {
	splitString := strings.Split(str, "x")
	length, err := strconv.Atoi(splitString[0])
	if err != nil {
		return nil
	}

	width, err := strconv.Atoi(splitString[1])
	if err != nil {
		return nil
	}

	height, err := strconv.Atoi(splitString[2])
	if err != nil {
		return nil
	}

	return &Box{
		l: uint64(length),
		w: uint64(width),
		h: uint64(height),
	}
}

func (b *Box) totalSurfaceArea() uint64 {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b *Box) smallestSurfaceArea() uint64 {
	return min(b.l*b.w, b.w*b.h, b.h*b.l)
}

func (b *Box) totalPaperNeeded() uint64 {
	return b.totalSurfaceArea() + b.smallestSurfaceArea()
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

	var total uint64 = 0

	for {
		r, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			break
		}

		box := newBox(string(r))
		total += box.totalPaperNeeded()
	}

	fmt.Println(total)
}
