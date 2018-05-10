package model

import (
	"image"
	"image/color"
	"fmt"
	"image/draw"
	// "math/rand"
	"sync"
)

// Mutex to be used to synchronize model modifications
var Mutex sync.Mutex

// The model/data of the board
var Board [][]Stone

var PlayerTurn bool = true

type PiecePos struct {
	// The position in the board in pixel coordinates
	// Pos struct {
		X, Y int
	// }
	// Imgs []*image.RGBA
}

// Tells if we won
var Won bool

// Channel to signal new game
var NewGameCh = make(chan int, 1)

// Constant for the right Mouse button value in the Click struct.
// Button value for left and middle may not be the same for older browsers, but right button always has this value.
const MouseBtnRight = 2

// Click describes a mouse click.
type Click struct {
	// X, Y are the mouse coordinates in pixel, in the coordinate system of the Labyrinth
	X, Y int
	// Btn is the mouse button
	Btn int
}

// Channel to receive mouse clicks on (view package sends, ctrl package (engine) processes them)
var ClickCh = make(chan Click, 10)

// InitNew initializes a new game.
func InitNew() {
	fmt.Println("BOARD INIT NEW........");
	BoardImg = image.NewRGBA(image.Rect(0, 0, BoardWidth, BoardHeight))
    bg := color.RGBA{100, 90, 100, 255}
    // backfill entire surface with green
    draw.Draw(BoardImg, BoardImg.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	initBoard()
	initBoardImg()
}

// initBoard initializes and generates a new Board.
func initBoard() {
	fmt.Println("INIT BOARD........ %v, %v", Rows, Cols);
	Board = make([][]Stone, Rows)

	for i := range Board {
		Board[i] = make([]Stone, Cols)
	}
	genBoard()
}


// initBoardImg initializes and draws the image of the Labyrinth.
func initBoardImg() {

	fmt.Println("DRAWING........");

	for ri, row := range Board {
		for ci, block := range row {
			if block == StoneEmpty {
				x, y := ci*BlockSize, ri*BlockSize

				myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))
			    green := color.RGBA{0, 100, 0, 255}

			    // backfill entire surface with green
			    draw.Draw(BoardImg, myimage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
			} 
		}
	}
}

func DrawColRow(col int, row int) {
	for ri, row := range Board {
		for ci, block := range row {
			x, y := ci*BlockSize, ri*BlockSize
			myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))

			pieceColor := color.RGBA{0, 100, 0, 255}
			if block == StoneBlack {
			    pieceColor = color.RGBA{50, 10, 10, 255}
			} else if block == StoneWhite {
			    pieceColor = color.RGBA{255, 255, 255, 255}
			} 
			draw.Draw(BoardImg, myimage.Bounds(), &image.Uniform{pieceColor}, image.ZP, draw.Src)
		}
	}
}

// genLab generates a random labyrinth.
func genBoard() {
	// Create a "frame":
	for ri := range Board {
		Board[ri][0] = StoneEmpty
	}
	for ci := range Board[0] {
		Board[0][ci] = StoneEmpty
	}
}
