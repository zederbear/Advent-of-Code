package main

import (
    "fmt"
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

func checkSafe(report []int) bool {
    var lastdif int
    for i := 1; i < len(report); i++ {
        // 5 5 4 3 2 1
        dif := report[i] - report[i-1]
        
        if dif == 0 {
            return false
        }
        if dif > 3 || dif < -3 {
            return false
        }
        if i == 1 {
            lastdif = dif
            continue
        }

        
        if dif < 0 && lastdif > 0 {
            return false
        }
        if dif > 0 && lastdif < 0 {
            return false
        }
        lastdif = dif
    }
    return true
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

func main() {
    file, err := os.Open("realdata.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    var reports [][]int
    for scanner.Scan() {
        line := scanner.Text()
        var report []int
        for _, value := range strings.Split(line, " ") {
            num, err := strconv.Atoi(value)
            if err != nil {
                log.Fatal(err)
            }
            report = append(report, num)
        }
        reports = append(reports, report)
    }
    var safeornot []bool
    for _, report := range reports {
        safeornot = append(safeornot, checkSafe(report))
    }

    for i, safety := range safeornot {
        fmt.Print(i)
        fmt.Println(" | ", safety)
        if safety {
            fmt.Print(reports[i])
            fmt.Println(" | SAFE")
        } else {
            fmt.Print(reports[i])
            fmt.Println(" | UNSAFE")
        }
    }

    var total int
    for _, safe := range safeornot {
        total += boolToInt(safe)
    }
    fmt.Println(total)
    
}
