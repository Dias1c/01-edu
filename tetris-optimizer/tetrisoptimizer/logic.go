package tetrisoptimizer

func _FindSmallestSquare(tetrimonies []*tetrimony, board *board, step byte) bool {
	// if len(tetrimonies) == 1 {
	// 	board.PrintBoard()
	// }
	if len(tetrimonies) == 0 {
		return true
	}
	for y := 0; y < board.Width; y++ {
		for x := 0; x < board.Width; x++ {
			if board.Field[y][x] == 0 {
				curTetrimony := tetrimonies[0]
				if x+curTetrimony.Width > board.Width || y+curTetrimony.Height > board.Width {
					break
				}
				isPushed := board.PushTetrimony(x, y, curTetrimony, step)
				if isPushed {
					isMatch := _FindSmallestSquare(tetrimonies[1:], board, step+1)
					if isMatch {
						return true
					}
				}
				board.RemoveTetrimony(x, y, curTetrimony, step)
			}
		}
	}
	return false
}

//65-46 = A-.
