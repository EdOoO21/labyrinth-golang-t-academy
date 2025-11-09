package domain

func (m *Maze) InBorders(row, col int) bool {
	return 0 <= row && row < m.height &&
		0 <= col && col < m.width
}
