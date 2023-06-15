package usecase

import (
	"math/rand"
	"time"

	"github.com/brunobrolesi/bingo/internal/model"
)

type ShuffleUseCase interface {
	Shuffle(balls model.Balls) model.Balls
}

type shuffleUseCase struct{}

func NewShuffleUseCase() ShuffleUseCase {
	return &shuffleUseCase{}
}

func (s *shuffleUseCase) Shuffle(balls model.Balls) model.Balls {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(balls), func(i, j int) { balls[i], balls[j] = balls[j], balls[i] })

	return balls
}
