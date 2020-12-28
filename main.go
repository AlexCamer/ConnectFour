package main

import (
	"connectfour/game"
	"connectfour/view"
	"fmt"
)

func main() {
	g := game.New()
	view.DisplayBoard(&g)
	for !game.Over(&g) { // logic loop; while game is not over...
		view.DisplayTurn(&g) // display which turn
		position, err := view.Input() // get move from user
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		err = game.Insert(&g, position) // input move into game
		if err != nil {
			fmt.Println(err)
			continue
		}
		view.DisplayBoard(&g) // display the board
	}
	view.DisplayOutcome(&g) // game is over; display outcome
}
