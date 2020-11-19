package main

import (
	"main/app"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app.Start()
}
