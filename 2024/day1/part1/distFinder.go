package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	file, err := os.Open("testdata.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var answer []int

	for i := 0; i < len(list1); i++ {
		answer = append(answer, int(math.Abs(float64(list1[i]-list2[i]))))
	}

	fmt.Println(sum(answer))
}
