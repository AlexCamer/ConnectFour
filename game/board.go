package game

import (
	"errors"
)

const (
	HEIGHT = 6
	WIDTH = 7
)

type board [HEIGHT][WIDTH]Colour

func insert(b *board, position int, colour Colour) error {
	if position < 1 || position > WIDTH {
		return errors.New("invalid position")
	}
	for i := HEIGHT - 1; i >= 0; i-- {
		if b[i][position - 1] == NONE {
			b[i][position - 1] = colour
			return nil
		}
	}
	return errors.New("column full")
}

func checkFour(b *board, y, x, yInc, xInc int, c chan<- Colour) {
	colour := b[y][x]
	for i := 1; i < 4; i++ {
		if b[y + i * yInc][x + i * xInc] != colour {
			c <- NONE
			return
		}
	}
	c <- colour
}

func connectFour(b *board) Colour {
	c := make(chan Colour)
	counter := 0
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if i < HEIGHT - 3 {
				go checkFour(b, i, j, 1, 0, c)
				counter++
			}
			if j < WIDTH - 3 {
				go checkFour(b, i, j, 0, 1, c)
				counter++
				if i < HEIGHT - 3 {
					go checkFour(b, i, j, 1, 1, c)
					counter++
				}
				if i > 2 {
					go checkFour(b, i, j, -1, 1, c)
					counter++
				}
			}
		}
	}
	results := make([]Colour, counter)
	for i := 0; i < counter; i++ {
		results = append(results, <- c)
	}
	close(c)
	for _, colour := range results {
		if colour != NONE {
			return colour
		}
	}
	return NONE
}

func full(b *board) bool {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if b[i][j] == NONE {
				return false
			}
		}
	}
	return true
}
