package main

import (
	"fmt"
	"testing"
)

func TestInitializeNewBoard(t *testing.T) {
	type test struct {
		i, j int
		ans  GameBoard
	}

	var test_case = test{1, 2, GameBoard{{0, 0}}}
	outcome := InitializeNewBoard(test_case.i, test_case.j)
	fmt.Println(outcome)
}

func TestResetBoard(t *testing.T) {
	type test struct {
		b *GameBoard
	}
	var b = &GameBoard{{0, 1, 0}, {1, 1, 1}, {2, 1, 1}}
	var test_case = test{b}

	(test_case.b).ResetBoard()
	fmt.Println(test_case.b)
}

func TestIsStable(t *testing.T) {
	type test struct {
		b   *GameBoard
		ans bool
	}
	var b = &GameBoard{{0, 1, 0}, {1, 1, 1}, {2, 1, 1}}
	var test_case = test{b, true}

	outcome := (test_case.b).IsStable()
	if outcome != test_case.ans {
		t.Errorf("Wrong stable detection!")
	} else {
		fmt.Println("Pass")
	}
}

func TestTopple(t *testing.T) {
	type test struct {
		b    *GameBoard
		i, j int
		ans  *GameBoard
	}

	var b = &GameBoard{{0, 0, 0}, {0, 4, 0}, {0, 0, 0}}
	var ans = &GameBoard{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}
	i := 1
	j := 1

	var test_case = test{b, i, j, ans}

	(test_case.b).Topple(test_case.i, test_case.j)
	fmt.Println(test_case.b)
	fmt.Println(test_case.ans)
}

func TestInField(t *testing.T) {
	type test struct {
		i, j, num_rows, num_cols int
		ans                      bool
	}

	i := 1
	j := 2
	num_rows := 3
	num_cols := 4
	ans := true
	var test_case = test{i, j, num_rows, num_cols, ans}
	outcome := InField(test_case.i, test_case.j, test_case.num_rows, test_case.num_cols)

	if outcome != test_case.ans {
		t.Errorf("Wrong InField detection!")
	} else {
		fmt.Println("Pass")
	}
}

func TestCreatePartialBoard(t *testing.T) {
	type test struct {
		b1                             GameBoard
		start, end, num_rows, num_cols int
		answer                         GameBoard
	}

	var b1 = GameBoard{{1, 2, 3}, {4, 5, 6}}
	start := 1
	end := 3
	num_rows := 5
	num_cols := 3
	var ans = GameBoard{{0, 0, 0}, {1, 2, 3}, {4, 5, 6}, {0, 0, 0}, {0, 0, 0}}
	var test_case = test{b1, start, end, num_rows, num_cols, ans}

	outcome := CreatePartialBoard(test_case.b1, test_case.start, test_case.end, test_case.num_rows, test_case.num_cols)
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if outcome[i][j] != test_case.answer[i][j] {
				t.Errorf("Error! At (%d, %d), output: (%d) but the answer is: (%d)", i, j, outcome[i][j], test_case.answer[i][j])
			}
		}
	}
	fmt.Println("Pass!")
}
