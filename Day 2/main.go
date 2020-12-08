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
	var validCount int
	var validCountNew int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		passwords = append(passwords, val)
	}

	for _, term := range passwords {
		split := strings.Split(term, ":")
		policy := split[0]
		password := strings.Trim(split[1], " ")
		letter, min, max, err := decodePolicy(policy)
		if err != nil {
			check(err)
		}
		if validatePassword(password, min, max, letter) {
			validCount++
		}
		if validatePasswordNew(password, min, max, letter) {
			validCountNew++
		}
	}
	fmt.Println(validCount)
	fmt.Println(validCountNew)

}

// decodePolicy returns the letter, and range of occurences
func decodePolicy(policy string) (string, int, int, error) {
	split := strings.Split(policy, " ")
	letter := split[1]

	minmax := strings.Trim(split[0], " ")
	minmaxsplit := strings.Split(minmax, "-")
	min, err := strconv.Atoi(minmaxsplit[0])
	if err != nil {
		return "", 0, 0, err
	}
	max, err := strconv.Atoi(minmaxsplit[1])
	if err != nil {
		return "", 0, 0, err
	}

	return letter, min, max, nil
}

func validatePassword(password string, min int, max int, letter string) bool {
	count := strings.Count(password, letter)
	if count >= min && count <= max {
		return true
	}
	return false
}

func validatePasswordNew(password string, min int, max int, letter string) bool {
	index1 := min - 1
	index2 := max - 1
	char1 := string(password[index1])
	char2 := string(password[index2])
	if (char1 == letter) != (char2 == letter) {
		return true
	}
	return false
}
