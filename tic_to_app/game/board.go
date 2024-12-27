// game/board.go
package game

type Board struct {
    Cells [9]string
}

func NewBoard() *Board {
    return &Board{
        Cells: [9]string{},
    }
}

func (b *Board) MakeMove(position int, player string) bool {
    if position < 0 || position > 8 || b.Cells[position] != "" {
        return false
    }
    b.Cells[position] = player
    return true
}

func (b *Board) CheckWinner() string {
    // Winning combinations
    lines := [][3]int{
        {0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
        {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
        {0, 4, 8}, {2, 4, 6}, // Diagonals
    }

    for _, line := range lines {
        if b.Cells[line[0]] != "" &&
            b.Cells[line[0]] == b.Cells[line[1]] &&
            b.Cells[line[0]] == b.Cells[line[2]] {
            return b.Cells[line[0]]
        }
    }
    return ""
}

