// Implementation of John Horton Conway's "Game of Life' as described on
// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	seed, err := cryptoSeed()
	if err != nil {
		log.Fatalf("seed: %v", err)
	}
	rand.Seed(seed)

	game := New(150, 50)
	for i := 0; i < 10000; i++ {
		clear()

		game.Step()
		fmt.Print(game)
		time.Sleep(time.Second / 30)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cryptoSeed() (int64, error) {
	var b [8]byte

	if _, err := crypto_rand.Read(b[:]); err != nil {
		return 0, errors.New("cannot seed math/rand package with cryptographically secure random number generator")
	}

	return int64(binary.LittleEndian.Uint64(b[:])), nil
}
