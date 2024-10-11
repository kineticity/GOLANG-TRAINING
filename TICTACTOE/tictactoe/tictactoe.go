package tictactoe

import (
	"errors"
	"fmt"
	"gameApp/board"
	"gameApp/player"
)

type TicTacToe struct {
	player1, player2 player.Player
	board            board.Board
	currentPlayer    player.Player
	status           string
	winner           string
}

func NewTicTacToe(player1Name, player2Name string) (*TicTacToe, error) {
	player1, err := player.NewPlayer(player1Name, "O")
	if err != nil {
		return nil, err
	}
	player2, err := player.NewPlayer(player2Name, "X")
	if err != nil {
		return nil, err
	}
	gameBoard := board.NewBoard()

	return &TicTacToe{
		player1:       *player1,
		player2:       *player2,
		board:         *gameBoard,
		currentPlayer: *player1,
		status:        "ongoing",
		winner:        "No one",
	}, nil
}

func (g *TicTacToe) Play(parameter ...interface{}) {
	if g.GetStatus() != "ongoing" {
		fmt.Println("Game is over! No more moves allowed.")
		fmt.Printf("The game was a %s. %s is the winner.\n", g.GetStatus(), g.GetWinner())
		return
	}

	if len(parameter) != 1 {
		fmt.Println("Invalid number of arguments! Tic-Tac-Toe requires exactly one integer position (0-8).")
		return
	}

	pos, ok := parameter[0].(int)
	if !ok || pos < 0 || pos > 8 {
		fmt.Println("Invalid move! Please enter a number between 0 and 8.")
		return
	}

	if !g.board.IsValidMove(pos) {
		fmt.Println("Invalid move, position already taken.")
		return
	}

	g.board.MakeMove(pos, g.currentPlayer.GetSymbol())
	g.board.PrintBoard()

	if g.board.CheckWin() {
		fmt.Printf("%s wins!\n", g.currentPlayer.GetName())
		g.SetStatus("win")
		g.SetWinner(g.currentPlayer.GetName())
		return
	}

	if g.board.IsDraw() {
		fmt.Println("It's a draw!")
		g.SetStatus("draw")
		// g.SetWinner("No one")
		return
	}

	g.switchPlayer()
}

func (g *TicTacToe) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}

func (g *TicTacToe) SetStatus(status string) error {
	if status != "ongoing" && status != "draw" && status != "win" {
		return errors.New("invalid status")
	}
	g.status = status
	return nil
}

func (g *TicTacToe) GetStatus() string {
	return g.status
}

func (g *TicTacToe) SetWinner(winner string) error {
	if winner != g.player1.GetName() && winner != g.player2.GetName() && winner != "No one" {
		return errors.New("invalid winner")
	}
	g.winner = winner
	return nil
}

func (g *TicTacToe) GetWinner() string {
	return g.winner
}
