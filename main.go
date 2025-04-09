package main

import (
	"fmt"
	"math/rand"
	"time"
)

// snippet: main
func main() {
	HelloWorld()
	fmt.Println("2 + 3 =", Add(2, 3))
	fmt.Println("Random number:", RandomNumber())
}

// snippet: main

func HelloWorld() {
	fmt.Println("Hello, world!")
}

// snippet: add
func Add(a int, b int) int {
	return a + b
}

// snippet: add

// snippet: randomNumber
func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

// snippet: randomNumber
