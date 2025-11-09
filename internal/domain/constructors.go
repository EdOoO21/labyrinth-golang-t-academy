package domain

import (
	"bufio"
	"fmt"
	"io"
)

func NewMaze(height, width int) *Maze {
	height = 2*height + 1
	width = 2*width + 1 // так надо делать, потому что между двумя стенками нам надо хранить инфу о заборе
	// то есть между каждым 2 клетками должно быть по 1 клетке показателя забора
	// но +1 еще так как (у нас эти клетки на четных индексах будут), получается, что в конце у нас не будет забора
	// тк конечный индекс - всегда нечетный. Закроем его просто для красоты вывода лабиринта
	scheme := make([][]CellType, height)
	for i := range height {
		scheme[i] = make([]CellType, width)
		for j := range width {
			scheme[i][j] = CellWall
		}
	}
	return &Maze{scheme: scheme,
		width: width, height: height}
}

func NewMazeFromReader(r io.Reader) (*Maze, error) {
	scheme := make([][]CellType, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		row := []CellType(line) // символы тут все однобайтовые, поэтому в данном случае:
		// 1 руна = 1 символ, 1 байт = 1 руна =>
		// 1 байт = 1 символ, получается что len(string) вернет ожидаемую длину
		if len(scheme) > 0 && len(row) != len(scheme[0]) {
			return nil, fmt.Errorf("file is broken: some rows are not equal by length")
		}
		scheme = append(scheme, row)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	if len(scheme) == 0 || len(scheme[0]) == 0 {
		return nil, fmt.Errorf("failed to fill scheme from file")
	}

	return &Maze{
		scheme: scheme,
		height: len(scheme),
		width:  len(scheme[0])}, nil
}

func NewMazeFromScheme(scheme [][]CellType) (*Maze, error) {
	if len(scheme) > 0 && len(scheme[0]) > 0 {
		maze := &Maze{scheme: scheme,
			height: len(scheme),
			width:  len(scheme[0])}
		return maze, nil
	}
	return nil, fmt.Errorf("empty scheme")
}

func NewPath() *Path {
	return &Path{path: make([]Cell, 0)}
}

func NewCell(row, col int) Cell {
	return Cell{Row: row, Col: col}
}
