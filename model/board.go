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

// The model/data of the labyrinth
var Lab [][]Stone

var PlayerTurn bool = true

// MovingObj is a struct describing a moving object.
type MovingObj struct {
	// The position in the labyrinth in pixel coordinates
	Pos struct {
		X, Y float64
	}

	// Images for each direction, each has zero Min point
	Imgs []*image.RGBA
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
	LabImg = image.NewRGBA(image.Rect(0, 0, LabWidth, LabHeight))

	initLab()
	initLabImg()
}

// initLab initializes and generates a new Labyrinth.
func initLab() {
	fmt.Println("INIT LAB........ %v, %v", Rows, Cols);
	Lab = make([][]Stone, Rows)

	for i := range Lab {
		Lab[i] = make([]Stone, Cols)
	}
	genLab()
}


// initLabImg initializes and draws the image of the Labyrinth.
func initLabImg() {
	// Clear the labyrinth image
	fmt.Println("DRAWING........");
	draw.Draw(LabImg, LabImg.Bounds(), EmptyImg, image.Pt(0, 0), draw.Over)

	// Draw walls
	// zeroPt := image.Point{}
	for ri, row := range Lab {
		for ci, block := range row {
			if block == StoneEmpty {
				x, y := ci*BlockSize, ri*BlockSize
				// rect := image.Rect(x, y, x+BlockSize, y+BlockSize)
				// draw.Draw(LabImg, rect, WallImg, zeroPt, draw.Over)

				myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))
			    green := color.RGBA{0, 100, 0, 255}

			    // backfill entire surface with green
			    draw.Draw(LabImg, myimage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
			} else {
				x, y := ci*BlockSize, ri*BlockSize
				myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))
			    green := color.RGBA{200, 100, 200, 255}

			    // backfill entire surface with green
			    draw.Draw(LabImg, myimage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
			}
		}
	}
}

func DrawColRow(col int, row int) {
	fmt.Printf("setting:  %v, %v", col, row)
	for ri, row := range Lab {
		for ci, block := range row {
			x, y := ci*BlockSize, ri*BlockSize
			myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))

			pieceColor := color.RGBA{0, 100, 0, 255}
			if block == StoneBlack {
			    pieceColor = color.RGBA{0, 0, 0, 255}
			} else if block == StoneWhite {
			    pieceColor = color.RGBA{255, 255, 255, 255}
			} 
			draw.Draw(LabImg, myimage.Bounds(), &image.Uniform{pieceColor}, image.ZP, draw.Src)
		}
	}
}

// genLab generates a random labyrinth.
func genLab() {
	// Create a "frame":
	for ri := range Lab {
		Lab[ri][0] = StoneEmpty
	}
	for ci := range Lab[0] {
		Lab[0][ci] = StoneEmpty
	}
}
