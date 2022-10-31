package types

import (
	"errors"
	"fmt"
)

type Board struct {
	solved uint8
	cells  []Cell
}

func NewBoard(values []*uint8) (*Board, error) {
	if len(values) != 81 {
		return nil, errors.New("Data must have 81 entries")
	}

	data := make([]Cell, 0)
	allCandidates := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	solved := uint8(0)
	for idx, val := range values {
		row := uint8(idx / 9)
		column := uint8(idx % 9)

		squareRow := row / 3
		squareColumn := column / 3
		square := squareRow*3 + squareColumn

		cellCandidates := allCandidates
		if val != nil {
			cellCandidates = nil
		}

		data = append(data, Cell{
			Value:      val,
			Row:        row,
			Column:     column,
			Square:     square,
			Candidates: cellCandidates,
		})
		if val != nil {
			solved = solved + 1
		}
	}

	b := Board{
		solved: solved,
		cells:  data,
	}

	return &b, nil
}

func (b *Board) Cells() []Cell {
	return b.cells
}

func (b *Board) Solved() bool {
	return b.solved == 81
}

func (b *Board) UpdateCell(cell *Cell) {
	if b.cells[cell.Row*9+cell.Column].Value == nil && cell.Value != nil {
		b.solved = b.solved + 1
	}
	b.cells[cell.Row*9+cell.Column].Value = cell.Value
	candidates := cell.Candidates
	if cell.Value != nil {
		candidates = nil
	}
	b.cells[cell.Row*9+cell.Column].Candidates = candidates
}

func (b *Board) Rows() []Row {
	data := make([]Row, 0)
	var i uint8
	for i < 9 {
		row, _ := b.Row(i)
		data = append(data, *row)
		i = i + 1
	}

	return data
}

func (b *Board) Row(index uint8) (*Row, error) {
	if index > 8 {
		return nil, errors.New("Index out of bounds")
	}

	return &Row{
		index: index + 1,
		data:  b.cells[index*9 : (index*9)+9],
	}, nil
}

func (b *Board) Columns() []Column {
	data := make([]Column, 0)
	var i uint8
	for i < 9 {
		col, _ := b.Column(i)
		data = append(data, *col)
		i = i + 1
	}

	return data
}

func (b *Board) Column(index uint8) (*Column, error) {
	if index > 8 {
		return nil, errors.New("Index out of bounds")
	}

	data := make([]Cell, 0)
	row := uint8(0)
	for row < 9 {
		data = append(data, b.cells[row*9+index])
		row += 1
	}

	return &Column{
		index: index,
		data:  data,
	}, nil
}

func (b *Board) Squares() []Square {
	data := make([]Square, 0)
	var i uint8
	for i < 9 {
		sq, _ := b.Square(i)
		data = append(data, *sq)
		i = i + 1
	}

	return data
}

func (b *Board) Square(index uint8) (*Square, error) {
	if index > 8 {
		return nil, errors.New("Index out of bounds")
	}

	firstRow := index / 3
	firstCol := (index % 3) * 3

	offset := firstRow*27 + firstCol

	data := make([]Cell, 0)
	row := uint8(0)
	for row < 3 {
		start := offset + row*9
		data = append(data, b.cells[start:start+3]...)
		row = row + 1
	}

	return &Square{
		index: index,
		data:  data,
	}, nil
}

func (b *Board) String() string {
	ret := ""
	for _, row := range b.Rows() {
		ret = ret + row.String() + "\n"
	}

	return fmt.Sprintf("Solved: %d\n%s\n", b.solved, ret)
}
