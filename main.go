package main

import (
	"math/rand"
	"time"

	"gotest/universe"
	"gotest/ark"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	universe.Test()
	ark.Test()
}
