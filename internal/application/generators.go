package application

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

// DFS ----------------------------------------------------------
func (d *DFS) Generate(rows, cols int) *domain.Maze {
	scheme := domain.NewMaze(rows, cols)
	stack := make([]domain.Cell, 0)
	stack = append(stack, domain.NewCell(2*d.rnd.GetRandom(rows)+1, 2*d.rnd.GetRandom(cols)+1)) // чтобы мы точно попали не в клетку забора
	_ = scheme.SetCell(stack[0].Row, stack[0].Col, domain.CellEmpty)
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		for _, val := range directions {
			if t, err := scheme.GetCell(p.Row+val[0]*2, p.Col+val[1]*2); err == nil && t == domain.CellWall {
				stack = append(stack, domain.NewCell(p.Row+val[0]*2, p.Col+val[1]*2))
				_ = scheme.SetCell(p.Row+val[0]*2, p.Col+val[1]*2, domain.CellEmpty)
				_ = scheme.SetCell(p.Row+val[0], p.Col+val[1], domain.CellEmpty)
			}
		}

	}

	return scheme
}

// Prim ----------------------------------------------------------
func (p *Prim) Generate(rows, cols int) *domain.Maze {
	scheme := domain.NewMaze(rows, cols)
	arr := make([]domain.Cell, 0)
	arr = append(arr, domain.NewCell(2*p.rnd.GetRandom(rows)+1, 2*p.rnd.GetRandom(cols)+1)) // чтобы мы точно попали не в клетку забора
	_ = scheme.SetCell(arr[0].Row, arr[0].Col, domain.CellEmpty)
	for len(arr) != 0 {
		ind := p.rnd.GetRandom(len(arr))
		p := arr[ind]
		arr = append(arr[:ind], arr[ind+1:]...)

		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		for _, val := range directions {
			if t, err := scheme.GetCell(p.Row+val[0]*2, p.Col+val[1]*2); err == nil && t == domain.CellWall {
				arr = append(arr, domain.NewCell(p.Row+val[0]*2, p.Col+val[1]*2))
				_ = scheme.SetCell(p.Row+val[0]*2, p.Col+val[1]*2, domain.CellEmpty)
				_ = scheme.SetCell(p.Row+val[0], p.Col+val[1], domain.CellEmpty)
			}
		}

	}

	return scheme
}
