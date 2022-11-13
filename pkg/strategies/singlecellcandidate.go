package strategies

import (
	"github.com/whites11/sugoku/pkg/types"
)

type SingleCellCandidate struct {
}

func NewSingleCellCandidate() *SingleCellCandidate {
	return &SingleCellCandidate{}
}

func singleCellCandidateStrategy(section types.Section) *types.Cell {
	candidatesCount := make(map[uint8]uint8)
	candidatesCell := make(map[uint8]int)

	// Collect values in the row
	for i, cell := range section.Cells() {
		if cell.Value == nil {
			for _, candidate := range cell.Candidates {
				candidatesCount[candidate] = candidatesCount[candidate] + 1
				candidatesCell[candidate] = i
			}
		}
	}

	for val, occourences := range candidatesCount {
		if occourences == 1 {
			// In this section only one cell can have `val`.
			cell := section.Cells()[candidatesCell[val]]
			cell.Value = &val
			return &cell
		}
	}

	return nil
}

func (s *SingleCellCandidate) Step(b *types.Board) (*types.Cell, error) {
	// Rows
	{
		for _, row := range b.Rows() {
			cell := singleCellCandidateStrategy(&row)
			if cell != nil {
				return cell, nil
			}
		}
	}

	// Columns
	{
		for _, column := range b.Columns() {
			cell := singleCellCandidateStrategy(&column)
			if cell != nil {
				return cell, nil
			}
		}
	}

	// Squares
	{
		for _, square := range b.Squares() {
			cell := singleCellCandidateStrategy(&square)
			if cell != nil {
				return cell, nil
			}
		}
	}

	return nil, nil
}
