package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

type GameBoard [][]int

func main() {
	size, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	num, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err1)
	}
	placement := os.Args[3]

	board := InitializeNewBoard(size, size)

	// place the coins based on the instruction
	if placement == "central" {
		board[int(size/2)][int(size/2)] = num
	} else {
		for i := 0; i < 100; i++ {
			x := rand.Intn(size)
			y := rand.Intn(size)
			board[x][y] = num / 100
		}
	}

	num_procs := runtime.NumCPU()
	copy_board := board.CopyBoard()
	board.SandpilesSerialTime()
	(&copy_board).SandpilesParallelTime(num_procs)
}

// SandpilesSerialTime finds the stable configuration of sandpiles serially and print the time cost.
// Input: an initial board.
func (board *GameBoard) SandpilesSerialTime() {
	start_serial := time.Now()
	board.SandpilesSerial()
	elapsed_serial := time.Since(start_serial)
	log.Printf("Simulating sandpiles in serial took %s", elapsed_serial)
	board.DrawGameBoard("serial")
}

// SandpilesParallelTime finds the stable configuration of sandpiles in parallel and print the time cost.
// Input: an initial board.
func (board *GameBoard) SandpilesParallelTime(num_procs int) {
	start_parallel := time.Now()
	board.SandpilesMultiprocs(num_procs)
	elapsed_parallel := time.Since(start_parallel)
	log.Printf("Simulating sandpiles in parallel took %s", elapsed_parallel)
	board.DrawGameBoard("parallel")
}
