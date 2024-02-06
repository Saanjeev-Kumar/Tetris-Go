package main

import (
	"log"
	"math/rand"
	"time"

	"tetris_Go/sanju"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	game := sanju.NewGame()

	ebiten.SetWindowSize(sanju.ScreenWidth, sanju.ScreenHeight)
	ebiten.SetWindowTitle("gtris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
