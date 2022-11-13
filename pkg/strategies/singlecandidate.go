package strategies

import (
	"github.com/whites11/sugoku/pkg/types"
)

type SingleCandidate struct {
}

func NewSingleCandidate() *SingleCandidate {
	return &SingleCandidate{}
}

func (s *SingleCandidate) Step(b *types.Board) (*types.Cell, error) {
	for _, cell := range b.Cells() {
		if cell.Value == nil && len(cell.Candidates) == 1 {
			val := cell.Candidates[0]
			cell.Value = &val
			return &cell, nil
		}
	}

	return nil, nil
}
