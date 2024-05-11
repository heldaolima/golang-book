package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value") // maravilhoso
	flag.Parse()

	fmt.Println(rand.Intn(*maxp))
}
