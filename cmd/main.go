package main

import (
	"fmt"
	"os"

	day2 "github.com/asahnoln/advent-go-day2"
)

func main() {
	sub := day2.NewSub()

	sub.ParseInstructions(os.Stdin)

	fmt.Printf("Multiply answer: %d", sub.Multiply())
}
