package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ones := make([]int, 12)
	zeros := make([]int, 12)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < 12; i++ {
			if line[i] == '1' {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}
	gamma := ""
	for i := 0; i < 12; i++ {
		if ones[i] > zeros[i] {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}
	gamma_val, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon_val := int(math.Pow(2, 12)) - int(gamma_val) - 1
	fmt.Println(gamma_val * int64(epsilon_val))

}
