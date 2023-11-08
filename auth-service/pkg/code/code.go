package code

import (
	"math/rand"
	"time"
)

// Code возвращает 4х значный код
func Code() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(10000-1000+1)
}
