package ctrl

import (
	"github.com/skiptomyliu/gomoku/model"

	// "math/rand"
	"fmt"
)



func NextMove() (bestMove model.PiecePos) {
	board := model.Board

	tmp := make([][]int, len(board))
	for i := range board {
		tmp[i] = make([]int, len(board))
	}
	// copy(tmp, board)
	maxCount := 0
	
	maxScore, bMaxScore, wMaxScore := 0, 0, 0
	bestMove = model.PiecePos{-1, -1}
	bestMoveScore := model.PiecePos{-1, -1}

	for i := range board {
		for j := range board {

			if board[i][j] == model.StoneEmpty {

				// fmt.Printf("@ %d,%d\n", i, j)
				bCount, wCount, bScore, wScore := 0, 0, 0, 0
				for k := 1; k < 5; k++ {


					// BLACK STONE
					// Check for column / x-axis					
					if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
						bCount++
						bScore += (5-k)
					}

					if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
						bCount++
						bScore += (5-k)
					}

					// row / y-axis
					if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
						bCount++
						bScore += (5-k)
					}

					if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
						bCount++
						bScore += (5-k)
					}




					//WHITE STONE
					if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
						wCount++
						wScore += (5-k)
					}

					if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
						wCount++
						wScore += (5-k)
					}

					// row / y-axis
					if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
						wCount++
						wScore += (5-k)
					}

					if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
						wCount++
						wScore += (5-k)
					}


					tmp[i][j] = bCount


					//block win					
					if wCount >= 4 {
						fmt.Println(tmp)
						fmt.Println("BLOCKING the win... %v", bestMove)
						bestMove = model.PiecePos{i, j}
						return bestMove
					}

					// win
					if bCount >= 4 {
						fmt.Println(tmp)
						fmt.Println("RETURNING the win... %v", bestMove)
						bestMove = model.PiecePos{i, j}
						return bestMove
					}



					// by count
					// if bCount > maxCount {
					// 	bestMove = model.PiecePos{i, j}
					// 	maxCount = bCount
					// }

					// by score
					if bScore > bMaxScore {
						
						bMaxScore = bScore
						// fmt.Println("BLACK MAX SCORE")
						// fmt.Println(bMaxScore)
					}

					if wScore > wMaxScore {
						// bestMoveScore = model.PiecePos{i, j}
						wMaxScore = wScore
						// fmt.Println("WHITE MAX SCORE")
						// fmt.Println(wMaxScore)
						// bestMoveScore = model.PiecePos{i, j}
						bestMoveScore = model.PiecePos{i, j}
					}

					if wMaxScore >= bMaxScore {
						maxScore = wMaxScore
						// bestMoveScore = model.PiecePos{i, j}
						// fmt.Println("max wscore.....")
						// fmt.Println(maxScore)
						// bestMoveScore = wMaxScore
					} else {
						fmt.Println("max score.....")
						fmt.Println(maxScore)
						// maxScore = bMaxScore
						// bestMoveScore = bMaxScore
						// bestMoveScore = model.PiecePos{i, j}
					}



				}

			}

			// if board[i][j] == model.StoneBlack {
			// 	fmt.Printf("black %d,%d\n", i, j)
			// }
		}

	}
	fmt.Println(tmp)
	// fmt.Println("BY MAX COUNT %v", bestMove)
	fmt.Printf("BY MAX SCORE %v - %v\n", bestMoveScore, maxScore)
	fmt.Printf("highest: %v - %v\n", bestMove, maxCount)
	if bestMove.X == -1 {
		return bestMoveScore
	}
	return bestMove
}
