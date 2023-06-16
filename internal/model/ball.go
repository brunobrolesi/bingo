package model

import "fmt"

type Ball struct {
	Number int
	Letter string
}

type Balls []Ball

func (b *Ball) String() string {
	return fmt.Sprintf("Letra: %s - Número: %d", b.Letter, b.Number)
}
