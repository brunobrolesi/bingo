package usecase

import "github.com/brunobrolesi/bingo/internal/model"

type StartGameUseCase interface {
	Start() model.Balls
}

type startGame struct{}

func NewStartGameUseCase() StartGameUseCase {
	return &startGame{}
}

func (s *startGame) Start() model.Balls {
	balls := model.Balls{}
	for i := 1; i <= 15; i++ {
		ball := model.Ball{Number: i, Letter: "B"}
		balls = append(balls, ball)
	}
	for i := 16; i <= 30; i++ {
		ball := model.Ball{Number: i, Letter: "I"}
		balls = append(balls, ball)
	}
	for i := 31; i <= 45; i++ {
		ball := model.Ball{Number: i, Letter: "N"}
		balls = append(balls, ball)
	}
	for i := 46; i <= 60; i++ {
		ball := model.Ball{Number: i, Letter: "G"}
		balls = append(balls, ball)
	}
	for i := 61; i <= 75; i++ {
		ball := model.Ball{Number: i, Letter: "O"}
		balls = append(balls, ball)
	}

	return balls
}
