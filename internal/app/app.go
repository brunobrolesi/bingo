package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/brunobrolesi/bingo/internal/tools"
	"github.com/brunobrolesi/bingo/internal/usecase"
)

func Start() {
	star := usecase.NewStartGameUseCase()
	shuffle := usecase.NewShuffleUseCase()
	getRandomBall := usecase.NewGetRandomBallUseCase()
	addDrawnBall := usecase.NewAddDrawnBallUseCase()

	scanner := bufio.NewScanner(os.Stdin)

	balls := star.Exec()

	for len(balls) > 0 {
		tools.CallClear()

		spins := getNumberOfSpins(scanner)
		for i := 0; i < spins; i++ {
			balls = shuffle.Exec(balls)
		}
		ball, newBalls := getRandomBall.Exec(balls)
		balls = newBalls

		fmt.Println("BOLA DA VEZ: ", ball.String())
		fmt.Println("-----------------------------------")

		drawnBalls := addDrawnBall.Exec(ball)

		fmt.Println("BOLAS SORTEADAS: ")
		fmt.Println(drawnBalls.String())
		fmt.Println("-----------------------------------")
		fmt.Println("BOLAS RESTANTES: ", len(balls))
		fmt.Println("-----------------------------------")
		readyForNext(scanner)
	}

	fmt.Println("FIM DE JOGO")
}

func getNumberOfSpins(scanner *bufio.Scanner) int {
	fmt.Println("Quantas vezes quer rodar a roleta?")
	scanner.Scan()
	text := scanner.Text()
	spins, err := strconv.Atoi(text)
	for err != nil || spins < 1 {
		fmt.Println("Quantas vezes quer rodar a roleta?")
		scanner.Scan()
		text = scanner.Text()
		spins, err = strconv.Atoi(text)
	}
	tools.CallClear()
	return spins
}

func readyForNext(scanner *bufio.Scanner) {
	fmt.Println("Pronto para a próxima?")
	scanner.Scan()
}
