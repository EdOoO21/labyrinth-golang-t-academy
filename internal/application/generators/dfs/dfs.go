package dfs

import (
	"math/rand"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

// DFS ----------------------------------------------------------
func (d *DFS) Generate(rows, cols int) *maze.Maze {
	scheme := maze.NewMaze(rows, cols)
	stack := make([]cell.Cell, 0)
	stack = append(stack, cell.NewCell(2*d.Rnd.GetRandom(rows)+1, 2*d.Rnd.GetRandom(cols)+1)) // чтобы мы точно попали не в клетку забора
	_ = scheme.SetCell(stack[0].Row, stack[0].Col, cell.CellEmpty)
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		for _, val := range directions {
			if t, err := scheme.GetCell(p.Row+val[0]*2, p.Col+val[1]*2); err == nil && t == cell.CellWall {
				stack = append(stack, cell.NewCell(p.Row+val[0]*2, p.Col+val[1]*2))
				_ = scheme.SetCell(p.Row+val[0]*2, p.Col+val[1]*2, cell.CellEmpty)
				_ = scheme.SetCell(p.Row+val[0], p.Col+val[1], cell.CellEmpty)
			}
		}

	}

	return scheme
}
