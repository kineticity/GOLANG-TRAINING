package main

import (
	"fmt"
)

func main() {
	var g1 Game = NewTicTacToe("Manny", "Gloria")
	
	g1.play(0,2)
	g1.play(4)
	g1.play(6)
	g1.play(5)
	g1.play(4)
	g1.play(9)
	g1.play(3)

}

type Game interface {
	play(parameter ...interface{})
}

type TicTacToe struct {
	player1, player2 Player
	board            Board
	currentPlayer    Player
	gameOver         bool 
}

func NewTicTacToe(player1Name, player2Name string) *TicTacToe { //tictactoe factory
	player1 := Player{name: player1Name, symbol: "X"}
	player2 := Player{name: player2Name, symbol: "O"}
	board := Board{cells: make([]string, 9)}
	return &TicTacToe{
		player1:      player1,
		player2:      player2,
		board:        board,
		currentPlayer: player1,
		gameOver:     false,
	}
}

func (g *TicTacToe) play(parameter ...interface{}) {
	if g.gameOver {
		fmt.Println("Game is over! No more moves are allowed.")
		return
	}

	if len(parameter) != 1 { //validate only one parameter actually passed in the variadic parameter argument
		fmt.Println("Invalid number of arguments! Tic-Tac-Toe requires exactly one integer position (0-8).")
		return
	}

	pos, ok := parameter[0].(int) //validate if parameter integer between0 and 8
	if !ok || pos < 0 || pos > 8 {
		fmt.Println("Invalid move! Please enter a number between 0 and 8.")
		return
	}

	if !g.board.isValidMove(pos) {
		fmt.Println("Invalid move")
		return
	}

	g.board.makeMove(pos, g.currentPlayer.symbol)

	g.board.printBoard()

	if g.board.checkWin() {
		fmt.Printf("%s wins!\n", g.currentPlayer.name)
		g.gameOver = true 
		return
	}

	if g.board.isDraw() {
		fmt.Println("It's a draw!")
		g.gameOver = true 
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

type Player struct {
	name   string
	symbol string
}

type Board struct {
	cells []string
}

func (b *Board) printBoard() {
	fmt.Println("Current board:")
	for i := 0; i < 9; i++ {
		if b.cells[i] == "" {
			fmt.Print("_ ")
		} else {
			fmt.Printf("%s ", b.cells[i])
		}

		if (i+1)%3 == 0 { //new line after 012 345 678
			fmt.Println()
		}
	}
}

//POSITION SHOULD BE VALID (0-8) CELL SHOULD BE EMPTY 
func (b *Board) isValidMove(pos int) bool {
	return pos >= 0 && pos < 9 && b.cells[pos] == ""
}

func (b *Board) makeMove(pos int, symbol string) {
	b.cells[pos] = symbol
}

func (b *Board) checkWin() bool {
	winPositions := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6},            // Diagonals
	}

	for _, positions := range winPositions {
		if b.cells[positions[0]] != "" &&
			b.cells[positions[0]] == b.cells[positions[1]] &&
			b.cells[positions[1]] == b.cells[positions[2]] {
			return true
		}
	}
	return false
}

func (b *Board) isDraw() bool {
	for i := 0; i < 9; i++ {
		if b.cells[i] == "" {
			return false
		}
	}
	return true
}
