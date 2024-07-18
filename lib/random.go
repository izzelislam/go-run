package lib

import (
	"math/rand"
	"time"
)

func Random() {
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	for i := 0; i < 100; i++ {
		println(random.Intn(100))
	}
	
}