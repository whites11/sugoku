package types

type Column struct {
	index uint8
	data  []Cell
}

func (r *Column) Cells() []Cell {
	return r.data
}
