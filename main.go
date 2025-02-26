package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	userHeight, userKg := getUserInput()
	fmt.Println("___ Калькулятор индекса массы тела ___")

	var IMT float64 = calculateBMI(userKg, userHeight, BMIPower)

	outputResult(IMT)
}

func outputResult(BMI float64) {
	fmt.Printf("Ваш индекс массы телы: %.0f", BMI)
}

func calculateBMI(userKg float64, userHeight float64, BMIPower float64) float64 {
	BMI := userKg / math.Pow(userHeight/100, BMIPower)
	return BMI
}

func getUserInput() (float64, float64) {
	var userHeight float64 = 1.8
	var userKg float64 = 100.0

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}
