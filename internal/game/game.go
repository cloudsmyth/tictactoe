package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Rows     []Row
	Col      int
	RowCount int
	Player   bool
}

type Row []string

func NewRow(size int) Row {
	row := make(Row, size)
	return row
}

func NewBoard(rowCount int) Board {
	var rows []Row
	for i := 0; i < rowCount; i++ {
		rows = append(rows, NewRow(rowCount))
	}
	return Board{
		Rows:     rows,
		RowCount: rowCount,
		Player:   true,
	}
}

func PlayGame(scanner *bufio.Scanner) bool {
	board := NewBoard(3)

	gameOver := false

	for !gameOver {
		board.PrintBoard()
		currentPlayer := board.GetCurrentPlayer()
		fmt.Printf("It is '%s' players turn. Pleas enter a row and column: ", currentPlayer)

		scanner.Scan()
		input := scanner.Text()

		// Get space to place move on
		coordinates := strings.Split(input, " ")
		if len(coordinates) != 2 {
			fmt.Println("Invalid input! Please enter a row and column number seperated by a space!")
			continue
		}

		row, errR := strconv.Atoi(coordinates[0])
		col, errC := strconv.Atoi(coordinates[1])

		if errR != nil || errC != nil {
			fmt.Println("Invalid input! Please enter a number for row and col")
			continue
		}

		validMove := board.MakeMove(row, col)
		if !validMove {
			continue
		}

		winner := board.CheckWinner()
		if winner != "" {
			board.PrintBoard()
			fmt.Printf("Player %s wins the game!\n", winner)
			gameOver = true
			continue
		}

		if board.GameFinished() {
			board.PrintBoard()
			fmt.Println("Draw!")
			gameOver = true
			continue
		}
	}

	fmt.Print("Play again? (y/n): ")
	scanner.Scan()
	playAgain := strings.ToLower(scanner.Text())

	return playAgain == "y"
}

func (b *Board) GetCurrentPlayer() string {
	if b.Player {
		return "X"
	}
	return "O"
}

func (b *Board) MakeMove(row int, col int) bool {
	piece := "X"
	if b.Player == false {
		piece = "O"
	}
	if row < 1 || row > b.RowCount || col < 1 || col > b.RowCount {
		fmt.Println("Move is out of bounds")
		return false
	}
	if b.Rows[row-1][col-1] != "" {
		fmt.Println("That space is already occupied!")
		return false
	}
	b.Rows[row-1][col-1] = piece
	b.Player = !b.Player
	return true
}

func (b *Board) CheckWinner() string {
	// For the regular rows
	for i := 0; i < b.RowCount; i++ {
		if b.CheckLine(b.Rows[i]) != "" {
			return b.CheckLine(b.Rows[i])
		}
	}

	// For the columns
	for i := 0; i < b.RowCount; i++ {
		column := make([]string, b.RowCount)
		for j := 0; j < b.RowCount; j++ {
			column[j] = b.Rows[j][i]
		}
		if b.CheckLine(column) != "" {
			return b.CheckLine(column)
		}
	}

	// For the diagonals
	diagonal := make([]string, b.RowCount)
	for i := 0; i < b.RowCount; i++ {
		diagonal[i] = b.Rows[i][i]
	}
	if b.CheckLine(diagonal) != "" {
		return b.CheckLine(diagonal)
	}

	// Next diagonal
	for i := 0; i < b.RowCount; i++ {
		diagonal[i] = b.Rows[i][b.RowCount-1-i]
	}
	if b.CheckLine(diagonal) != "" {
		return b.CheckLine(diagonal)
	}
	return ""
}

func (b *Board) GameFinished() bool {
	for i := 0; i < b.RowCount; i++ {
		for j := 0; j < b.RowCount; j++ {
			if b.Rows[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func (b *Board) CheckLine(line []string) string {
	if line[0] == "" {
		return ""
	}

	for i := 1; i < len(line); i++ {
		if line[i] != line[0] {
			return ""
		}
	}

	return line[0]
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.RowCount; i++ {
		fmt.Println(b.Rows[i])
	}
	fmt.Println("----------------")
}
