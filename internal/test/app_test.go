package test

import (
	"testing"

	"github.com/brunobrolesi/bingo/internal/model"
	"github.com/brunobrolesi/bingo/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	t.Run("should run the game correctly", func(t *testing.T) {
		start := usecase.NewStartGameUseCase()
		referenceBalls := start.Exec()
		balls := make(model.Balls, len(referenceBalls))
		copy(balls, referenceBalls)
		shuffle := usecase.NewShuffleUseCase()
		getRandomBall := usecase.NewGetRandomBallUseCase()
		addDrawnBall := usecase.NewAddDrawnBallUseCase()
		for len(balls) > 0 {
			balls = shuffle.Exec(balls)
			ball, newBalls := getRandomBall.Exec(balls)
			balls = newBalls
			drawnBall := addDrawnBall.Exec(ball)
			assert.NotContains(t, balls, ball)
			assert.Contains(t, drawnBall[ball.Letter], ball.Number)
		}
	})
}
