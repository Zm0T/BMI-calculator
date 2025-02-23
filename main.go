package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight float64 = 1.8
	var userKg float64 = 100.0
	fmt.Print("Калькулятор индекса массы тела ___ \n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	BMI := userKg / math.Pow(userHeight, BMIPower)
	fmt.Print("Ваш индекс массы телы: ")
	fmt.Print(BMI)
}
