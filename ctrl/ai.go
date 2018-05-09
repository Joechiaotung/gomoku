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
			return 50
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
			return 500
		} else if targetCount == 4 {
			return 10000000
		}
	}
	return 0
}


func hHeatDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (j < len(board)-k) {
			if (board[i][j+k] == model.StoneBlack) {
				count0 -= 1
			} 

			if (board[i][j+k] == model.StoneWhite) {
				count0 += 1
			} 
		}

		if (j-k > 0) {
			if (board[i][j-k] == model.StoneBlack) {
				count1 -= 1
			} 

			if  (board[i][j-k] == model.StoneWhite) {
				count1 += 1
			}
		}
	}

	if count0 > count1 {
		return defenseScore(count0, targetCount)		
	} else {
		return defenseScore(count1, targetCount)
	}

}

func dHeatDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {
		if (i < len(board)-k) && (j < len(board)-k) {
			if (board[i+k][j+k] == model.StoneBlack) {
				count0 -= 1
			} 

			if (board[i+k][j+k] == model.StoneWhite) {
				count0 += 1
			}
		}

		if (i-k > 0) && (j-k > 0) {
			if (board[i-k][j-k] == model.StoneBlack) {
				count1 -= 1
			} 

			if (board[i-k][j-k] == model.StoneWhite) {
				count1 += 1
			}	
		}
	}
	
	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}
}


func dHeatDefense2(i int, j int, targetCount int, board [][]model.Stone)(int) {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (i-k > 0) && (j < len(board)-k) {
			if (board[i-k][j+k] == model.StoneBlack) {
				count0 -= 1
			} 

			if (board[i-k][j+k] == model.StoneWhite) {
				count0 += 1
			} 
		}

		if (i < len(board)-k) && (j-k > 0) {
			if (board[i+k][j-k] == model.StoneBlack) {
				count1 -= 1
			} 

			if (board[i+k][j-k] == model.StoneWhite) {
				count1 += 1
			}
		}

	}
	
	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}
}


func vHeatDefense(i int, j int, targetCount int, board [][]model.Stone)(int) {
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

func vPotential(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore 		
	} 


	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}
			if (i < len(board)-k-m) && (i+k-m > 0) {
				if (board[i+k-m][j] == targetStone) {
					count += 1
				}

				if (board[i+k-m][j] == notTargetStone) {
					count = 0
					break
				}
			}
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return scoreFunc(maxCount, maxCount)
}

func hPotential(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore 		
	} 

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if (j < len(board)-k-m) && (j+k-m > 0) {
				if (board[i][j+k-m] == targetStone) {
					count += 1
				}

				if (board[i][j+k-m] == notTargetStone) {
					count = 0
					break
				}
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}


func dPotential(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore 		
	} 

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if (j < len(board)-k-m) && (j+k-m > 0) && (i < len(board)-k-m) && (i+k-m > 0) {
				if (board[i+k-m][j+k-m] == targetStone) {
					count += 1
				}

				if (board[i+k-m][j+k-m] == notTargetStone) {
					count = 0
					break
				}
			}
			

			
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}


func dPotential2(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore 		
	} 

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if (j < len(board)-k-m) && (j+k-m > 0) && (i-k-m > 0) {
				if (board[i-k-m][j+k-m] == targetStone) {
					count += 1
				}

				if (board[i-k-m][j+k-m] == notTargetStone) {
					count = 0
					break
				}
			}
			
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}

// --ooo

// --ooo--
// -o-oo-
// -oo-o-


// -o-o-o-
// -ooo-o-

func threeHorizontal(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count := 0
	for k := 0; k < 7; k++ {
		
		if (j < len(board)-k) {
			stone := piece
			if (k == 0) || (k == 1) || (k == 5) || (k == 6) {
				stone = model.StoneEmpty
			} else {
				stone = piece
			}

			if (board[i][j+k] == stone) {
				count += 1
			} 
		}	
	}

	if count == 7 {
		return 100
	}

	count = 0
	for k := 0; k < 6; k++ {
		if (j < len(board)-k) {
			if (k == 0) || (k == 5) {
				if board[i][j+k] != model.StoneEmpty {
					return 0
				}
			} else {
				if board[i][j+k] == piece {
					count += 1
				} else if board[i][j+k] != model.StoneEmpty {
					return 0
				}
			}
		}
		
	}

	if count == 3 {
		return 100		
	} else {
		return 0
	}
}


