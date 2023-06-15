package usecase

import (
	"math/rand"
	"time"

	"github.com/brunobrolesi/bingo/internal/model"
)

type GetRandomBallUseCase interface {
	GetRandomBall(balls model.Balls) (model.Ball, model.Balls)
}

type getRandomBall struct{}

func NewGetRandomBallUseCase() GetRandomBallUseCase {
	return &getRandomBall{}
}

func (g *getRandomBall) GetRandomBall(balls model.Balls) (model.Ball, model.Balls) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(balls))
	randomBall := balls[index]
	newBalls := g.removeBall(balls, index)
	return randomBall, newBalls
}

func (g *getRandomBall) removeBall(balls model.Balls, index int) model.Balls {
	return append(balls[:index], balls[index+1:]...)
}
