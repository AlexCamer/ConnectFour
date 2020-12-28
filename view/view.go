package view

import (
	"connectfour/game"
	"fmt"
)

var viewTemplate = []rune(
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	"[ ][ ][ ][ ][ ][ ][ ]\n" +
	" 1  2  3  4  5  6  7\n")

func DisplayBoard(g *game.Game) {
	view := viewTemplate
	board := game.Board(g)
	for i := 0; i < game.HEIGHT; i++ {
		for j := 0; j < game.WIDTH; j++ {
			var symbol rune
			switch board[i][j] {
			case game.RED:
				symbol = 'R'
			case game.YELLOW:
				symbol = 'Y'
			case game.NONE:
				continue
			}
			view[22 * i + 3 * j + 1] = symbol
		}
	}
	fmt.Print(string(view))
}

func toString(colour game.Colour) string {
	switch colour {
	case game.RED:
		return "red"
	case game.YELLOW:
		return "yellow"
	default:
		return ""
	}
}

func DisplayTurn(g *game.Game) {
	fmt.Println(toString(game.Turn(g)), "turn:")
}

func DisplayOutcome (g *game.Game) {
	if game.Winner(g) == game.NONE {
		fmt.Println("draw")
	} else {
		fmt.Println(toString(game.Winner(g)), "wins!")
	}
}

func Input() (int, error) {
	var input int
	_, err := fmt.Scanf("%d", &input)
	return input, err
}
