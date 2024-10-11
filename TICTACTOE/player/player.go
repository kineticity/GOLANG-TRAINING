package player

import "errors"

type Player struct {
	name   string
	symbol string
}

func NewPlayer(name, symbol string) (*Player, error) {
	if name == "" {
		return nil, errors.New("player name cannot be empty")
	}
	if symbol != "X" && symbol != "O" {
		return nil, errors.New("symbol must be 'X' or 'O'")
	}
	return &Player{name: name, symbol: symbol}, nil
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetSymbol() string {
	return p.symbol
}
