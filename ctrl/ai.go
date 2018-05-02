package ctrl

import (
	"github.com/skiptomyliu/gomoku/model"

	// "math/rand"
	"fmt"
)



func NextMove() {
	board := model.Board

	tmp := make([][]int, len(board))
	for i := range board {
		tmp[i] = make([]int, len(board))
	}
	// copy(tmp, board)
	maxCount := 0
	bestMove := model.PiecePos{-1, -1}
	for i := range board {
		for j := range board {

			if board[i][j] == model.StoneEmpty {

				// fmt.Printf("@ %d,%d\n", i, j)
				count := 0
				for k := 1; k < 5; k++ {

					// Check for column / x-axis					
					if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
						// count++
						count += (5-k)
						// tmp[i][j] += 888
						// fmt.Printf("---pub %v,%v\n", i, j)
					}

					if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
						// count++
						count += (5-k)
						// tmp[i][j] += 888
						// fmt.Printf("---sub %v,%v\n", i, j)
					}

					// row / y-axis
					if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
						// count++
						count += (5-k)
					}

					if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
						// count++
						count += (5-k)
					}


					if count > maxCount {
						bestMove = model.PiecePos{i, j}
						maxCount = count
					}
					tmp[i][j] = count
				}

			}


			// if board[i][j] == model.StoneBlack {
			// 	fmt.Printf("black %d,%d\n", i, j)
			// }
		}
		// Lab[i] = make([]Stone, Cols)
	}
	fmt.Println(tmp)
	fmt.Printf("highest: %v\n", bestMove)
}

// ----
// -oo-
// ----

// ------
// -oooo-
// ------