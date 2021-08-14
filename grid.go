package main

import (
	"math/rand"
)

// Grid is the two-dimensional representation of all cells.
// It uses the coordinate system [y][x].
type Grid struct {
	width  int
	height int
	cells  [][]bool
}

// NewGrid returns a new Grid instance and optionally populates it with
// approximate 20 % live cells.
func NewGrid(width, height int, populate bool) *Grid {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	grid := &Grid{
		width:  width,
		height: height,
		cells:  cells,
	}

	if populate {
		for i := 0; i < (width * height / 5); i++ {
			grid.SetLive(rand.Intn(width), rand.Intn(height), true)
		}
	}

	return grid
}

// SetLive sets the 'live' state of a cell.
func (g *Grid) SetLive(x, y int, v bool) {
	g.cells[y][x] = v
}

// IsLive returns the 'live' state of a cell.
func (g *Grid) IsLive(x, y int) bool {
	return g.cells[y][x]
}

// CellShouldLive determines if the cell should be live in the next generation
// of the game based on its eight neighbours.
func (g *Grid) IsLiveNextGen(x, y int) bool {
	live := 0
	for xx := x - 1; xx <= x+1; xx++ {
		for yy := y - 1; yy <= y+1; yy++ {
			// Don't count the current cell.
			if yy == y && xx == x {
				continue
			}

			// Overflow to the other side of the grid.
			xxx := xx
			xxx += g.width
			xxx %= g.width

			yyy := yy
			yyy += g.height
			yyy %= g.height

			if g.IsLive(xxx, yyy) {
				live++
			}
		}
	}

	// Simplified rules for next generation:
	// 1. Any cell with three live neighbours is live next generation.
	if live == 3 {
		return true
	}

	// 2. Any live cell with two live neighbours stays alive.
	if g.IsLive(x, y) && live == 2 {
		return true
	}

	// 3. All other cells die.
	return false
}
