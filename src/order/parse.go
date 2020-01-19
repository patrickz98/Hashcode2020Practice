package order

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Constrains struct {
	inputFile     string
	SlicesMaximum int
	Size          int
	Slices        []int
}

const submissonPath = "submisson"

func (cons Constrains) Score(indices []int) int {

	sum := 0

	for _, index := range indices {
		sum += cons.Slices[index]
	}

	return sum
}

func (cons Constrains) Submission(indices []int) int {

	sum := 0

	submissonTxt := strconv.Itoa(len(indices)) + "\n"
	indicesStr := make([]string, len(indices))

	integrity := make(map[int]bool)

	for inx, index := range indices {
		sum += cons.Slices[index]
		indicesStr[inx] = strconv.Itoa(index)

		if !integrity[index] {
			integrity[index] = true
		} else {
			panic("index: " + strconv.Itoa(index))
		}
	}

	submissonTxt += strings.Join(indicesStr, " ") + "\n"

	err := os.MkdirAll(submissonPath, 0755)
	check(err)

	outFile := submissonPath + "/" + cons.inputFile + ".out"
	err = ioutil.WriteFile(outFile, []byte(submissonTxt), 0755)
	check(err)

	return sum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func New(inputFile string) Constrains {

	bytes, err := ioutil.ReadFile("input/" + inputFile + ".in")
	check(err)

	input := strings.Split(string(bytes), "\n")
	constrains := strings.Split(input[0], " ")

	slicesMaximum, err := strconv.Atoi(constrains[0])
	check(err)

	typesOfPizza, err := strconv.Atoi(constrains[1])
	check(err)

	pizzaParts := strings.Split(input[1], " ")

	pizzaSlices := make([]int, typesOfPizza)

	for inx, pizzaStr := range pizzaParts {

		pizzaSlice, err := strconv.Atoi(pizzaStr)
		check(err)

		pizzaSlices[inx] = pizzaSlice
	}

	return Constrains{
		inputFile:     inputFile,
		SlicesMaximum: slicesMaximum,
		Size:          typesOfPizza,
		Slices:        pizzaSlices,
	}
}
