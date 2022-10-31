package solver

import (
	"errors"
	"fmt"

	"github.com/whites11/sugoku/pkg/strategies"
	"github.com/whites11/sugoku/pkg/types"
)

type Solver struct {
	strategies []strategies.Strategy
}

func New(strategies []strategies.Strategy) (*Solver, error) {
	if len(strategies) == 0 {
		return nil, errors.New("Need at least one strategy")
	}
	return &Solver{
		strategies: strategies,
	}, nil
}

func (s *Solver) Solve(b *types.Board) error {
	for !b.Solved() {
		anyStrategyWorked := false
		for _, strategy := range s.strategies {
			cell, err := strategy.Step(b)
			if err != nil {
				return err
			}

			if cell != nil {
				b.UpdateCell(cell)
				anyStrategyWorked = true

				fmt.Println(b)
				break
			}
		}

		if !anyStrategyWorked {
			return errors.New("Sorry I'm unable to solve this board")
		}
	}

	return nil
}
