package ctrl

import (
	"github.com/skiptomyliu/gomoku/model"

	// "math/rand"
	"fmt"
)


func defenseScore(count int, targetCount int)(int) {
	if count >= targetCount {
		if targetCount == 1 {
			return 10
		} else if targetCount == 2 {
			return 20
		} else if targetCount == 3 {
			return 100
		} else if targetCount == 4 {
			return 100000
		}
	}
	return 0
}


func offenseScore(count int, targetCount int)(int) {
	if count >= targetCount {
		if targetCount == 1 {
			return 15
		} else if targetCount == 2 {
			return 30
		} else if targetCount == 3 {
			return 2000
		} else if targetCount == 4 {
			return 10000000
		}
	}
	return 0
}


func horizontalDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count0 -= 1
		} 

		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count0 += 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count1 -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count1 += 1
		}
	}

	if count0 > count1 {
		return defenseScore(count0, targetCount)		
	} else {
		return defenseScore(count1, targetCount)
	}

}

func diagonalDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (i < len(board)-k) && (j < len(board)-k) && (board[i+k][j+k] == model.StoneBlack) {
			count0 -= 1
		} 

		if (i < len(board)-k) && j < (len(board)-k) && (board[i+k][j+k] == model.StoneWhite) {
			count0 += 1
		} 

		if (i-k > 0) && (j-k > 0) && (board[i-k][j-k] == model.StoneBlack) {
			count1 -= 1
		} 

		if (i-k > 0) && (j-k > 0) && (board[i-k][j-k] == model.StoneWhite) {
			count1 += 1
		}	
	}
	
	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}
}


func diagonalDefense2(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneBlack) {
			count0 -= 1
		} 

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneWhite) {
			count0 += 1
		} 

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneBlack) {
			count1 -= 1
		} 

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneWhite) {
			count1 += 1
		}
	}
	
	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}
}


func verticalDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
			count0 -= 1
		} 

		if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
			count1 -= 1
		}

		if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
			count0 += 1
		} 

		if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
			count1 += 1
		}

	}

	if count0 > count1 {
		return defenseScore(count0, targetCount)		
	} else {
		return defenseScore(count1, targetCount)		
	}
	
}

/*
-0
0-
-00
0-0
00-
-000
0-00
00-0
000-

-0000
0-000
00-00
000-0
0000-

*/
func voWin(i int, j int, board [][]model.Stone)(int) {
	count, maxCount := 0, 0
	// fmt.Printf("%v,%v -- ", i, j)
	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}
			if (i < len(board)-k-m) && (i+k-m > 0) && (board[i+k-m][j] == model.StoneBlack) {
				count += 1
			}

			if (i < len(board)-k-m) && (i+k-m > 0) && (board[i+k-m][j] == model.StoneWhite) {
				return 0
			}
			// fmt.Printf("(%v,%v), ", i+k-m, j)
		}
		if count > maxCount {
			maxCount = count
		}
		// fmt.Printf("\n")
	}
	if maxCount > 1 {
		fmt.Printf("OFFENSE VERT SCORE COUNT: (%v,%v) - %v\n", i, j, maxCount)		
	}

	return offenseScore(maxCount, 0)
}

func hoWin(i int, j int, board [][]model.Stone)(int) {
	count, maxCount := 0, 0
	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}
			if (j < len(board)-k-m) && (j+k-m > 0) && (board[i][j+k-m] == model.StoneBlack) {
				count += 1
			}

			if (j < len(board)-k-m) && (j+k-m > 0) && (board[i][j+k-m] == model.StoneWhite) {
				return 0
			}
			// fmt.Printf("(%v,%v), ", i, j+k-m)
		}
		if count > maxCount {
			maxCount = count
		}
		// fmt.Printf("\n")
	}
	if maxCount > 1 {
		fmt.Printf("OFFENSE HORIZ SCORE COUNT: (%v,%v) - %v | score: %v\n", i, j, maxCount, offenseScore(maxCount, maxCount))		
	}

	return offenseScore(maxCount, maxCount)
}


func horizontalOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count -= 1
		} 

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count += 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count -= 1
		} 

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count += 1
		}
	}

	return offenseScore(count, targetCount)		
}

func verticalOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
			count -= 1
		} 

		if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
			count += 1
		} 

		if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
			count -= 1
		} 

		if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
			count += 1
		}
	}
	

	return offenseScore(count, targetCount)		
}


func diagonalOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (i < len(board)-k) && (j < len(board)-k) && (board[i+k][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (i < len(board)-k) && j < (len(board)-k) && (board[i+k][j+k] == model.StoneWhite) {
			count += 1
		}

		if (i-k > 0) && (j-k > 0) && (board[i-k][j-k] == model.StoneBlack) {
			count -= 1
		}

		if (i-k > 0) && (j-k > 0) && (board[i-k][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	

	return offenseScore(count, targetCount)		

}


func diagonalOffense2(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneBlack) {
			count -= 1
		} 

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneWhite) {
			count += 1
		} 

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneBlack) {
			count -= 1
		} 

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneWhite) {
			count += 1
		}
	}
	
	return offenseScore(count, targetCount)		

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


				// isWin(i, j, 0, board)

				for k := 1; k < 5; k++ {

					

					tmp[i][j] += horizontalDefense(i, j, k, board)					
					tmp[i][j] += verticalDefense(i, j, k, board)					
					tmp[i][j] += diagonalDefense(i, j, k, board)
					tmp[i][j] += diagonalDefense2(i, j, k, board)

					// tmp[i][j] += horizontalOffense(i, j, k, board)
					// tmp[i][j] += verticalOffense(i, j, k, board)
					tmp[i][j] += diagonalOffense(i, j, k, board)
					tmp[i][j] += diagonalOffense2(i, j, k, board)

				}

				tmp[i][j] += hoWin(i, j, board)
				tmp[i][j] += voWin(i, j, board)

				if tmp[i][j] > wMaxScore {
					wMaxScore = tmp[i][j]
					bestMoveScore = model.PiecePos{i, j}
				}


			}

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
