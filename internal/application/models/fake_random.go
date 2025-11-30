package models

import "math/rand"

type FakeRandom struct {
	R *rand.Rand
}

func (s *FakeRandom) GetRandom(n int) int {
	if n <= 0 {
		return 0
	}
	return s.R.Intn(n)
} // пришлось написать сюда еще рандом,
// чтобы не импортировать его из inf, если
// есть другие решения, можно в фидбек
// было бы интересно почитать
