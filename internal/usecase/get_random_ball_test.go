package usecase

import (
	"testing"

	"github.com/brunobrolesi/bingo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomBallUseCase(t *testing.T) {
	referenceBalls := model.Balls{
		model.Ball{Number: 1, Letter: "B"},
		model.Ball{Number: 2, Letter: "I"},
		model.Ball{Number: 3, Letter: "N"},
		model.Ball{Number: 4, Letter: "G"},
		model.Ball{Number: 5, Letter: "O"},
	}
	t.Run("should get random ball", func(t *testing.T) {
		balls := make(model.Balls, len(referenceBalls))
		copy(balls, referenceBalls)

		sut := NewGetRandomBallUseCase()
		ball, newBalls := sut.GetRandomBall(balls)

		assert.Equal(t, len(referenceBalls)-1, len(newBalls))
		assert.Contains(t, referenceBalls, ball)
		assert.NotContains(t, newBalls, ball)
	})

}
