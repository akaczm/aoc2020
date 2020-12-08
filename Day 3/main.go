package main

import (
	"bufio"
	"fmt"
	"os"
)

//Toboggan has a determined position
type Toboggan struct {
	PosX int
	PosY int
}

//Move moves the toboggan by a specified amount
func (t *Toboggan) Move(x int, y int) {
	t.PosX += x
	t.PosY += y
}

//Reset resets the toboggan to starting position
func (t *Toboggan) Reset() {
	t.PosX = 0
	t.PosY = 0
}

//Forest is ultimately a useless structure but I'm just learning Go
type Forest struct {
	Row []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	forest := &Forest{}
	toboggan := &Toboggan{}
	trees := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		val := scanner.Text()

		forest.Row = append(forest.Row, val)
	}

	trees = append(trees, traverse(forest, toboggan, 1, 1))
	trees = append(trees, traverse(forest, toboggan, 3, 1))
	trees = append(trees, traverse(forest, toboggan, 5, 1))
	trees = append(trees, traverse(forest, toboggan, 7, 1))
	trees = append(trees, traverse(forest, toboggan, 1, 2))

	arraysize := len(trees)

	result := multiply(trees, arraysize-1)

	fmt.Println(result)

}

func traverse(forest *Forest, toboggan *Toboggan, speedX int, speedY int) int {
	var treesEncountered int
	for index, row := range forest.Row {
		if index == toboggan.PosY {
			if string(row[toboggan.PosX]) == "#" {
				treesEncountered++
			}
			toboggan.Move(speedX, speedY)
			if toboggan.PosX >= len(row) {
				toboggan.PosX -= len(row)
			}
		}
	}
	toboggan.Reset()
	return treesEncountered
}

func multiply(array []int, n int) int {
	if n == 0 {
		return array[n]
	}
	return array[n] * multiply(array, n-1)
}
