package model

import "fmt"

type Ball struct {
	Number int
	Letter string
}

type Balls []Ball

func (b *Ball) String() string {
	return b.Letter + "-" + fmt.Sprint(b.Number)
}
