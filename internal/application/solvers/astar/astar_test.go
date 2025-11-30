package astar

import (
	"math/rand"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/generators/prim"
	mods "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/models"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/solvers/support"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain/cell"
)

func TestAStarSolver(t *testing.T) {
	dfs := &prim.Prim{Rnd: &mods.FakeRandom{R: rand.New(rand.NewSource(42))}}

	AStar := &AStar{}
	tests := []struct {
		rows, cols     int
		startX, startY int
		endX, endY     int
	}{
		{1, 1, 1, 1, 1, 1},     // минимальный случай
		{2, 2, 1, 1, 3, 3},     // маленький квадрат
		{3, 3, 1, 1, 5, 5},     // маленький, есть внутренняя клетка
		{4, 4, 1, 1, 7, 7},     // четный размер
		{5, 5, 1, 1, 9, 9},     // нечетный размер
		{10, 10, 1, 1, 19, 19}, // средний лабиринт
		{15, 20, 1, 1, 29, 39}, // прямоугольный лабиринт
		{20, 15, 1, 1, 39, 29}, // другой прямоугольник
		{50, 50, 1, 1, 99, 99}, // большой лабиринт
	}
	// тут значения клеток уже домножены, потому что клетки
	// еще на стадии парсинга масштабируются, а тут вызов без парсинга
	// в общем понятно

	for _, tc := range tests {
		tc := tc
		t.Run("AStar_solver_check", func(t *testing.T) {

			maze := dfs.Generate(tc.rows, tc.cols)
			mazeScheme := maze.GetScheme()

			start := cell.NewCell(tc.startX, tc.startY)
			end := cell.NewCell(tc.endX, tc.endY)

			path := AStar.Solve(maze, start, end)

			realPath := path.GetPath()
			if len(realPath) == 0 {
				if start != end {
					t.Fatalf("empty path but start!=end")
				}
				return
			}

			if realPath[0] != start || realPath[len(realPath)-1] != end {
				t.Fatalf("incorrect path endpoints: got start=%v end=%v, want start=%v end=%v",
					realPath[0], realPath[len(realPath)-1], start, end)
			}

			if support.CheckPath(mazeScheme, maze, realPath) {
				t.Fatalf("incorrect path: some walls were broken or steps are invalid")

			}
		})
	}
}
