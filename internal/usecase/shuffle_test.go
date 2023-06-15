package usecase

import (
	"testing"

	"github.com/brunobrolesi/bingo/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestShuffleUseCase(t *testing.T) {
	referenceBalls := model.Balls{
		model.Ball{Number: 1, Letter: "B"},
		model.Ball{Number: 2, Letter: "I"},
		model.Ball{Number: 3, Letter: "N"},
		model.Ball{Number: 4, Letter: "G"},
		model.Ball{Number: 5, Letter: "O"},
	}
	t.Run("should shuffle balls", func(t *testing.T) {
		balls := make(model.Balls, len(referenceBalls))
		copy(balls, referenceBalls)
		shuffledBalls := NewShuffleUseCase().Exec(balls)

		assert.Equal(t, len(referenceBalls), len(shuffledBalls))
		assert.NotEqual(t, referenceBalls, shuffledBalls)
		assert.ElementsMatch(t, referenceBalls, shuffledBalls)
	})
}
