package domain

type Generator interface {
	Generate(rows, cols int) *Maze
}

type Solver interface {
	Solve(maze *Maze, start, end Cell) *Path
}

type CellType rune

const ( // символы использовать только из 1 руны,
	//  иначе надо менять логику в NewMazeFromReader
	CellWall  CellType = '#'
	CellEmpty CellType = ' '
	CellStart CellType = 'O'
	CellEnd   CellType = 'X'
	CellPath  CellType = '.'
)

type Maze struct {
	scheme        [][]CellType
	width, height int
}

type Path struct {
	path []Cell
}

type Cell struct {
	Row, Col int
}
