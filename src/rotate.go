package src

import "fmt"

func Quote_rotation() (Quote, error) {
	for i := 0; i < len(Quote_bank); i++ {
		if !Quote_bank[i].Used {
			Quote_bank[i].Used = true
			return Quote_bank[i], nil
		} else {
			i++
			Quote_bank[i].Used = false
		}
	}
	return Quote{}, fmt.Errorf("[rotation error]")
}
