package game

import (
	"math/rand"
)

// piece colour type
type Colour int

const (
	NONE Colour = iota
	RED
	YELLOW
)

// returns a random colour that is not NONE
func randomColour() Colour {
	colours := []Colour{RED, YELLOW}
	return colours[rand.Int() % len(colours)]
}
