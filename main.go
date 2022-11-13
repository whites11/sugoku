package main

import (
	"fmt"

	"github.com/whites11/sugoku/pkg/solver"
	"github.com/whites11/sugoku/pkg/strategies"
	"github.com/whites11/sugoku/pkg/types"
)

func main() {
	board, err := types.NewBoard([]*uint8{
		nil, intPtr(2), nil, nil, nil, intPtr(1), nil, intPtr(8), intPtr(4),
		nil, nil, intPtr(8), intPtr(6), nil, nil, intPtr(2), intPtr(5), nil,
		nil, nil, intPtr(5), nil, intPtr(2), intPtr(7), nil, nil, nil,
		nil, intPtr(5), nil, nil, nil, nil, nil, nil, intPtr(8),
		nil, intPtr(3), intPtr(7), nil, nil, nil, intPtr(4), nil, nil,
		nil, nil, nil, intPtr(3), nil, intPtr(4), nil, intPtr(6), nil,
		nil, intPtr(7), nil, nil, nil, nil, intPtr(8), nil, intPtr(2),
		intPtr(8), nil, nil, nil, nil, intPtr(3), intPtr(1), nil, nil,
		nil, nil, nil, intPtr(2), intPtr(1), nil, nil, intPtr(4), nil,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(board)

	s, err := solver.New([]strategies.Strategy{
		strategies.NewCleanupCandidates(),
		strategies.NewSingleCandidate(),
		strategies.NewSingleCellCandidate(),
	})
	if err != nil {
		panic(err)
	}

	err = s.Solve(board)
	if err != nil {
		panic(err)
	}
}

func intPtr(v uint8) *uint8 {
	return &v
}
