package main

import (
	"fmt"
	"main/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	os.Chdir("/Users")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Direcotry: %s\n", newDir)
	rand.Seed(time.Now().UnixNano())
	app.Start()
}
