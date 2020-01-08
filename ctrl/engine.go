package ctrl

import (
	// "image"
	// "image/draw"
	// "math"
	"fmt"
	"math/rand"
	"time"
)

// LoopDelay is the delay between the iterations of the main loop of the game engine, in milliseconds.
var LoopDelay = 20 // ~20 FPS

// InitNew initializes a new game.
func initNew() {
	// Initialize random number generator
	rand.Seed(time.Now().Unix())

	fmt.Println("Trying to init new")
	model.InitNew()
	view.InitNew()
}

// StartEngine starts the game engine in a new goroutine and returns as soon as possible.
func StartEngine() {
	model.NewGameCh <- 1 // Cannot block as application was just started, no incoming requests processed yet

	// initNew()
	model.Mutex.Lock()

	go simulate()
}

// Delta time since our last iteration
var dt float64

// simulate implements the game cycle
func simulate() {

	for {
		// Check if we have to start a new game
		select {
		case <-model.NewGameCh:
			initNew()
		default:
		}

		// Process mouse clicks
	clickLoop:
		for {
			select {
			case click := <-model.ClickCh:
				handleClick(click)
			default:
				break clickLoop
			}
		}

		// now := time.Now().UnixNano()
		// dt = float64(now-t) / 1e9
		// t = now

		// Sleep some time.
		// Iterations might not be exact, but we don't rely on it:
		// We calculate delta time and calculate moving and next positions
		// based on the delta time.

		model.Mutex.Unlock() // While sleeping, clients can request view images
		if model.Won {
			// If won, nothing has to be done, just wait for a new game signal
			<-model.NewGameCh // Blocking receive
			// Send back value to detect it at the proper place
			model.NewGameCh <- 1
		}
		time.Sleep(time.Millisecond * time.Duration(LoopDelay))
		model.Mutex.Lock() // We will modify model now, labyrinth image might change so lock.
	}
}

// handleClick handles a mouse click
func handleClick(c model.Click) {

	col, row := c.X/model.BlockSize, c.Y/model.BlockSize

	// AI
	model.PlayerTurn = true
	if model.PlayerTurn == true {
		if model.Board[row][col] != model.StoneWhite && model.Board[row][col] != model.StoneBlack {
			model.Board[row][col] = model.StoneWhite

			aiMove := NextMove()
			fmt.Println(aiMove)
			model.Board[aiMove.X][aiMove.Y] = model.StoneBlack

			model.PlayerTurn = !model.PlayerTurn

			ai_flag := handleWinning(aiMove.X, aiMove.Y)
			if ai_flag == 1 {
				fmt.Println("AI win")
				//fmt.Println("aimove", aiMove.X, aiMove.Y)
				model.Won = true
			}

			model.DrawColaRow(col, row)

			flag := handleWinning(row, col)
			if flag == 1 {
				fmt.Println("client win")
				//fmt.Println("col", col, row)
				model.Won = true
			}
		}
	} else {
		// aiMove := NextMove()
		// fmt.Println(aiMove)
		// model.Board[aiMove.X][aiMove.Y] = model.StoneBlack
		// model.Board[row][col] = model.StoneBlack
	}

	// model.Board[aiMove.X][aiMove.Y] = model.StoneBlack

	// model.PlayerTurn = false

}

// handleWinning handles the winning of game event.
func handleWinning(x int, y int) int {

	//var a,b int
	//color := model.Board[x][y]
	flag := 0
	/*
		if(stone[x][y]==1)
			wcscpy_s(win,_T("黑棋勝利！"));
		else
			wcscpy_s(win,_T("白棋勝利！"));
	*/
	for a := x - 4; a <= x+4; a++ {
		//判斷橫
		if model.Board[a][y] == model.Board[x][y] && model.Board[a+1][y] == model.Board[x][y] && model.Board[a+2][y] == model.Board[x][y] && model.Board[a+3][y] == model.Board[x][y] && model.Board[a+4][y] == model.Board[x][y] {
			//fmt.Println("row")
			flag = 1
			break
		}
	}
	for b := y - 4; b <= y+4; b++ { //判斷豎
		if (model.Board[x][b] == model.Board[x][y]) && (model.Board[x][b+1] == model.Board[x][y]) && (model.Board[x][b+2] == model.Board[x][y]) && (model.Board[x][b+3] == model.Board[x][y]) && (model.Board[x][b+4] == model.Board[x][y]) {
			//fmt.Println("column")
			flag = 1
			break
		}
	}
	for a, b := x-4, y-4; a <= x+4; a, b = a+1, b+1 { //判斷右斜
		if model.Board[a][b] == model.Board[x][y] && model.Board[a+1][b+1] == model.Board[x][y] && model.Board[a+2][b+2] == model.Board[x][y] && model.Board[a+3][b+3] == model.Board[x][y] && model.Board[a+4][b+4] == model.Board[x][y] {
			//fmt.Println("+x +y")
			flag = 1
			break
		}
	}
	for a, b := x-4, y+4; a <= x+4; a, b = a+1, b-1 { //判斷左斜
		//fmt.Println(a, b)
		if model.Board[a][b] == model.Board[x][y] && model.Board[a+1][b-1] == model.Board[x][y] && model.Board[a+2][b-2] == model.Board[x][y] && model.Board[a+3][b-3] == model.Board[x][y] && model.Board[a+4][b-4] == model.Board[x][y] {
			//fmt.Println("+x -y")
			flag = 1
			break
		}
	}
	if flag == 1 {

		return 1

	}
	return 0

}

func init() {

}
