package main

import (
	"fmt"

	"github.com/whites11/sugoku/pkg/solver"
	"github.com/whites11/sugoku/pkg/strategies"
	"github.com/whites11/sugoku/pkg/types"
)

func main() {
	board, err := types.NewBoard([]*uint8{
		intPtr(7), intPtr(5), nil, nil, nil, intPtr(1), nil, intPtr(4), intPtr(9),
		nil, nil, intPtr(9), intPtr(6), intPtr(7), intPtr(2), nil, intPtr(3), intPtr(1),
		nil, nil, intPtr(6), nil, intPtr(4), nil, intPtr(8), nil, nil,
		nil, intPtr(3), intPtr(8), nil, intPtr(5), nil, nil, nil, intPtr(6),
		intPtr(5), nil, nil, nil, nil, nil, nil, nil, nil,
		intPtr(6), nil, intPtr(1), intPtr(2), nil, nil, intPtr(7), nil, intPtr(3),
		intPtr(8), nil, nil, nil, nil, intPtr(5), nil, intPtr(6), nil,
		nil, nil, intPtr(4), intPtr(7), intPtr(6), intPtr(3), intPtr(1), intPtr(8), intPtr(5),
		nil, nil, nil, intPtr(9), nil, nil, intPtr(3), intPtr(7), intPtr(4),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(board)

	s, err := solver.New([]strategies.Strategy{
		strategies.NewCleanupCandidates(),
		strategies.NewSingleCandidate(),
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
