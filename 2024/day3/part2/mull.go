package main

import (
	"fmt"
	// "bufio"
	"io"
	"os"

	// "strings"
	"log"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("testdata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)

	mulPattern := `mul\((\d+),(\d+)\)`
	dontPattern := `don't\(\)`
	doPattern := `do\(\)`

	pattern := fmt.Sprintf("%s|%s|%s", mulPattern, dontPattern, doPattern)

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(text, -1)
	fmt.Println(matches)

	do := true

	var trueMatches [][]string

	for _, val := range matches {
		if val[0] == "do()" {
			do = true
		} else if val[0] == "don't()" {
			do = false
		} else if do {
			trueMatches = append(trueMatches, val)
		}
	}

	fmt.Println(trueMatches)

	var total int
	for _, pair := range trueMatches {
		num1, err1 := strconv.Atoi(pair[1])
		num2, err2 := strconv.Atoi(pair[2])

		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}

		total += num1 * num2
	}

	fmt.Println("Total: ")
	fmt.Println(total)
}
