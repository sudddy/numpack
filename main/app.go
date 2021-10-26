package main

import (
	"fmt"

	"github.com/sudddy/numpack/calc"
)

func main() {
	fmt.Println("Calcualtion")
	result := calc.Add(1,2)
	fmt.Println("result, ", result)
}