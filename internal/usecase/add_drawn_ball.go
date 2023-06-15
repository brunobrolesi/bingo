package usecase

import (
	"sort"

	"github.com/brunobrolesi/bingo/internal/model"
)

var drawnBalls model.DrawnBalls

func init() {
	drawnBalls = model.NewDrawnBalls()
}

type AddDrawnBallUseCase interface {
	Exec(ball model.Ball) model.DrawnBalls
}

type addDrawnBallUseCase struct{}

func NewAddDrawnBallUseCase() AddDrawnBallUseCase {
	return addDrawnBallUseCase{}
}

func (a addDrawnBallUseCase) Exec(ball model.Ball) model.DrawnBalls {
	drawnBalls[ball.Letter] = append(drawnBalls[ball.Letter], ball.Number)
	sort.Slice(drawnBalls[ball.Letter], func(i, j int) bool { return drawnBalls[ball.Letter][i] < drawnBalls[ball.Letter][j] })
	return drawnBalls
}
