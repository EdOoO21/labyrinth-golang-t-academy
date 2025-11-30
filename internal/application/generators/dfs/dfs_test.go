package dfs

import (
	"math/rand"
	"testing"

	support "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/generators/support"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/models"
)

func TestDfsGenerate(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	dfs := &DFS{Rnd: &models.FakeRandom{R: rng}}

	tests := []struct {
		rows, cols int
	}{
		{1, 1},   // минимальный случай
		{2, 2},   // маленький квадрат
		{3, 3},   // маленький, но есть внутренняя клетка
		{4, 4},   // четный размер
		{5, 5},   // нечетный размер
		{10, 10}, // средний лабиринт
		{15, 20}, // прямоугольный лабиринт
		{20, 15}, // другой прямоугольник
		{50, 50}, // большой лабиринт
	}

	for _, tc := range tests {
		tc := tc
		t.Run("DFS_generator_check", func(t *testing.T) {
			maze := dfs.Generate(tc.rows, tc.cols)
			height, width := maze.GetHeight(), maze.GetWidth()
			if height != 2*tc.rows+1 ||
				width != 2*tc.cols+1 {
				t.Fatalf("dfs generated maze with incorrect shape: (%v, %v), expected: (%v, %v)",
					height, width, 2*tc.rows+1, 2*tc.cols+1)
			}
			if !support.CheckReachable(maze) {
				t.Fatal("dfs generated disconnected maze")
			}
		})
	}
}
