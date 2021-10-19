package main

import (
	"math/rand"
	"time"

	"un/universe"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	universe.Test()
}
