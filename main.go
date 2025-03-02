package main

import (
	"errors"
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	var userInput string
	for {
		userHeight, userKg := getUserInput()
		IMT, err := calculateBMI(userKg, userHeight, BMIPower)
		if err != nil {
			fmt.Println("Неверно введены значения")
			continue
		} else {
			outputResult(IMT)
		}

		fmt.Println("Хотите ли вы продолжить? Введите y-да, n-нет")
		fmt.Scan(&userInput)
		if userInput == "n" {
			fmt.Println("Программа завершила свою работу")
			break
		}
	}
	//if IMT < 16 {
	//	fmt.Println("У вас сильный недостаток веса")
	//} else if IMT >= 16 && IMT < 18.5 {
	//	fmt.Println("Дефицит массы тела")
	//} else if IMT >= 18.5 && IMT < 25 {
	//	fmt.Println("Нормальный вес")
	//} else if IMT >= 25 && IMT < 30 {
	//	fmt.Println("Избиточный вес")
	//} else {
	//	fmt.Println("Степень ожирения")
	//}
}

func outputResult(BMI float64) {
	switch {
	case BMI < 16:

		fmt.Println("У вас сильный недостаток веса")
	case BMI < 18.5:

		fmt.Println("Дефицит массы тела")
	case BMI < 25:

		fmt.Println("Нормальный вес")
	case BMI < 30:
		fmt.Println("Избиточный вес")
	default:
		fmt.Println("Степень ожирения")
	}
	fmt.Printf("Ваш индекс массы телы: %.0f\n", BMI)
}

func calculateBMI(userKg float64, userHeight float64, BMIPower float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("INPUT_ERROR")
	}
	BMI := userKg / math.Pow(userHeight/100, BMIPower)
	return BMI, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64 = 1.8
	var userKg float64 = 100.0

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scanln(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scanln(&userKg)

	return userHeight, userKg
}
