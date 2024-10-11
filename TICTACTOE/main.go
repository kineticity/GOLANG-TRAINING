package main

import (
	"fmt"
	"gameApp/game"
	"gameApp/tictactoe"
)

func main() {
	var g1 game.Game
	g1, err := tictactoe.NewTicTacToe("Manny", "Gloria")
	if err != nil {
		fmt.Println(err)
		return
	}
	g1.Play(0)
	g1.Play(1)
	g1.Play(3)
	g1.Play(4)
	g1.Play(6) //Player 1 wins!
	g1.Play(2)

	g2, err := tictactoe.NewTicTacToe("Nancy", "Shane")
	if err != nil {
		fmt.Println(err)
		return
	}

	g2.Play(1)
	g2.Play(0)
	g2.Play(3)
	g2.Play(2)
	g2.Play(5)
	g2.Play(4)
	g2.Play(6)
	g2.Play(7)
	g2.Play(8) //Game ends in a draw!
	g2.Play(3)
}
