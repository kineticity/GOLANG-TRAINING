package game

type Game interface {
	Play(parameter ...interface{})
	GetStatus() string
	GetWinner() string
}
