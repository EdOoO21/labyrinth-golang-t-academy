package application

import (
	"math/rand"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type fakeRandom struct {
	r *rand.Rand
}

func (s *fakeRandom) GetRandom(n int) int {
	if n <= 0 {
		return 0
	}
	return s.r.Intn(n)
} // пришлось написать сюда еще рандом,
// чтобы не импортировать его из inf, если
// есть другие решения, можно в фидбек
// было бы интересно почитать

func TestDfsGenerate(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	dfs := &DFS{rnd: &fakeRandom{r: rng}}

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
			if !checkReachable(maze) {
				t.Fatal("dfs generated disconnected maze")
			}
		})
	}
}

func TestPrimGenerate(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	prim := &Prim{rnd: &fakeRandom{r: rng}}

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
		t.Run("PRIM_generator_check", func(t *testing.T) {
			maze := prim.Generate(tc.rows, tc.cols)
			height, width := maze.GetHeight(), maze.GetWidth()
			if height != 2*tc.rows+1 ||
				width != 2*tc.cols+1 {
				t.Fatalf("prim generated maze with incorrect shape: (%v, %v), expected: (%v, %v)",
					height, width, 2*tc.rows+1, 2*tc.cols+1)
			}
			if !checkReachable(maze) {
				t.Fatal("prim generated disconnected maze")
			}
		})
	}
}

func checkReachable(maze *domain.Maze) bool {
	height, width := maze.GetHeight(), maze.GetWidth()
	visited := make(map[domain.Cell]struct{})
	stack := []domain.Cell{domain.NewCell(1, 1)}
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

			if val, err := maze.GetCell(r, c); err == nil && val == domain.CellEmpty {
				p1 := domain.NewCell(r, c)
				if _, ok := visited[p1]; !ok {
					stack = append(stack, p1)
				}
			}

		}
	}

	for i := range height {
		for j := range width {
			if val, _ := maze.GetCell(i, j); val == domain.CellEmpty {
				count -= 1
			}
		}
	}

	return count == 0
}
