package random

import (
	"math/rand"
	"time"
)

func (rn *RandomNumber) GetRandom(num int) int {
	if num <= 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(num)
}
