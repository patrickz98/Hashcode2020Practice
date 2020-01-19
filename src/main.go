package main

import (
	"fmt"
	"github.com/patrickz98/Hashcode2020Quali/src/order"
)

func main() {

	inputFile := "a_example"
	// inputFile := "b_small"
	// inputFile := "c_medium"
	// inputFile := "d_quite_big"
	// inputFile := "e_also_big"

	constrains := order.New(inputFile)

	fmt.Println("SlicesMaximum", constrains.SlicesMaximum)
	// fmt.Println("Slices", constrains.Slices)

	sum := 0
	indices := make([]int, 0)

	for inx := 0; inx < constrains.Size; inx++ {

		item := constrains.Slices[constrains.Size-inx-1]

		if sum+item < constrains.SlicesMaximum {
			sum += item
			indices = append(indices, constrains.Size-inx-1)
		}
	}

	score := constrains.Submission(indices)
	fmt.Println("Score", score)
	fmt.Println("Missing", constrains.SlicesMaximum-score)
}
