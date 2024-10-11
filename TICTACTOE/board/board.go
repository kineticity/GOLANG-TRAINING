package board

import (
	"fmt"
	"gameApp/cell"
)

type Board struct {
	cells [9]cell.Cell
}

func NewBoard() *Board {

	board := Board{}
	for i := 0; i < 9; i++ {
		board.cells[i] = cell.NewCell()
	}
	return &board
}

func (b *Board) PrintBoard() {
	fmt.Println("Current board:")
	for i := 0; i < 9; i++ {
		fmt.Printf("%s ", b.cells[i].GetSymbol())
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func (b *Board) IsValidMove(pos int) bool {
	return b.cells[pos].IsEmpty()
}

func (b *Board) MakeMove(pos int, symbol string) {
	b.cells[pos].SetSymbol(symbol)
}

func (b *Board) CheckWin() bool {
	winPositions := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, positions := range winPositions {
		if b.cells[positions[0]].GetSymbol() != "_" &&
			b.cells[positions[0]].GetSymbol() == b.cells[positions[1]].GetSymbol() &&
			b.cells[positions[1]].GetSymbol() == b.cells[positions[2]].GetSymbol() {
			return true
		}
	}
	return false
}

func (b *Board) IsDraw() bool {
	for i := 0; i < 9; i++ {
		if b.cells[i].IsEmpty() {
			return false
		}
	}
	return true
}
