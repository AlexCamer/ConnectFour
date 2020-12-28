package game

import (
	"errors"
	"math/rand"
	"time"
)

// connect-four game struct
type Game struct {
	board  board
	turn   Colour
	winner Colour
}

// returns new game with random initial turn
func New() Game {
	rand.Seed(time.Now().UnixNano())
	return Game{board: board{}, turn: randomColour(), winner: NONE}
}

// insert a piece at a position of the respective turn colour
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

// switches turn
func switchTurn(g *Game) {
	if g.turn == RED {
		g.turn = YELLOW
	} else {
		g.turn = RED
	}
}

// is game over?
func Over(g *Game) bool {
	return g.winner != NONE || full(&g.board)
}

// board getter; for view
func Board(g *Game) board {
	return g.board
}

// turn getter; for view
func Turn(g *Game) Colour {
	return g.turn
}

// winner getter; for view
func Winner(g *Game) Colour {
	return g.winner
}
