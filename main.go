// Implementation of John Horton Conway's "Game of Life' as described on
// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	seed, err := cryptoSeed()
	if err != nil {
		log.Fatalf("seed: %v", err)
	}
	rand.Seed(seed)

	err = termbox.Init()
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

	game := NewGame(150, 50)
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
			time.Sleep(time.Second / 30)
		}
	}
}

func cryptoSeed() (int64, error) {
	var b [8]byte

	if _, err := crypto_rand.Read(b[:]); err != nil {
		return 0, errors.New("cannot seed math/rand package with cryptographically secure random number generator")
	}

	return int64(binary.LittleEndian.Uint64(b[:])), nil
}
