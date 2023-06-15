package usecase

import (
	"fmt"
	"sort"

	"github.com/brunobrolesi/bingo/internal/model"
)

var drawnBalls model.DrawnBalls

func init() {
	drawnBalls = model.NewDrawnBalls()
}

type AddDrawnBallUseCase interface {
	AddDrawnBall(ball model.Ball) model.DrawnBalls
}

type addDrawnBallUseCase struct{}

func NewAddDrawnBallUseCase() AddDrawnBallUseCase {
	return addDrawnBallUseCase{}
}

func (a addDrawnBallUseCase) AddDrawnBall(ball model.Ball) model.DrawnBalls {
	drawnBalls[ball.Letter] = append(drawnBalls[ball.Letter], ball.Number)
	sort.Slice(drawnBalls[ball.Letter], func(i, j int) bool { return drawnBalls[ball.Letter][i] < drawnBalls[ball.Letter][j] })
	fmt.Println(drawnBalls.String())
	return drawnBalls
}
