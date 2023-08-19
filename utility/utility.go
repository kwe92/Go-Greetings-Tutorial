package utility

import (
	"math/rand"
)

func RandInt(exclusive_range int) (random_integer int) {
	return rand.Intn(exclusive_range)
}
