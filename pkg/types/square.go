package types

import (
	"fmt"
)

type Square struct {
	index uint8
	data  []Cell
}

func (s *Square) Cells() []Cell {
	return s.data
}

func (s *Square) String() string {
	return fmt.Sprintf(
		"%d %d %d\n%d %d %d\n%d %d %d\n",
		s.data[0], s.data[1], s.data[2],
		s.data[3], s.data[4], s.data[5],
		s.data[6], s.data[7], s.data[8],
	)
}
