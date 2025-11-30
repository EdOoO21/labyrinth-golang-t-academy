package path

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"

func NewPath() *Path {
	return &Path{path: make([]cell.Cell, 0)}
}

func (p *Path) AppendCell(c cell.Cell) {
	p.path = append(p.path, c)
}

func (p *Path) Reverse() {
	for i, j := 0, len(p.path)-1; i < j; i, j = i+1, j-1 {
		p.path[i], p.path[j] = p.path[j], p.path[i]
	}
}

func (p *Path) GetPath() []cell.Cell {
	return p.path
}
