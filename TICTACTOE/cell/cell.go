package cell

type Cell struct {
	symbol  string
	isEmpty bool
}

func NewCell() Cell {
	return Cell{
		symbol:  "_",
		isEmpty: true,
	}
}

func (c *Cell) GetSymbol() string {
	return c.symbol
}

func (c *Cell) SetSymbol(symbol string) {
	c.symbol = symbol
	c.isEmpty = (symbol == "_")
}

func (c *Cell) IsEmpty() bool {
	return c.isEmpty
}
