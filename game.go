package main

import "bytes"

// Game is the main struct for this application.
type Game struct {
	currentGrid *Grid
	nextGrid    *Grid
	gen         int
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		currentGrid: NewGrid(width, height, true),
		nextGrid:    NewGrid(width, height, false),
	}

}

// Step moves the game to the next generation.
func (g *Game) Step() {
	for y := range g.currentGrid.cells {
		for x := range g.currentGrid.cells[y] {
			g.nextGrid.SetLive(x, y, g.currentGrid.IsLiveNextGen(x, y))
		}
	}

	// Switch nextGrid and currentGrid
	g.currentGrid, g.nextGrid = g.nextGrid, g.currentGrid
	g.gen++
}

// String returns a string visualization of the game's current state.
func (g *Game) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.currentGrid.height; y++ {
		for x := 0; x < g.currentGrid.width; x++ {
			b := byte(' ')
			if g.currentGrid.IsLive(x, y) {
				b = '+'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
