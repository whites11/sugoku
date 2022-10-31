package strategies

import "github.com/whites11/sugoku/pkg/types"

type Strategy interface {
	Step(b *types.Board) (*types.Cell, error)
}
