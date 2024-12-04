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

	pattern := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(text, -1)

	fmt.Println("Matches found: ")

	for _, match := range matches {
		fmt.Println(match)
	}

	var result [][]int

	for _, match := range matches {
		if len(match) == 3 {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				log.Fatal(err1, err2)
			}

			result = append(result, []int{num1, num2})
		}
	}

	fmt.Println("Extracted numbers:")
	for _, pair := range result {
		fmt.Println(pair)
	}

	var total int
	for _, pair := range result {
		total += pair[0] * pair[1]
	}

	fmt.Println("Total: ")
	fmt.Println(total)

}
