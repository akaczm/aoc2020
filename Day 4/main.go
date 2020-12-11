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

func NewPassport(input []string, validators []Validator) (Passport, error) {
	passport, err := decodePassport(input)

	if err != nil {
		return Passport{}, errors.Wrap(err, "error decoding passport")
	}

	for _, validator := range validators {
		valid, err := validator.Validate(passport)
		if err != nil {
			return Passport{}, errors.Wrap(err, "error validating password")
		}
		if !valid {
			return Passport{}, &PassportInvalidError{}
		}
	}
	return passport, nil
}

type PassportInvalidError struct{}

func (e *PassportInvalidError) Error() string {
	return "Passport Invalid"
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
	Validate(passport Passport) (bool, error)
}

// ValidateBirthyear validates the birthyear according to requirements.
type ValidateBirthyear struct{}

// Validate validates according to requirements.
func (v ValidateBirthyear) Validate(passport Passport) (bool, error) {
	if passport.Birthyear == "" {
		return false, nil
	}
	birthyear, err := strconv.Atoi(passport.Birthyear)
	if err != nil {
		return false, errors.Wrap(err, "error decoding birth year")
	}
	if birthyear < 1920 || birthyear > 2002 {
		return false, nil
	}
	return true, nil
}

type ValidateIssueyear struct{}

// Validate validates according to requirements.
func (v ValidateIssueyear) Validate(passport Passport) (bool, error) {
	if passport.Issueyear == "" {
		return false, nil
	}
	issueyear, err := strconv.Atoi(passport.Issueyear)
	if err != nil {
		return false, errors.Wrap(err, "error decoding issue year")
	}
	if issueyear < 2010 || issueyear > 2020 {
		return false, nil
	}
	return true, nil
}

type ValidateExpyear struct{}

// Validate validates according to requirements.
func (v ValidateExpyear) Validate(passport Passport) (bool, error) {
	if passport.Expyear == "" {
		return false, nil
	}
	expyear, err := strconv.Atoi(passport.Expyear)
	if err != nil {
		return false, errors.Wrap(err, "error decoding expiration year")
	}
	if expyear < 2020 || expyear > 2030 {
		return false, nil
	}
	return true, nil
}

type ValidateHeight struct{}

// Validate validates according to requirements.
func (v ValidateHeight) Validate(passport Passport) (bool, error) {
	if passport.Height == "" {
		return false, nil
	}
	valuestr := passport.Height[:len(passport.Height)-2]
	value, err := strconv.Atoi(valuestr)
	if err != nil {
		return false, err
	}
	suffix := passport.Height[len(passport.Height)-2:]
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
func (v ValidateHaircolor) Validate(passport Passport) (bool, error) {
	const validchars string = "0123456789abcdef"

	if passport.Haircolor == "" {
		return false, nil
	}

	prefix := string(passport.Haircolor[0])
	value := passport.Haircolor[1:]
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
func (v ValidateEyecolor) Validate(passport Passport) (bool, error) {
	switch passport.Eyecolor {
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
func (v ValidatePID) Validate(passport Passport) (bool, error) {
	_, err := strconv.Atoi(passport.PassportID)
	if err != nil {
		return false, errors.Wrap(err, "error decoding PID")
	}
	if len(passport.PassportID) == 9 {
		return true, nil
	}
	return false, nil
}

type ValidateCID struct{}

// Validate validates according to requirements.
func (v ValidateCID) Validate(passport Passport) (bool, error) {
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
	validators := []Validator{
		ValidateBirthyear{},
		ValidateCID{},
		ValidateExpyear{},
		ValidateEyecolor{},
		ValidateHaircolor{},
		ValidateHeight{},
		ValidateIssueyear{},
		ValidatePID{},
	}

	validPassports := make([]Passport, 0)
	f, err := os.Open("testinput")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanbuffer := make([]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			decodedpassport, err := NewPassport(scanbuffer, validators)
			if _, ok := err.(*PassportInvalidError); ok {
				//continue
			}
			if err != nil {
				fmt.Println(err)
			} else {
				validPassports = append(validPassports, decodedpassport)
				scanbuffer = nil
			}
		} else {
			split := strings.Split(scanner.Text(), " ")
			for _, item := range split {
				scanbuffer = append(scanbuffer, item)
			}
		}
	}
	decodedpassport, err := NewPassport(scanbuffer, validators)
	if _, ok := err.(*PassportInvalidError); ok {
		//continue
	}
	if err != nil {
		fmt.Println(err)
	} else {
		validPassports = append(validPassports, decodedpassport)
	}
	fmt.Println(len(validPassports))

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
