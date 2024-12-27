// game/computer.go
package game

import "math/rand"

func ComputerMove(board *Board) int {
    // Simple AI: Look for winning move
    for i := 0; i < 9; i++ {
        if board.Cells[i] == "" {
            board.Cells[i] = "O"
            if board.CheckWinner() == "O" {
                board.Cells[i] = ""
                return i
            }
            board.Cells[i] = ""
        }
    }

    // Block player's winning move
    for i := 0; i < 9; i++ {
        if board.Cells[i] == "" {
            board.Cells[i] = "X"
            if board.CheckWinner() == "X" {
                board.Cells[i] = ""
                return i
            }
            board.Cells[i] = ""
        }
    }

    // Random move
    var emptyCells []int
    for i := 0; i < 9; i++ {
        if board.Cells[i] == "" {
            emptyCells = append(emptyCells, i)
        }
    }
    
    if len(emptyCells) > 0 {
        return emptyCells[rand.Intn(len(emptyCells))]
    }
    return -1
}

