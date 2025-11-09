package application

import (
	"container/heap"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func (d *Dijkstra) Solve(maze *domain.Maze, start, end domain.Cell) *domain.Path {
	pq := &PQDijkstra{}
	heap.Init(pq)

	prevs := make(map[domain.Cell]domain.Cell)
	path := domain.NewPath()

	if start == end {
		return path
	}

	dists := make(map[domain.Cell]int)
	dists[start] = 0
	visited := make(map[domain.Cell]struct{})
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT

	heap.Push(pq, pqItem{dist: 0, c: start})

	for pq.Len() != 0 {
		item := heap.Pop(pq).(pqItem)
		if _, ok := visited[item.c]; ok {
			continue
		}
		visited[item.c] = struct{}{}
		for _, way := range directions {
			p := domain.NewCell(item.c.Row+way[0]*2, item.c.Col+way[1]*2)

			if oldDist, ok := dists[p]; (maze.InBorders(p.Row, p.Col)) &&
				(!ok || oldDist > item.dist+1) {
				if v, _ := maze.GetCell(item.c.Row+way[0], item.c.Col+way[1]); v == domain.CellEmpty {
					dists[p] = item.dist + 1
					prevs[p] = item.c
					heap.Push(pq, pqItem{dist: item.dist + 1, c: p})
				}

			}
		}
	}

	now := end
	for now != start {
		path.AppendCell(now)
		path.AppendCell(domain.NewCell((now.Row+prevs[now].Row)/2,
			(now.Col+prevs[now].Col)/2))
		now = prevs[now]
	}
	path.AppendCell(now)
	path.Reverse()

	return path
}

func (a *AStar) Solve(maze *domain.Maze, start, end domain.Cell) *domain.Path {
	pq := &PQAStar{}
	heap.Init(pq)

	prevs := make(map[domain.Cell]domain.Cell)
	path := domain.NewPath()

	if start == end {
		return path
	}

	dists := make(map[domain.Cell]int)
	dists[start] = 0
	visited := make(map[domain.Cell]struct{})
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 1=UP, 2=DOWN, 3=LEFT, 4=RIGHT

	heap.Push(pq, astarItem{dist: 0, h: heuristic(start, end), c: start})

	for pq.Len() != 0 {
		item := heap.Pop(pq).(astarItem)
		if _, ok := visited[item.c]; ok {
			continue
		}
		visited[item.c] = struct{}{}
		for _, way := range directions {
			p := domain.NewCell(item.c.Row+way[0]*2, item.c.Col+way[1]*2)

			if oldDist, ok := dists[p]; (maze.InBorders(p.Row, p.Col)) &&
				(!ok || oldDist > item.dist+1) {
				if v, _ := maze.GetCell(item.c.Row+way[0], item.c.Col+way[1]); v == domain.CellEmpty {
					dists[p] = item.dist + 1
					prevs[p] = item.c
					heap.Push(pq, astarItem{dist: item.dist + 1, h: heuristic(p, end), c: p})
				}

			}
		}
	}

	now := end
	for now != start {
		path.AppendCell(now)
		path.AppendCell(domain.NewCell((now.Row+prevs[now].Row)/2,
			(now.Col+prevs[now].Col)/2))
		now = prevs[now]
	}
	path.AppendCell(now)
	path.Reverse()

	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(c, end domain.Cell) int {
	return abs(c.Row-end.Row) + abs(c.Col-end.Col)
}
