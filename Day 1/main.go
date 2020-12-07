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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		values = append(values, val)
	}
	fmt.Println(values)

	for index1, val1 := range values {
		for _, val2 := range values[index1:] {
			if val1+val2 == 2020 {
				fmt.Println(val1 * val2)
			}
		}
	}

}
