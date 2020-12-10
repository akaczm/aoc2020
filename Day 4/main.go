// https://adventofcode.com/2020/day/4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Passport is a structure that holds passport information and allows validation of data.
type Passport struct {
	Birthyear  string
	Issueyear  string
	Expyear    string
	Height     string
	Haircolor  string
	Eyecolor   string
	PassportID string
	CountryID  string
}

type ValueValidatorPair struct {
	Val   Validator
	Value string
}

// GetFields returns the fields of Passport struct as a map.
func (p Passport) GetFields() map[string]string {
	fields := make(map[string]string)
	fields["byr"] = p.Birthyear
	fields["iyr"] = p.Issueyear
	fields["eyr"] = p.Expyear
	fields["hgt"] = p.Height
	fields["hcl"] = p.Haircolor
	fields["ecl"] = p.Eyecolor
	fields["pid"] = p.PassportID
	fields["cid"] = p.CountryID

	return fields
}

// Validator interface expects a string input, returns a boolean and error if applicable.
type Validator interface {
	Validate(data string) (bool, error)
}

// ValidateBirthyear validates the birthyear according to requirements.
type ValidateBirthyear struct{}

// Validate validates according to requirements.
func (v ValidateBirthyear) Validate(data string) (bool, error) {
	if data == "" {
		return false, nil
	}
	birthyear, err := strconv.Atoi(data)
	if err != nil {
		return false, err
	}
	if birthyear < 1920 || birthyear > 2002 {
		return false, nil
	}
	return true, nil
}

type ValidateIssueyear struct{}

// Validate validates according to requirements.
func (v ValidateIssueyear) Validate(data string) (bool, error) {
	if data == "" {
		return false, nil
	}
	issueyear, err := strconv.Atoi(data)
	if err != nil {
		return false, err
	}
	if issueyear < 2010 || issueyear > 2020 {
		return false, nil
	}
	return true, nil
}

type ValidateExpyear struct{}

// Validate validates according to requirements.
func (v ValidateExpyear) Validate(data string) (bool, error) {
	if data == "" {
		return false, nil
	}
	expyear, err := strconv.Atoi(data)
	if err != nil {
		return false, err
	}
	if expyear < 2020 || expyear > 2030 {
		return false, nil
	}
	return true, nil
}

type ValidateHeight struct{}

// Validate validates according to requirements.
func (v ValidateHeight) Validate(data string) (bool, error) {
	if data == "" {
		return false, nil
	}
	valuestr := data[:len(data)-2]
	value, err := strconv.Atoi(valuestr)
	if err != nil {
		return false, err
	}
	suffix := data[len(data)-2:]
	if suffix == "in" {
		if value < 59 || value > 76 {
			return false, nil
		}
	} else if suffix == "cm" {
		if value < 150 || value > 193 {
			return false, nil
		}
	}
	return true, nil
}

type ValidateHaircolor struct{}

// Validate validates according to requirements.
func (v ValidateHaircolor) Validate(data string) (bool, error) {
	const validchars string = "0123456789abcdef"

	if data == "" {
		return false, nil
	}

	prefix := string(data[0])
	value := data[1:]
	if prefix != "#" {
		return false, nil
	}
	if len(value) != 6 {
		return false, nil
	}
	for _, character := range value {
		if !strings.Contains(validchars, string(character)) {
			return false, nil
		}
	}
	return true, nil
}

type ValidateEyecolor struct{}

// Validate validates according to requirements.
func (v ValidateEyecolor) Validate(data string) (bool, error) {
	switch data {
	case "amb":
		return true, nil
	case "blu":
		return true, nil
	case "brn":
		return true, nil
	case "gry":
		return true, nil
	case "grn":
		return true, nil
	case "hzl":
		return true, nil
	case "oth":
		return true, nil
	}
	return false, nil
}

type ValidatePID struct{}

// Validate validates according to requirements.
func (v ValidatePID) Validate(data string) (bool, error) {
	_, err := strconv.Atoi(data)
	if err != nil {
		return false, errors.Wrap(err, "error decoding PID")
	}
	if len(data) == 9 {
		return true, nil
	}
	return false, nil
}

type ValidateCID struct{}

// Validate validates according to requirements.
func (v ValidateCID) Validate(data string) (bool, error) {
	return true, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDigits(i int) int {
	if i < 10 {
		return 1
	}
	return 1 + getDigits(i/10)

}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	passports := make([]Passport, 0)

	scanbuffer := make([]string, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			decodedpassport, err := decodePassport(scanbuffer)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				passports = append(passports, decodedpassport)
				scanbuffer = nil
			}
		} else {
			split := strings.Split(scanner.Text(), " ")
			for _, item := range split {
				scanbuffer = append(scanbuffer, item)
			}
		}
	}
	decodedpassport, err := decodePassport(scanbuffer)
	if err != nil {
		fmt.Println(err)
	} else {
		passports = append(passports, decodedpassport)
	}
	scanbuffer = nil
	var validPassports int
	// Simple check for part 1
	for _, passport := range passports {
		if validatePassport(passport) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
	// In-depth check for part 2
	var validPassportsExtensive int
	for _, passport := range passports {
		var validatorPair []ValueValidatorPair
		isValid := false
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateBirthyear{}, passport.Birthyear})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateCID{}, passport.CountryID})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateExpyear{}, passport.Expyear})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateEyecolor{}, passport.Eyecolor})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateHaircolor{}, passport.Haircolor})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateHeight{}, passport.Height})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidateIssueyear{}, passport.Issueyear})
		validatorPair = append(validatorPair, ValueValidatorPair{ValidatePID{}, passport.PassportID})
		for _, pair := range validatorPair {
			result, err := validatePassportExtensive(pair.Val, pair.Value)
			if err != nil {
				isValid = false
				break
			}
			if result == false {
				isValid = false
				break
			} else if result == true {
				isValid = true
			}
		}
		if isValid {
			validPassportsExtensive++
		}
	}
	fmt.Println(validPassportsExtensive)
}

func decodePassport(raw []string) (Passport, error) {
	passport := Passport{}
	for _, item := range raw {
		itemsplit := strings.Split(item, ":")
		switch itemsplit[0] {
		case "byr":
			passport.Birthyear = itemsplit[1]
		case "iyr":
			passport.Issueyear = itemsplit[1]
		case "eyr":
			passport.Expyear = itemsplit[1]
		case "hgt":
			passport.Height = itemsplit[1]
		case "hcl":
			passport.Haircolor = itemsplit[1]
		case "ecl":
			passport.Eyecolor = itemsplit[1]
		case "pid":
			passport.PassportID = itemsplit[1]
		case "cid":
			passport.CountryID = itemsplit[1]
		}
	}
	return passport, nil
}

func validatePassport(passport Passport) bool {
	for key, field := range passport.GetFields() {
		if field == "" {
			if key != "cid" {
				return false
			}
		}
	}
	return true
}

func validatePassportExtensive(validator Validator, data string) (bool, error) {
	return validator.Validate(data)
}
