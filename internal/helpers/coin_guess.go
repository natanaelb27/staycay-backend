package helpers

import "math/rand"

type CoinGuess interface {
	Randomizer(n int) int
}

func Randomizer(n int) int {
	return rand.Intn(n)
}
