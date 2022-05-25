package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    curr_value := 1000000
    ans := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        value, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal(err)
        }
        if value > curr_value {
            ans += 1
        }
        curr_value = value
    }
    fmt.Println(ans)

}
