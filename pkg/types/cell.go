package types

import "fmt"

type Cell struct {
	Value      *uint8
	Row        uint8
	Column     uint8
	Square     uint8
	Candidates []uint8
}

func (c *Cell) String() string {
	val := "_"
	if c.Value != nil {
		val = fmt.Sprint(*c.Value)
	}

	return fmt.Sprintf("%s %v", val, c.Candidates)
}
