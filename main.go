package main

import (
	"connectfour/game"
	"connectfour/view"
	"fmt"
)

func main() {
	g := game.New()
	view.DisplayBoard(&g)
	for !game.Over(&g) {
		view.DisplayTurn(&g)
		position, err := view.Input()
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		err = game.Insert(&g, position)
		if err != nil {
			fmt.Println(err)
			continue
		}
		view.DisplayBoard(&g)
	}
	view.DisplayOutcome(&g)
}
