package main

import (
	"fmt"
	"github.com/patrickz98/Hashcode2020Quali/src/order"
	"sort"
)

func containsInt(array []int, search int) bool {
	i := sort.SearchInts(array, search)
	return i < len(array) && array[i] == search
}

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

	sort.Ints(indices)
	// fmt.Println("indices", indices)

	gains := make(map[int][]int)
	big := make(map[int]int)

	left := constrains.SlicesMaximum - constrains.Score(indices)

	stopIndex := -1

	for inz, index := range indices {

		if stopIndex > 0 {
			break
		}

		item := constrains.Slices[index]

		// fmt.Printf("%d/%d\r", len(indices), count)

		for inx, valX := range constrains.Slices {

			if stopIndex > 0 {
				break
			}

			if containsInt(indices, inx) {
				continue
			}

			for iny, valY := range constrains.Slices {

				if inx == iny {
					continue
				}

				if item >= valX+valY {
					continue
				}

				gain := (valX + valY) - item

				if gain > left {
					continue
				}

				if containsInt(indices, iny) {
					continue
				}

				if big[inz] < valX+valY {
					gains[inz] = []int{inx, iny}
					big[inz] = valX + valY

					fmt.Printf("index: %6d, item: %6d, new: %6d, inx: %4d, iny: %4d, gain: %3d\n",
						index, item, valX+valY, inx, iny, gain)

					if gain == left {
						stopIndex = inz
						break
					}
				}
			}
		}
	}

	if stopIndex > 0 {
		val := gains[stopIndex]
		indices[stopIndex] = val[0]
		indices = append(indices, val[1])
	}

	sort.Ints(indices)

	fmt.Println("gains", gains)

	score := constrains.Submission(indices)
	fmt.Println("Score", score)
	fmt.Println("Missing", constrains.SlicesMaximum-score)
}

func mainn() {

	inputFiles := []string{
		"a_example",
		"b_small",
		"c_medium",
		"d_quite_big",
		"e_also_big",
	}

	scoreTotal := 0

	for _, file := range inputFiles {

		fmt.Println("File", file)
		constrains := order.New(file)

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

		scoreTotal += constrains.SlicesMaximum - score
	}

	fmt.Println("Total Score", scoreTotal)
}
