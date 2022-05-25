package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func find_req_bit(a []string, bit_num int, oxygen int) int {
    zero_count, one_count := 0, 0
    for _, val := range a {
        if val[bit_num] == '1' {
            one_count++
        } else {
            zero_count++
        }
    }
    ans := 0
    if oxygen == 1 {
        if zero_count > one_count {
            ans = 0
        } else {
            ans = 1
        }
    } else {
        if zero_count > one_count {
            ans = 1
        } else {
            ans = 0
        }
    }
    return ans
}

func find_rating(a []string, oxygen int) string {
    ans := a
    curr_bit := 0
    for true {
        new_ans := []string{}
        if len(ans) == 1 || curr_bit >= 12 {
            break
        }
        req_bit := find_req_bit(ans, curr_bit, oxygen)
        for _, val := range ans {
            curr_val := val[curr_bit] - 48
            if curr_val == byte(req_bit) {
                new_ans = append(new_ans, val)
            }
        }
        curr_bit++
        ans = new_ans
    }
    return ans[0]
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    a := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        a = append(a, scanner.Text())
    }
    // fmt.Println(len(a))
    oxygen_rating := find_rating(a, 1)
    co2_rating := find_rating(a, 0)
    oxygen_rating_val, err := strconv.ParseInt(oxygen_rating, 2, 64)
    co2_rating_val, err := strconv.ParseInt(co2_rating, 2, 64)
    fmt.Println(oxygen_rating_val*co2_rating_val)

}
