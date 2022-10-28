package helpers

import "math/rand"

type CoinGuessHelper interface {
	Randomizer(n int) int
}

type coinGuessHelper struct {
}

type CGHConfig struct {
}

func NewCoinGuessHelper(c *CGHConfig) CoinGuessHelper {
	return &coinGuessHelper{}
}

func (h *coinGuessHelper) Randomizer(n int) int {
	return rand.Intn(n)
}
