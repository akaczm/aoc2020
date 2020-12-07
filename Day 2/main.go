package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	passwords := make([]string, 0)
	validCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		passwords = append(passwords, val)
	}

	for _, term := range passwords {
		split := strings.Split(term, ":")
		policy := split[0]
		password := strings.Trim(split[1], " ")
		letter, min, max := decodePolicy(policy)
		if validatePassword(password, min, max, letter) {
			validCount++
		}
	}
	fmt.Println(validCount)

}

// decodePolicy returns the letter, and range of occurences
func decodePolicy(policy string) (string, int, int) {
	split := strings.Split(policy, " ")
	letter := split[1]

	minmax := strings.Trim(split[0], " ")
	minmaxsplit := strings.Split(minmax, "-")
	min, err := strconv.Atoi(minmaxsplit[0])
	check(err)
	max, err := strconv.Atoi(minmaxsplit[1])
	check(err)

	return letter, min, max
}

func validatePassword(password string, min int, max int, letter string) bool {
	count := strings.Count(password, letter)
	if count >= min && count <= max {
		return true
	}
	return false
}
