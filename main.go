package main

import (
	"math/rand"
	"time"

	"gotest/universe"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	universe.Test()
}
