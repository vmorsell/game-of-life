package main

import "bytes"

// Game is the main struct for this application.
// It holds a reference to the grid for the current generation.
type Game struct {
	grid *Grid
	gen  int
}

// New returns a new Game instance with the provided Grid.
func New(grid *Grid) *Game {
	return &Game{
		grid: grid,
	}

}

// Step moves the game to the next generation.
func (g *Game) Step() {
	nextGrid := NewGrid(g.grid.width, g.grid.height, false)

	for y := range g.grid.cells {
		for x := range g.grid.cells[y] {
			if live := g.grid.IsLiveNextGen(x, y); live {
				nextGrid.cells[y][x] = true
			}
		}
	}

	g.grid = nextGrid
	g.gen++
}

// String returns a string visualization of the game's current state.
func (g *Game) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.grid.height; y++ {
		for x := 0; x < g.grid.width; x++ {
			b := byte(' ')
			if g.grid.IsLive(x, y) {
				b = '+'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
