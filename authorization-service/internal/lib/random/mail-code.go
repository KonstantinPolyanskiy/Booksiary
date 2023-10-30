package random

import "math/rand"

func Code() int {
	return 1000 + rand.Intn(10000-1000+1)
}
