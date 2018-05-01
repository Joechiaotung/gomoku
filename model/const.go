package model

const (
	// BlockSize is the size of the board unit in pixels.
	BlockSize = 40
)

var (
	// Rows is the number of rows in the Labyrinth
	Rows int
	// Cols is the number of columns in the Labyrinth
	Cols int

	// LabWidth is the width of the labyrinth's image in pixels.
	LabWidth int

	// LabHeight is the height of the labyrinth's image in pixels.
	LabHeight int
)

// V is the moving speed of Gopher and the Buddlogs in pixel/sec.
var V float64

// Type of the unit of the labyrinth
type Stone int

// Pieces
const (
	StoneEmpty Stone = iota
	StoneWhite
	StoneBlack
)

type Dir int

// Directions of Gopher (facing directions)
const (
	DirRight Dir = iota
	DirLeft
	DirUp
	DirDown

	// Not a valid direction: just to tell how many directions are there
	DirLength
)

func (d Dir) String() string {
	switch d {
	case DirRight:
		return "right"
	case DirLeft:
		return "left"
	case DirUp:
		return "up"
	case DirDown:
		return "down"
	}
	return ""
}