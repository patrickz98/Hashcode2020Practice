package main

import (
	"fmt"
	"github.com/patrickz98/Hashcode2020Quali/src/order"
)

func main() {

	inputFile := "b_small"
	constrains := order.New(inputFile)

	fmt.Println("SlicesMaximum", constrains.SlicesMaximum)
	fmt.Println("Slices", constrains.Slices)

	score := constrains.Submission([]int{0, 1})

	fmt.Println("Score", score)
}
