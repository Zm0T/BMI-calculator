package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.9
	var userKg float64 = 95
	var BMI = userKg / math.Pow(userHeight, 2)
	fmt.Println(BMI)
}
