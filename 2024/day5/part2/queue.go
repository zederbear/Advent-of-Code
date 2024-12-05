package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strSliceToIntSlice(slice []string) []int {
	var intSlice []int
	for _, value := range slice {
		intValue, _ := strconv.Atoi(value)
		intSlice = append(intSlice, intValue)
	}
	return intSlice
}

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func correctUpdate(update []int, rules [][]int) []int {
	for {
		changed := false
		for _, rule := range rules {
			index0 := indexOf(update, rule[0])
			index1 := indexOf(update, rule[1])

			if index0 == -1 || index1 == -1 {
				continue
			}

			if index0 > index1 {
				update[index0], update[index1] = update[index1], update[index0]
				changed = true
			}
		}

		if !changed {
			break
		}
	}
	return update
}

func main() {
	file, err := os.Open("realdata.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var updates = [][]int{}
	var rules = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strSliceToIntSlice(strings.Split(line, "|"))
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			update := strSliceToIntSlice(strings.Split(line, ","))
			updates = append(updates, update)
		}
	}

	fmt.Println("Rules:", rules)
	fmt.Println("Updates:", updates)

	var reorderedUpdates [][]int
	for i, update := range updates {
		originalUpdate := append([]int{}, update...)
		correctedUpdate := correctUpdate(update, rules)

		if slicesEqual(originalUpdate, correctedUpdate) {
			fmt.Printf("Update %d is already valid: %v\n", i+1, correctedUpdate)
		} else {
			fmt.Printf("Reordered Update %d: %v\n", i+1, correctedUpdate)
			reorderedUpdates = append(reorderedUpdates, correctedUpdate)
		}
	}

	total := 0
	for _, update := range reorderedUpdates {
		total += update[len(update)/2]
	}
	fmt.Println("Total (Reordered Middle Sums):", total)
}

