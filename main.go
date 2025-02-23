package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight = 1.9
	var userKg float64 = 95
	var BMI = userKg / math.Pow(userHeight, BMIPower)
	fmt.Println(BMI)
}
