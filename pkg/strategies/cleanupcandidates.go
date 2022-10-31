package strategies

import (
	"reflect"

	"github.com/whites11/sugoku/pkg/types"
)

type CleanupCandidates struct {
}

func NewCleanupCandidates() *CleanupCandidates {
	return &CleanupCandidates{}
}

func (s *CleanupCandidates) Step(b *types.Board) (*types.Cell, error) {
	// Rows
	{
		for _, row := range b.Rows() {
			// Collect values in the row
			vals := make([]uint8, 0)
			for _, cell := range row.Cells() {
				if cell.Value != nil {
					vals = append(vals, *cell.Value)
				}
			}

			// Remove values from other cells in the row.
			for _, cell := range row.Cells() {
				if cell.Value == nil {
					newCandidates := cell.Candidates
					for _, candidate := range vals {
						newCandidates = remove(newCandidates, candidate)
					}

					if !reflect.DeepEqual(newCandidates, cell.Candidates) {
						cell.Candidates = newCandidates
						return &cell, nil
					}
				}
			}
		}
	}

	// Columns
	{
		for _, column := range b.Columns() {
			// Collect values in the column
			vals := make([]uint8, 0)
			for _, cell := range column.Cells() {
				if cell.Value != nil {
					vals = append(vals, *cell.Value)
				}
			}

			// Remove values from other cells in the row.
			for _, cell := range column.Cells() {
				if cell.Value == nil {
					newCandidates := cell.Candidates
					for _, candidate := range vals {
						newCandidates = remove(newCandidates, candidate)
					}

					if len(newCandidates) != len(cell.Candidates) {
						cell.Candidates = newCandidates
						return &cell, nil
					}
				}
			}
		}
	}

	// Squares
	{
		for _, square := range b.Squares() {
			// Collect values in the square
			vals := make([]uint8, 0)
			for _, cell := range square.Cells() {
				if cell.Value != nil {
					vals = append(vals, *cell.Value)
				}
			}

			// Remove values from other cells in the row.
			for _, cell := range square.Cells() {
				if cell.Value == nil {
					newCandidates := cell.Candidates
					for _, candidate := range vals {
						newCandidates = remove(newCandidates, candidate)
					}

					if len(newCandidates) != len(cell.Candidates) {
						cell.Candidates = newCandidates
						return &cell, nil
					}
				}
			}
		}
	}

	return nil, nil
}

func remove(s []uint8, val uint8) []uint8 {
	ret := make([]uint8, 0)
	for _, v := range s {
		if v != val {
			ret = append(ret, v)
		}
	}
	return ret
}
