package support

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
)

func CheckReachable(maze *maze.Maze) bool {
	height, width := maze.GetHeight(), maze.GetWidth()
	visited := make(map[cell.Cell]struct{})
	stack := []cell.Cell{cell.NewCell(1, 1)}
	count := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[p]; ok {
			continue
		}

		visited[p] = struct{}{}
		count += 1

		for _, way := range directions {
			r, c := p.Row+way[0], p.Col+way[1]

			if val, err := maze.GetCell(r, c); err == nil && val == cell.CellEmpty {
				p1 := cell.NewCell(r, c)
				if _, ok := visited[p1]; !ok {
					stack = append(stack, p1)
				}
			}

		}
	}

	for i := range height {
		for j := range width {
			if val, _ := maze.GetCell(i, j); val == cell.CellEmpty {
				count -= 1
			}
		}
	}

	return count == 0
}
