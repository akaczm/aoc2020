package main

import (
	"bufio"
	"fmt"
	"os"
)

// Plane holds its size as well as a 2d slice of its seats holding the occupancy status as a boolean.
type Plane struct {
	Row   int
	Col   int
	Seats [][]bool
}

// NewPlane creates a new plane and initializes and empty 2d array with all the seats free (false).
func NewPlane(x int, y int) Plane {
	xdim := make([][]bool, x)
	for ydim := range xdim {
		xdim[ydim] = make([]bool, y)
	}
	plane := Plane{x, y, xdim}
	return plane
}

// Marks a designated seat as taken, returns changed plane struct.
func (p Plane) takeSeat(x int, y int) Plane {
	p.Seats[x][y] = true
	return p
}

// BoardingPass is a struct that holds the instructions of a boarding pass.
type BoardingPass struct {
	SeatRow string
	SeatCol string
}

// NewBoardingPass creates a new BoardingPass struct by decoding the raw data.
func NewBoardingPass(input string) BoardingPass {
	boardingPass := BoardingPass{input[:7], input[7:]}
	return boardingPass
}

// findSeat calculates the seat a boarding pass points towards.
func (pass BoardingPass) findSeat(plane Plane) (int, int) {
	row := findRow(pass.SeatRow, plane.Row)
	col := findColumn(pass.SeatCol, plane.Col)
	return row, col
}

// getID gets the ID of a seat.
func getID(row int, col int) int {
	return row*8 + col
}

// findRow finds the row a boarding pass points to.
func findRow(pass string, rowCount int) int {
	rows := makeRange(0, rowCount-1)
	for _, instruction := range pass {
		switch string(instruction) {
		case "B":
			rows = rows[len(rows)/2:]
		case "F":
			rows = rows[:len(rows)/2]
		}

	}
	return rows[0]
}

// findColumn finds the column a boarding pass points to.
func findColumn(pass string, colCount int) int {
	cols := makeRange(0, colCount-1)
	for _, instruction := range pass {
		switch string(instruction) {
		case "R":
			cols = cols[len(cols)/2:]
		case "L":
			cols = cols[:len(cols)/2]
		}

	}
	return cols[0]
}

// findEmpty finds the empty seats in a plane and returns their IDs.
func (p Plane) findEmpty() []int {
	emptySeatIDs := make([]int, 0)
	for x, row := range p.Seats {
		for y, col := range p.Seats[x] {
			if !row[y] && !col {
				id := getID(x, y)
				emptySeatIDs = append(emptySeatIDs, id)
			}
		}
	}
	return emptySeatIDs
}

// makeRange is a helper function to create an array of numbers within the specified range.
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	plane := NewPlane(128, 8)
	boardingPasses := make([]BoardingPass, 0)
	// Read data and create array of boarding passes.
	for scanner.Scan() {
		boardingPasses = append(boardingPasses, NewBoardingPass(scanner.Text()))
	}

	var highestID int
	takenIDs := make([]int, 0)
	for _, pass := range boardingPasses {
		row, col := pass.findSeat(plane)
		plane = plane.takeSeat(row, col)
		id := getID(row, col)
		takenIDs = append(takenIDs, id)
		if id > highestID {
			highestID = id
		}
	}
	fmt.Println(highestID)
	emptyseatsIDs := plane.findEmpty()
	potentialSeatIDs := make([]int, 0)

	// Certainly this can be made better with recurrency.
	for _, takenseatID := range takenIDs {
		for _, emptyseatID := range emptyseatsIDs {
			if takenseatID-1 == emptyseatID {
				potentialSeatIDs = append(potentialSeatIDs, emptyseatID)
			}
		}
	}
	// I miss Python's map and list comprehensions.

	// Take the seats that match previous condition and iterate through again.
	for _, takenseatID := range takenIDs {
		for _, emptyseatID := range potentialSeatIDs {
			if takenseatID+1 == emptyseatID {
				fmt.Println(emptyseatID)
			}
		}
	}
}
