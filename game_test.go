package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	got := New(10, 10)

	require.NotNil(t, got.currentGrid)
	require.NotNil(t, got.nextGrid)
	require.NotSame(t, got.currentGrid, got.nextGrid)
	require.Zero(t, got.gen)
}

func TestStep(t *testing.T) {
	// Check if the Blinker oscillator works.
	// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#/media/File:Game_of_life_blinker.gif
	grid := &Grid{
		width:  5,
		height: 5,
		cells: [][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		},
	}

	want := &Grid{
		width:  5,
		height: 5,
		cells: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}

	g := &Game{
		currentGrid: grid,
		nextGrid:    NewGrid(grid.width, grid.height, true), // Random populated next grid.
	}
	g.Step()
	require.Equal(t, want, g.currentGrid)
}

func TestString(t *testing.T) {
	want := "+ \n +\n"
	grid := &Grid{
		width:  2,
		height: 2,
		cells: [][]bool{
			{true, false},
			{false, true},
		},
	}
	g := &Game{
		currentGrid: grid,
	}

	got := g.String()
	require.Equal(t, want, got)
}
