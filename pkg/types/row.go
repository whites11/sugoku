package types

import "fmt"

type Row struct {
	index uint8
	data  []Cell
}

func (r *Row) Cells() []Cell {
	return r.data
}

func (r *Row) String() string {
	ret := ""
	for _, c := range r.data {
		ret = ret + fmt.Sprintf("%s ", c.String())
	}

	return ret
}
