package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	counter := 0

	for {

		hash := getMD5Hash("bgvyzdsv" + strconv.Itoa(counter))

		if isStartsWithSixZeros(hash) {
			break
		}

		counter++
	}

	fmt.Println(counter)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func isStartsWithSixZeros(text string) bool {
	return strings.HasPrefix(text, "000000")
}
