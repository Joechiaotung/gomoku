package ctrl

import (
	"github.com/skiptomyliu/gomoku/model"

	// "math/rand"
	"fmt"
)

func isOne(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count -= 1
		} 


		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count += 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	
	if count == 1 {
		return 10
	} else if count == 2 {
		return 20
	} else if count == 3 {
		return 40
	} else if count >= 4 {
		return 60
	} else {
		return 0
	}
}

func isTwo(i int, j int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k < 3; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count -= 1
		}  


		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count += 1
		} 
		
		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	
	if count >= 2 {
		return 20
	} 
	return 0
}


func isThree(i int, j int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k < 4; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count -= 1
		} 

		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count += 1
		} 
		
		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	
	if count >= 3 {
		return 30
	} 
	return 0
}


func isFour(i int, j int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k < 5; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count -= 1
		}  

		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count += 1
		} 
		
		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	
	if count >= 4 {
		return 40
	} 
	return 0
}


func NextMove() (bestMove model.PiecePos) {
	board := model.Board

	tmp := make([][]int, len(board))
	for i := range board {
		tmp[i] = make([]int, len(board))
	}
	// copy(tmp, board)
	maxCount := 0
	
	// maxScore, bMaxScore, wMaxScore := 0, 0, 0
	maxScore, wMaxScore := 0, 0
	bestMove = model.PiecePos{-1, -1}
	bestMoveScore := model.PiecePos{-1, -1}

	for i := range board {
		for j := range board {

			if board[i][j] == model.StoneEmpty {

				// fmt.Printf("@ %d,%d\n", i, j)
				// bCount, wCount, bScore, wScore := 0, 0, 0, 0
				// count := 0

				// score := isOne(i, j, 1, board)
				tmp[i][j] += isOne(i, j, 1, board)
				// tmp[i][j] += isOne(i, j, 2, board)
				// tmp[i][j] += isOne(i, j, 3, board)
				// tmp[i][j] += isOne(i, j, 4, board)
				// score := isOne(i, j, 2, board)
				// score := isOne(i, j, 3, board)
				// score := isOne(i, j, 4, board)
				// tmp[i][j] += score
				tmp[i][j] += isTwo(i, j, board)
				tmp[i][j] += isThree(i, j, board)
				tmp[i][j] += isFour(i, j, board)
				// if score > 0 {
				// 	fmt.Println(tmp[i][j])
				// }

				if tmp[i][j] > wMaxScore {
					wMaxScore = tmp[i][j]
					bestMoveScore = model.PiecePos{i, j}
				}


				// for k := 1; k < 5; k++ {

				// 	// BLACK STONE
				// 	// Check for column / x-axis					
				// 	if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
				// 		bCount++
				// 		bScore += (5-k)
				// 	}

				// 	if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
				// 		bCount++
				// 		bScore += (5-k)
				// 	}

				// 	// row / y-axis
				// 	if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
				// 		bCount++
				// 		bScore += (5-k)
				// 	}

				// 	if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
				// 		bCount++
				// 		bScore += (5-k)
				// 	}




				// 	//WHITE STONE
				// 	if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
				// 		wCount++
				// 		wScore += (5-k)

				// 	}

				// 	if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
				// 		wCount++
				// 		wScore += (5-k)
				// 	}

				// 	// row / y-axis
				// 	if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
				// 		wCount++
				// 		wScore += (5-k)
				// 	}

				// 	if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
				// 		wCount++
				// 		wScore += (5-k)
				// 	}

				// 	tmp[i][j] = bCount
				// 	tmp[i][j] = wScore

				// 	// defense
				// 	if wScore > wMaxScore && wMaxScore >= bMaxScore {
				// 		// bestMoveScore = model.PiecePos{i, j}
				// 		wMaxScore = wScore
				// 		// fmt.Println("WHITE MAX SCORE")
				// 		// fmt.Println(wMaxScore)
				// 		// bestMoveScore = model.PiecePos{i, j}
				// 		bestMoveScore = model.PiecePos{i, j}
				// 	}  
				// }

			}

			// if board[i][j] == model.StoneBlack {
			// 	fmt.Printf("black %d,%d\n", i, j)
			// }
		}

	}


	for i := range board {
		fmt.Println(tmp[i])
	} 

	// fmt.Println(tmp)
	// fmt.Println("BY MAX COUNT %v", bestMove)
	fmt.Printf("BY MAX SCORE %v - %v\n", bestMoveScore, maxScore)
	fmt.Printf("highest: %v - %v\n", bestMove, maxCount)
	if bestMove.X == -1 {
		return bestMoveScore
	}
	return bestMove
}
