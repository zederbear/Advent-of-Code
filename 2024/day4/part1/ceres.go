package main

import (
    "fmt"
    rotate "github.com/hooverz/advent-of-code/day4/rotate"
    "os"
    "bufio"
)

// reverseSlice reverses the elements of a slice and returns the reversed slice
func reverseSlice(slice []rune) []rune {
	if len(slice) == 0 {
        return slice
	}
    return append(reverseSlice(slice[1:]), slice[0])
}

// reverseLists reverses each slice within a list of slices and returns the modified list
func reverseLists(lists [][]rune) [][]rune {
    newLists := make([][]rune, len(lists))
	for i := range lists {
		newLists[i] = reverseSlice(lists[i])
	}
	return newLists
}

func findSubsequence(slice []rune, subseq []rune) int {
    count := 0
    subLen := len(subseq)
    for i := 0; i <= len(slice)-subLen; i++ {
        if matchSlice(slice[i:i+subLen], subseq) {
            count++
        }
    }
    return count
}

func matchSlice(slice []rune, subseq []rune) bool {
    for i := 0; i < len(subseq); i++ {
        if slice[i] != subseq[i] {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("realdata.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    var listOfRuneLists [][]rune
    for _, line := range lines {
        listOfRuneLists = append(listOfRuneLists, []rune(line))
    }

    for _, line := range listOfRuneLists {
        fmt.Println(line)
    }

    
    fmt.Println("Original Grid:")
    for _, line := range listOfRuneLists {
        fmt.Println(line)
    }

    fmt.Println("Rotated 45 Degrees:")
    rotated45 := rotate.Rotate(listOfRuneLists, 45)
    for _, line := range rotated45 {
        fmt.Println(line)
    }

    fmt.Println("Rotated 90 Degrees:")
    rotated90 := rotate.Rotate(listOfRuneLists, 90)
    for _, line := range rotated90 {
        fmt.Println(line)
    }
    reversed := reverseLists(listOfRuneLists)
    reversedRotated45 := reverseLists(rotated45)
    reversedRotated90 := reverseLists(rotated90)
    rotated135 := rotate.Rotate(rotate.Rotate(listOfRuneLists, 90), 45)
    reversedRotated135 := reverseLists(rotated135)

    totalHorizontal := 0
    for _, line := range listOfRuneLists {
        totalHorizontal += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    totalDiagonal45 := 0
    for _, line := range rotated45 {
        totalDiagonal45 += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    totalVertical := 0
    for _, line := range rotated90 {
        totalVertical += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    totalHorizontalReversed := 0
    for _, line := range reversed {
        totalHorizontalReversed += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    totalDiagonal45Reversed := 0
    for _, line := range reversedRotated45 {
        totalDiagonal45Reversed += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    totalVerticalReversed := 0
    for _, line := range reversedRotated90 {
        totalVerticalReversed += findSubsequence(line, []rune{88, 77, 65, 83})
    }
    totalDiagonal135 := 0
    for _, line := range rotated135 {
        totalDiagonal135 += findSubsequence(line, []rune{88, 77, 65, 83})
    }
    totalDiagonal135Reversed := 0
    for _, line := range reversedRotated135 {
        totalDiagonal135Reversed += findSubsequence(line, []rune{88, 77, 65, 83})
    }

    finalTotal := totalHorizontal + totalDiagonal45 + totalVertical +
        totalHorizontalReversed + totalDiagonal45Reversed + totalVerticalReversed +
        totalDiagonal135 + totalDiagonal135Reversed
    fmt.Println("Total Horizontal:", totalHorizontal)
    fmt.Println("Total Diagonal 45:", totalDiagonal45)
    fmt.Println("Total Vertical:", totalVertical)
    fmt.Println("Total Horizontal Reversed:", totalHorizontalReversed)
    fmt.Println("Total Diagonal 45 Reversed:", totalDiagonal45Reversed)
    fmt.Println("Total Vertical Reversed:", totalVerticalReversed)
    fmt.Println("Total Diagonal 135:", totalDiagonal135)
    fmt.Println("Total Diagonal 135 Reversed:", totalDiagonal135Reversed)
    fmt.Println("Final Total:", finalTotal)

}
