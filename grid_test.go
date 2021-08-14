package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGrid(t *testing.T) {
	t.Run("not populated", func(t *testing.T) {
		g := NewGrid(10, 10, false)

		for i := 0; i < g.width; i++ {
			for j := 0; j < g.height; j++ {
				require.False(t, g.cells[j][i])
			}
		}
	})

	t.Run("populated", func(t *testing.T) {
		width := 10
		height := 10

		g := NewGrid(width, height, true)

		want := width * height / 5
		live := 0
		for i := 0; i < g.width; i++ {
			for j := 0; j < g.height; j++ {
				if g.cells[j][i] == true {
					live++
				}
			}
		}

		// Accept an error of 10%
		require.InDelta(t, want, live, float64(width*height)*0.1)
	})
}

func TestSetLive(t *testing.T) {
	g := NewGrid(10, 10, false)

	x := 5
	y := 2

	g.SetLive(x, y, true)

	require.True(t, g.cells[y][x])
}

func TestIsLive(t *testing.T) {
	g := NewGrid(10, 10, false)

	x := 5
	y := 2

	g.cells[y][x] = true

	got := g.IsLive(x, y)
	require.True(t, got)
}

type XY struct {
	x, y int
}

func TestIsLiveNextGen(t *testing.T) {
	tests := []struct {
		name       string
		grid       *Grid
		x          int
		y          int
		wantIfLive bool
		wantIfDead bool
	}{
		{
			name:       "no live neighbours",
			grid:       gridWithLiveCells(nil),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "one live neighbour",
			grid: gridWithLiveCells([]XY{
				{5, 4},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "two live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
			}),
			x:          5,
			y:          5,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "three live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
			}),
			x:          5,
			y:          5,
			wantIfLive: true,
			wantIfDead: true,
		},
		{
			name: "four live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
				{6, 5},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "five live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
				{6, 5},
				{4, 4},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "six live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
				{6, 5},
				{4, 4},
				{6, 6},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "seven live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
				{6, 5},
				{4, 4},
				{6, 6},
				{4, 6},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "eight live neighbours",
			grid: gridWithLiveCells([]XY{
				{5, 4},
				{5, 6},
				{4, 5},
				{6, 5},
				{4, 4},
				{6, 6},
				{4, 6},
				{6, 4},
			}),
			x:          5,
			y:          5,
			wantIfLive: false,
			wantIfDead: false,
		},
		{
			name: "grid underflow x",
			grid: gridWithLiveCells([]XY{
				{0, 4},
				{9, 4},
			}),
			x:          0,
			y:          5,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "grid underflow y",
			grid: gridWithLiveCells([]XY{
				{5, 1},
				{4, 9},
			}),
			x:          5,
			y:          0,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "grid underflow x and y",
			grid: gridWithLiveCells([]XY{
				{9, 0},
				{9, 9},
			}),
			x:          0,
			y:          0,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "grid overflow x",
			grid: gridWithLiveCells([]XY{
				{0, 4},
				{9, 4},
			}),
			x:          9,
			y:          5,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "grid overflow y",
			grid: gridWithLiveCells([]XY{
				{5, 0},
				{4, 9},
			}),
			x:          5,
			y:          9,
			wantIfLive: true,
			wantIfDead: false,
		},
		{
			name: "grid underflow x and y",
			grid: gridWithLiveCells([]XY{
				{9, 0},
				{0, 9},
			}),
			x:          9,
			y:          9,
			wantIfLive: true,
			wantIfDead: false,
		},
	}

	for _, tt := range tests {
		t.Run("if cell is live", func(t *testing.T) {
			tt.grid.SetLive(tt.x, tt.y, true)
			got := tt.grid.IsLiveNextGen(tt.x, tt.y)
			require.Equal(t, tt.wantIfLive, got)
		})

		t.Run("if cell is dead", func(t *testing.T) {
			tt.grid.SetLive(tt.x, tt.y, false)
			got := tt.grid.IsLiveNextGen(tt.x, tt.y)
			require.Equal(t, tt.wantIfDead, got)
		})
	}
}

func gridWithLiveCells(cells []XY) *Grid {
	g := NewGrid(10, 10, false)

	for _, c := range cells {
		g.SetLive(c.x, c.y, true)
	}

	return g
}
