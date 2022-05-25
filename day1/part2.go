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

    curr_value := 0
    curr_num := 0
    arr := []int{}
    ans := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        value, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal(err)
        }
        if curr_num > 2 {
            new_value := curr_value + value - arr[curr_num - 3]
            if new_value > curr_value {
                ans++
            }
            curr_value = new_value
        } else {
            curr_value += value
        }
        arr = append(arr, value)
        curr_num++
    }
    fmt.Println(ans)

}
