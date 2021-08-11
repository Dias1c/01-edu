package tetrisoptimizer

func _FindSmallestSquare(tetrimonies []*tetrimony, board *board, step byte) bool {
	if len(tetrimonies) == 0 {
		return true
	}
	curTetrimony := tetrimonies[0]
	for y := 0; y+curTetrimony.Height <= board.Width; y++ {
		for x := 0; x+curTetrimony.Width <= board.Width; x++ {
			canPush := board.CanPushTetrimony(x, y, curTetrimony)
			if canPush {
				board.PushTetrimony(x, y, curTetrimony, step)
				isMatch := _FindSmallestSquare(tetrimonies[1:], board, step+1)
				if isMatch {
					return true
				}
				board.RemoveTetrimony(x, y, curTetrimony, step)
			}
		}
	}
	return false
}
