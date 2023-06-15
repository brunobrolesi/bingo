package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartUseCase(t *testing.T) {
	t.Run("should return all game balls and correct letters", func(t *testing.T) {
		sut := NewStartGameUseCase()
		balls := sut.Start()

		assert.Equal(t, 75, len(balls))
		for i := 1; i <= 15; i++ {
			assert.Equal(t, i, balls[i-1].Number)
			assert.Equal(t, "B", balls[i-1].Letter)
		}
		for i := 16; i <= 30; i++ {
			assert.Equal(t, i, balls[i-1].Number)
			assert.Equal(t, "I", balls[i-1].Letter)
		}
		for i := 31; i <= 45; i++ {
			assert.Equal(t, i, balls[i-1].Number)
			assert.Equal(t, "N", balls[i-1].Letter)
		}
		for i := 46; i <= 60; i++ {
			assert.Equal(t, i, balls[i-1].Number)
			assert.Equal(t, "G", balls[i-1].Letter)
		}
		for i := 61; i <= 75; i++ {
			assert.Equal(t, i, balls[i-1].Number)
			assert.Equal(t, "O", balls[i-1].Letter)
		}
	})
}