func threeVertical(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count := 0
	for k := 0; k < 7; k++ {

		if (i < len(board)-k) {
			stone := piece
			if (k == 0) || (k == 1) || (k == 5) || (k == 6) {
				stone = model.StoneEmpty
			} else {
				stone = piece
			}

			if (board[i+k][j] == stone) {
				count += 1
			} 
		}
	}

	if count == 7 {
		return 100
	}

	count = 0
	for k := 0; k < 6; k++ {
		if (i < len(board)-k) {
			if (k == 0) || (k == 5) {
				if board[i+k][j] != model.StoneEmpty {
					return 0
				}
			} else {
				if board[i+k][j] == piece {
					count += 1
				} else if board[i+k][j] != model.StoneEmpty {
					return 0
				}
			}
		}
		
	}

	if count == 3 {
		return 100		
	} else {
		return 0
	}
}



func threeDiagonal(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count := 0
	for k := 0; k < 7; k++ {

		stone := piece
		if (k == 0) || (k == 1) || (k == 5) || (k == 6) {
			stone = model.StoneEmpty
		} else {
			stone = piece
		}

		if (j < len(board)-k) && (i < len(board)-k) && (board[i+k][j+k] == stone) {
			count += 1
		}
	}

	if count == 7 {
		return 100
	}

	count = 0
	for k := 0; k < 6; k++ {
		if (j < len(board)-k) && (i < len(board)-k) {
			if (k == 0) || (k == 5) {
				if board[i+k][j+k] != model.StoneEmpty {
					return 0
				}
			} else {
				if board[i+k][j+k] == piece {
					count += 1
				} else if board[i+k][j+k] != model.StoneEmpty {
					return 0
				}
			}
		}
		
	}

	if count == 3 {
		return 100		
	} else {
		return 0
	}
}


func threeDiagonal2(i int, j int, piece model.Stone, board [][]model.Stone)(int) {
	count := 0
	for k := 0; k < 7; k++ {

		stone := piece
		if (k == 0) || (k == 1) || (k == 5) || (k == 6) {
			stone = model.StoneEmpty
		} else {
			stone = piece
		}

		if (j < len(board)-k) && (i-k > 0) && (board[i-k][j+k] == stone) {
			count += 1
		}
	}

	if count == 7 {
		return 100
	}

	count = 0
	for k := 0; k < 6; k++ {
		if (j < len(board)-k) && (i-k > 0) {
			if (k == 0) || (k == 5) {
				if board[i-k][j+k] != model.StoneEmpty {
					return 0
				}
			} else {
				if board[i-k][j+k] == piece {
					count += 1
				} else if board[i-k][j+k] != model.StoneEmpty {
					return 0
				}
			}
		}
	}

	if count == 3 {
		return 100
	} else {
		return 0
	}
}


func hHeatOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
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

func vHeatOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
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


func dHeatOffense(i int, j int, targetCount int, board [][]model.Stone)(int) {
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


func dHeatOffense2(i int, j int, targetCount int, board [][]model.Stone)(int) {
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

				for k := 1; k < 4; k++ {
					tmp[i][j] += hHeatDefense(i, j, k, board)					
					tmp[i][j] += vHeatDefense(i, j, k, board)					

					// tmp[i][j] += dHeatDefense(i, j, k, board)
					// tmp[i][j] += dHeatDefense2(i, j, k, board)

					tmp[i][j] += hHeatOffense(i, j, k, board)
					tmp[i][j] += vHeatOffense(i, j, k, board)

					// tmp[i][j] += dHeatOffense(i, j, k, board)
					// tmp[i][j] += dHeatOffense2(i, j, k, board)
				}

				tmp[i][j] += hPotential(i, j, model.StoneBlack, board)
				tmp[i][j] += hPotential(i, j, model.StoneWhite, board)
				tmp[i][j] += vPotential(i, j, model.StoneBlack, board)
				tmp[i][j] += vPotential(i, j, model.StoneWhite, board)

				// tmp[i][j] += dPotential(i, j, model.StoneBlack, board)
				// tmp[i][j] += dPotential(i, j, model.StoneWhite, board)
				// tmp[i][j] += dPotential2(i, j, model.StoneBlack, board)
				// tmp[i][j] += dPotential2(i, j, model.StoneWhite, board)


				tmp[i][j] += threeHorizontal(i, j, model.StoneWhite, board)
				tmp[i][j] += threeHorizontal(i, j, model.StoneBlack, board)
				tmp[i][j] += threeVertical(i, j, model.StoneWhite, board)
				tmp[i][j] += threeVertical(i, j, model.StoneBlack, board)
				// tmp[i][j] += threeDiagonal(i, j, model.StoneWhite, board)
				// tmp[i][j] += threeDiagonal(i, j, model.StoneBlack, board)
				// tmp[i][j] += threeDiagonal2(i, j, model.StoneWhite, board)
				// tmp[i][j] += threeDiagonal2(i, j, model.StoneBlack, board)

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
