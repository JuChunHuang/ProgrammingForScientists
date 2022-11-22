package main

// SandpilesMultiprocs finds the stable configuration of sandpiles in parallel.
// Input: an initial GameBoard.
func (board *GameBoard) SandpilesMultiprocs(num_procs int) {
	c := make(chan GameBoard, num_procs)
	num_rows := len(*board)
	num_cols := len((*board)[0])

	for !board.IsStable() {
		// split the GameBoard into num_procs approximately equal pieces
		for i := 0; i < num_procs; i++ {
			start_index := i * (num_rows / num_procs)
			end_index := (i + 1) * (num_rows / num_procs)
			if i < num_procs-1 {
				go SandpileSingleproc(CreatePartialBoard((*board)[start_index:end_index], start_index, end_index, num_rows, num_cols), c)

			} else {
				go SandpileSingleproc(CreatePartialBoard((*board)[start_index:], start_index, num_rows, num_rows, num_cols), c)
			}
		}

		board.ResetBoard()

		for i := 0; i < num_procs; i++ {
			board.Combine(c)
		}
	}
}

// CreatePartialBoard expands the size of the given subboard to the original board size by filling with 0.
// Input: a subboard, a start index and an end index of the subboard when it was in the original board,
//        a number of total rows and a number of total columns in the original board.
// Output: a GameBoard which has the same size with the original board.
func CreatePartialBoard(board GameBoard, start, end, num_rows, num_cols int) GameBoard {
	new_board := InitializeNewBoard(num_rows, num_cols)
	for i := start; i < end; i++ {
		for j := 0; j < num_cols; j++ {
			new_board[i][j] = board[i-start][j]
		}
	}
	return new_board
}

// SandpileSingleproc finds the stable configuration of subboard (but with original board size) in parallel.
// Input: a subboard and a channel to store the result.
func SandpileSingleproc(board GameBoard, c chan GameBoard) {
	num_rows := len(board)
	num_cols := len(board[0])
	for !board.IsStable() {
		for i := 0; i < num_rows; i++ {
			for j := 0; j < num_cols; j++ {
				(&board).Topple(i, j)
			}
		}
	}
	c <- board
}

// Combine combines the subboards to the reseted board.
// Input: a reseted board and a channel that store the subboard.
func (board *GameBoard) Combine(c chan GameBoard) {
	tmpBoard := <-c
	num_rows := len(*board)
	num_cols := len((*board)[0])
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			(*board)[i][j] += tmpBoard[i][j]
		}
	}
}
