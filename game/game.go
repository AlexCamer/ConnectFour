package game

import "errors"

type Game struct {
	board  board
	turn   Colour
	winner Colour
}

func New() Game {
	return Game{board: board{}, turn: randomColour(), winner: NONE}
}

func Insert(g *Game, position int) error {
	if Over(g) {
		return errors.New("game is over")
	}
	error := insert(&g.board, position, g.turn)
	if error != nil {
		return error
	}
	g.winner = connectFour(&g.board)
	switchTurn(g)
	return nil
}

func switchTurn(g *Game) {
	if g.turn == RED {
		g.turn = YELLOW
	} else {
		g.turn = RED
	}
}

func Over(g *Game) bool {
	return g.winner != NONE || full(&g.board)
}

func Board(g *Game) board {
	return g.board
}

func Turn(g *Game) Colour {
	return g.turn
}

func Winner(g *Game) Colour {
	return g.winner
}
