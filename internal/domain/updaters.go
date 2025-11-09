package domain

import "fmt"

// Maze methods ------------------------------------
func (m *Maze) SetCell(row, col int, cell CellType) error {
	if !m.InBorders(row, col) {
		return fmt.Errorf("incorrect indexes row = %d, col = %d for maze size (%d, %d)",
			row, col, m.height, m.width)
	}
	m.scheme[row][col] = cell
	return nil
}

// Path methods ------------------------------------
func (p *Path) AppendCell(c Cell) {
	p.path = append(p.path, c)
}

func (p *Path) Reverse() {
	for i, j := 0, len(p.path)-1; i < j; i, j = i+1, j-1 {
		p.path[i], p.path[j] = p.path[j], p.path[i]
	}
}
