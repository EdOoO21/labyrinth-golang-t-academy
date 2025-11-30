package cell

type Cell struct {
	Row, Col int
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
