package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func (b *Box) getVolume() uint64 {
	return b.l * b.w * b.h
}

func (b *Box) getSmallestTwoDimension() []uint64 {
	dimensions := []int{
		int(b.h), int(b.l), int(b.w),
	}

	sort.Ints(dimensions)

	return []uint64{
		uint64(dimensions[0]),
		uint64(dimensions[1]),
	}
}

func (b *Box) getBowLength() uint64 {
	dimensions := b.getSmallestTwoDimension()

	return 2*dimensions[0] + 2*dimensions[1]
}

func (b *Box) getRibbonLength() uint64 {
	return b.getVolume() + b.getBowLength()
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
		total += box.getRibbonLength()
	}

	fmt.Println(total)
}
