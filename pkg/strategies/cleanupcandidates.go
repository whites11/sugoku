package strategies

import (
	"github.com/whites11/sugoku/pkg/types"
)

type CleanupCandidates struct {
}

func NewCleanupCandidates() *CleanupCandidates {
	return &CleanupCandidates{}
}

func cleanupCandidatesStrategy(section types.Section) (*types.Cell, error) {
	// Collect values in the section.
	vals := make([]uint8, 0)
	for _, cell := range section.Cells() {
		if cell.Value != nil {
			vals = append(vals, *cell.Value)
		}
	}

	// Remove values from other cells in the section.
	for _, cell := range section.Cells() {
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

	return nil, nil
}

func (s *CleanupCandidates) Step(b *types.Board) (*types.Cell, error) {
	// Rows
	{
		for _, row := range b.Rows() {
			cell, err := cleanupCandidatesStrategy(&row)
			if err != nil {
				return nil, err
			}
			if cell != nil {
				return cell, nil
			}
		}
	}

	// Columns
	{
		for _, column := range b.Columns() {
			cell, err := cleanupCandidatesStrategy(&column)
			if err != nil {
				return nil, err
			}
			if cell != nil {
				return cell, nil
			}
		}
	}

	// Squares
	{
		for _, square := range b.Squares() {
			cell, err := cleanupCandidatesStrategy(&square)
			if err != nil {
				return nil, err
			}
			if cell != nil {
				return cell, nil
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
