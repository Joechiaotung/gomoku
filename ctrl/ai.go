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

	for i := range board {
		for j := range board {

			if board[i][j] == model.StoneEmpty {

				// fmt.Printf("@ %d,%d\n", i, j)


				// Check for column / x-axis
				count := 0
				for k := 1; k < 5; k++ {
					if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
						count++
						// tmp[i][j] += 888
						// fmt.Printf("---pub %v,%v\n", i, j)
					}

					if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
						count++
						// tmp[i][j] += 888
						// fmt.Printf("---sub %v,%v\n", i, j)
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
}

// ----
// -oo-
// ----

// ------
// -oooo-
// ------