package generatevalue

import (
	"math/rand"
	"time"
)

func GenerateValue() int {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100)

	return a
}
