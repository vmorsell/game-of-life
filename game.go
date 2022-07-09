package main

import (
	"github.com/nsf/termbox-go"
)

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

// Render uses Termbox to render the current state of the game.
func (g *Game) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < g.currentGrid.height; y++ {
		for x := 0; x < g.currentGrid.width; x++ {
			if g.currentGrid.IsLive(x, y) {
				termbox.SetCell(x, y, '*', termbox.ColorGreen, termbox.ColorDefault)
			}
		}
	}
	termbox.Flush()
}
