package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/saleh/game/game"
)

func main() {

	g := game.NewGame()
	err := ebiten.RunGame(g)
	panicError(err)
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
