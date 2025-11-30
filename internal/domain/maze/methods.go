package maze

import (
	"fmt"

	cell "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

// setters ---------------------------------------------------------------------
func (m *Maze) SetCell(row, col int, cell cell.CellType) error {
	if !m.InBorders(row, col) {
		return fmt.Errorf("incorrect indexes row = %d, col = %d for maze size (%d, %d)",
			row, col, m.height, m.width)
	}
	m.scheme[row][col] = cell
	return nil
}

// boolean ---------------------------------------------------------------------
func (m *Maze) InBorders(row, col int) bool {
	return 0 <= row && row < m.height &&
		0 <= col && col < m.width
}

// getters ---------------------------------------------------------------------
func (m *Maze) GetCell(row, col int) (cell.CellType, error) {
	if !m.InBorders(row, col) {
		return cell.CellEmpty, fmt.Errorf("incorrect indexes row = %d, col = %d for maze (with fence) size (%d, %d)",
			row, col, m.height, m.width)
	}
	return m.scheme[row][col], nil
}

func (m *Maze) GetScheme() [][]cell.CellType {
	return m.scheme
}

func (m *Maze) GetHeight() int {
	return m.height
}

func (m *Maze) GetWidth() int {
	return m.width
}
