package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	values := make([]int, 0)

	//builds an array of integers based off file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		values = append(values, val)
	}

	//iterates over it and returns the result of multiplying the two values
	//that sum to 2020
	for index1, val1 := range values {
		for _, val2 := range values[index1:] {
			if val1+val2 == 2020 {
				fmt.Println(val1 * val2)
			}
		}
	}

}
