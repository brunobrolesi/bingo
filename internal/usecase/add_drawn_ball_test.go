package usecase

import (
	"testing"

	"github.com/brunobrolesi/bingo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAddDrawnBall(t *testing.T) {
	referenceBalls := model.Balls{
		model.Ball{Number: 1, Letter: "B"},
		model.Ball{Number: 2, Letter: "I"},
		model.Ball{Number: 3, Letter: "N"},
		model.Ball{Number: 4, Letter: "G"},
		model.Ball{Number: 5, Letter: "O"},
	}
	t.Run("should add a ball to the drawn balls", func(t *testing.T) {
		sut := NewAddDrawnBallUseCase()
		for _, ball := range referenceBalls {
			sut.AddDrawnBall(ball)
			assert.Contains(t, drawnBalls[ball.Letter], ball.Number)
		}
	})
}
