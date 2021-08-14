package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	grid := NewGrid(10, 10, false)

	g := New(grid)

	t.Run("current grid", func(t *testing.T) {
		require.Equal(t, grid, g.currentGrid)
	})

	t.Run("next grid", func(t *testing.T) {
		require.Equal(t, grid, g.nextGrid)
	})
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

	g := New(grid)
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
	g := New(grid)

	got := g.String()
	require.Equal(t, want, got)
}
