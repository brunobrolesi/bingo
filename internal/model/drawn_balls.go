package model

import (
	"fmt"
)

type DrawnBalls map[string][]int

func NewDrawnBalls() DrawnBalls {
	return DrawnBalls{"B": {}, "I": {}, "N": {}, "G": {}, "O": {}}
}

func (d DrawnBalls) String() string {
	s := ""
	s += fmt.Sprintln("B: " + fmt.Sprint(d["B"]))
	s += fmt.Sprintln("I: " + fmt.Sprint(d["I"]))
	s += fmt.Sprintln("N: " + fmt.Sprint(d["N"]))
	s += fmt.Sprintln("G: " + fmt.Sprint(d["G"]))
	s += fmt.Sprintln("O: " + fmt.Sprint(d["O"]))
	return s
}
