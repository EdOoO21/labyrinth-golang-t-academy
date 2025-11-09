package domain

import "fmt"

// Maze methods ------------------------------------
func (m *Maze) GetCell(row, col int) (CellType, error) {
	if !m.InBorders(row, col) {
		return CellEmpty, fmt.Errorf("incorrect indexes row = %d, col = %d for maze (with fence) size (%d, %d)",
			row, col, m.height, m.width)
	}
	return m.scheme[row][col], nil
}

func (m *Maze) GetScheme() [][]CellType {
	return m.scheme
}

func (m *Maze) GetHeight() int {
	return m.height
}

func (m *Maze) GetWidth() int {
	return m.width
}

// Path methods ------------------------------------
func (p *Path) GetPath() []Cell {
	return p.path
}
