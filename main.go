package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2

	userHeight := 1.8
	userWeight := 67.0

	// Body Mass Index
	BMI := userWeight / math.Pow(userHeight, BMIPower)

	fmt.Printf("%f\n", BMI)

	fmt.Print(BMI)
	fmt.Printf("Test2")
}
