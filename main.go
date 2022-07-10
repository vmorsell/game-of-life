// Implementation of John Horton Conway's "Game of Life' as described on
// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	width                = 80
	height               = 25
	freq   time.Duration = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		log.Fatalf("termbox: %v", err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	game := NewGame(width, height)
	game.Render()

loop:
	for i := 0; i < 10000; i++ {
		select {
		case ev := <-eventQueue:
			if ev.Ch == 'q' || ev.Key == termbox.KeyEsc {
				break loop
			}

		default:
			game.Render()
			game.Step()
			time.Sleep(time.Second / freq)
		}
	}
}
