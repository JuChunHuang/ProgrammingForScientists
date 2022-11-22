package main

// SandpilesSerial finds the stable configuration of sandpiles serially.
// Input: an initial GameBoard.
func (board *GameBoard) SandpilesSerial() {
	num_rows := len(*board)
	num_cols := len((*board)[0])

	for !board.IsStable() {
		// range over all cells and do topple operations
		for i := 0; i < num_rows; i++ {
			for j := 0; j < num_cols; j++ {
				board.Topple(i, j)
			}
		}
	}
}

// Topple does a topple operation of the given cell.
// Input: a row index and column index of the target cell.
func (board *GameBoard) Topple(i, j int) {
	num_rows := len(*board)
	num_cols := len((*board)[0])
	if (*board)[i][j] >= 4 {
		(*board)[i][j] -= 4
		if InField(i-1, j, num_rows, num_cols) {
			(*board)[i-1][j]++
		}
		if InField(i+1, j, num_rows, num_cols) {
			(*board)[i+1][j]++
		}
		if InField(i, j-1, num_rows, num_cols) {
			(*board)[i][j-1]++
		}
		if InField(i, j+1, num_rows, num_cols) {
			(*board)[i][j+1]++
		}
	}
}
