package prim

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// Prim ----------------------------------------------------------
func (p *Prim) Generate(rows, cols int) *maze.Maze {
	scheme := maze.NewMaze(rows, cols)
	arr := make([]cell.Cell, 0)
	arr = append(arr, cell.NewCell(2*p.Rnd.GetRandom(rows)+1, 2*p.Rnd.GetRandom(cols)+1)) // чтобы мы точно попали не в клетку забора
	_ = scheme.SetCell(arr[0].Row, arr[0].Col, cell.CellEmpty)
	for len(arr) != 0 {
		ind := p.Rnd.GetRandom(len(arr))
		p := arr[ind]
		arr = append(arr[:ind], arr[ind+1:]...)

		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		for _, val := range directions {
			if t, err := scheme.GetCell(p.Row+val[0]*2, p.Col+val[1]*2); err == nil && t == cell.CellWall {
				arr = append(arr, cell.NewCell(p.Row+val[0]*2, p.Col+val[1]*2))
				_ = scheme.SetCell(p.Row+val[0]*2, p.Col+val[1]*2, cell.CellEmpty)
				_ = scheme.SetCell(p.Row+val[0], p.Col+val[1], cell.CellEmpty)
			}
		}

	}

	return scheme
}
