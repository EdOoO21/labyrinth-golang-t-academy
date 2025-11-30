package dijkstra

import (
	"container/heap"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/solvers/support"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/maze"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/path"
)

func (d *Dijkstra) Solve(maze *maze.Maze, start, end cell.Cell) *path.Path {
	pq := &support.PQDijkstra{}
	heap.Init(pq)

	prevs := make(map[cell.Cell]cell.Cell)
	path := path.NewPath()

	if start == end {
		return path
	}

	dists := make(map[cell.Cell]int)
	dists[start] = 0
	visited := make(map[cell.Cell]struct{})
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT

	heap.Push(pq, support.PqItem{Dist: 0, C: start})

	for pq.Len() != 0 {
		item := heap.Pop(pq).(support.PqItem)
		if _, ok := visited[item.C]; ok {
			continue
		}
		visited[item.C] = struct{}{}
		for _, way := range directions {
			p := cell.NewCell(item.C.Row+way[0]*2, item.C.Col+way[1]*2)

			if oldDist, ok := dists[p]; (maze.InBorders(p.Row, p.Col)) &&
				(!ok || oldDist > item.Dist+1) {
				if v, _ := maze.GetCell(item.C.Row+way[0], item.C.Col+way[1]); v == cell.CellEmpty {
					dists[p] = item.Dist + 1
					prevs[p] = item.C
					heap.Push(pq, support.PqItem{Dist: item.Dist + 1, C: p})
				}

			}
		}
	}

	now := end
	for now != start {
		path.AppendCell(now)
		path.AppendCell(cell.NewCell((now.Row+prevs[now].Row)/2,
			(now.Col+prevs[now].Col)/2))
		now = prevs[now]
	}
	path.AppendCell(now)
	path.Reverse()

	return path
}
