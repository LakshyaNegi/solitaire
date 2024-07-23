package main

import (
	"solitaire/game"
)

func main() {
	game := game.NewGame()

	game.Start()

	for game.IsRunning {
		game.Draw()
	}
}
