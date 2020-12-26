package game

import (
	"math/rand"
)

type Colour int

const (
	NONE Colour = iota
	RED
	YELLOW
)

func randomColour() Colour {
	colours := []Colour{RED, YELLOW}
	return colours[rand.Int() % len(colours)]
}
