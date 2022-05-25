package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    horizontal_position := 0
    depth := 0
    aim := 0

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        splitted_line := strings.Split(line, " ")
        command, value := splitted_line[0], splitted_line[1]
        int_value, err := strconv.Atoi(value)
        if err != nil {
            log.Fatal(err)
        }
        if command == "forward" {
            horizontal_position += int_value
            depth += aim*int_value
        } else if command == "down" {
            aim += int_value
        } else if command == "up" {
            aim -= int_value
        }
    }
    fmt.Println(horizontal_position*depth)

}
