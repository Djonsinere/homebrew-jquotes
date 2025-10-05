package src

import (
	"fmt"
	"math/rand/v2"
)

func Random_rotation() (Quote, error) {
	random_quote_index := rand.IntN(len(Quote_bank))
	if random_quote_index < 0 || random_quote_index > len(Quote_bank) {
		return Quote{}, fmt.Errorf("[random error]")
	}
	return Quote_bank[random_quote_index], nil
}
