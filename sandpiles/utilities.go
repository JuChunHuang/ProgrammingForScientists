package main

import "canvas"

// DrawGameBoard draws the board into a PNG file.
// Input: a GameBoard and a file name of the PNG file.
func (board *GameBoard) DrawGameBoard(file_name string) {
	height := len(*board)
	width := len((*board)[0])
	c := canvas.CreateNewCanvas(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (*board)[i][j] == 0 {
				c.SetFillColor(canvas.MakeColor(0, 0, 0))
			} else if (*board)[i][j] == 1 {
				c.SetFillColor(canvas.MakeColor(85, 85, 85))
			} else if (*board)[i][j] == 2 {
				c.SetFillColor(canvas.MakeColor(170, 170, 170))
			} else if (*board)[i][j] == 3 {
				c.SetFillColor(canvas.MakeColor(255, 255, 255))
			} else {
				panic("Error: sandpiles value out of range!")
			}
			c.ClearRect(j, i, j+1, i+1)
			c.Fill()
		}
	}

	c.SaveToPNG(file_name + ".png")
}

// IsStable determines whether the configuration of coins in the given board is stable.
// Input: a GameBoard.
// Output: a boolean shows that the given GameBoard is stable or not.
func (board *GameBoard) IsStable() bool {
	num_rows := len(*board)
	num_cols := len((*board)[0])
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if (*board)[i][j] >= 4 {
				return false
			}
		}
	}
	return true
}

// InField determines whether the given cell is inside the GameBoard.
// Input: a row index and column index of the target cell; a number of rows and a number of columns.
// Output: a boolean shows that the given cell is inside the GameBoard or not.
func InField(i, j, num_rows, num_cols int) bool {
	if i >= 0 && i < num_rows && j >= 0 && j < num_cols {
		return true
	}

	return false
}

// InitializeNewBoards intializes a default GameBoard based on the given number of rows and columns.
// Input: a number of rows and a number of columns.
// Output: a GameBoard.
func InitializeNewBoard(numRows, numCols int) GameBoard {
	var board GameBoard
	board = make(GameBoard, numRows)

	for r := range board {
		board[r] = make([]int, numCols)
	}

	return board
}

// ResetBoard resets all the cells in the given GameBoard to 0.
// Input: a GameBoard.
func (board *GameBoard) ResetBoard() {
	num_rows := len(*board)
	num_cols := len((*board)[0])
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			(*board)[i][j] = 0
		}
	}
}

// CopyBoard makes a copy of the given board.
func (board *GameBoard) CopyBoard() GameBoard {
	num_rows := len(*board)
	num_cols := len((*board)[0])
	new_board := make(GameBoard, num_rows)
	for i := 0; i < num_rows; i++ {
		new_board[i] = make([]int, num_cols)
		new_board[i] = (*board)[i]
	}

	return new_board
}
