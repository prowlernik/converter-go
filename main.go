package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("__Калькулятор валют__")
	for {
		value := getUserInputValue()
		num := getUserInputNum()
		value_end := getUserInputValueEnd(value)
		output(value, num, value_end)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}

}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Println("")
	fmt.Print("Хотите повторить операцию (y/n) ")
	fmt.Scan(&userChoise)
	fmt.Println("")
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false

}

func getUserInputValueEnd(value string) string {
	for {
		usd, rub, eur := "usd", "rub", "eur"
		var userInput string
		fmt.Printf("Введите конечную валюту (%s,%s,%s) : ", usd, eur, rub)
		fmt.Scan(&userInput)
		if value == userInput {
			fmt.Println("Вы уже выбрали данную валюту")
			continue
		}
		res, err := checkValue(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return res

	}

}

func getUserInputValue() string {
	for {
		usd, rub, eur := "usd", "rub", "eur"
		var userInput string
		fmt.Printf("Введите исходную валюту (%s,%s,%s) : ", usd, eur, rub)
		fmt.Scan(&userInput)
		res, err := checkValue(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return res
	}

}

func checkValue(userInput string) (string, error) {
	if userInput != "usd" && userInput != "rub" && userInput != "eur" {
		return "", errors.New("Неверная валюта")
	}
	return userInput, nil
}

func getUserInputNum() float64 {
	var userInput float64
	for {
		fmt.Print("Введите число :")
		fmt.Scan(&userInput)
		userInput, err := checkNum(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return userInput
	}
}

func checkNum(userInput float64) (float64, error) {
	if userInput <= 0 {
		return 0, errors.New("Неверное число")
	}
	return userInput, nil
}

func output(value string, num float64, value_end string) {
	const USD_TO_EUR, USD_TO_RUB, EUR_TO_USD, EUR_TO_RUB, RUB_TO_EUR, RUB_TO_USD float64 = 0.88, 81.99, 1.14, 93.23, 0.011, 0.012
	switch {
	case value == "usd" && value_end == "eur":
		sum := num * USD_TO_EUR
		fmt.Printf("Результат: %.2f\n", sum)
	case value == "usd" && value_end == "rub":
		sum := num * USD_TO_RUB
		fmt.Printf("Результат: %.2f\n", sum)
	case value == "eur" && value_end == "usd":
		sum := num * EUR_TO_USD
		fmt.Printf("Результат: %.2f\n", sum)
	case value == "eur" && value_end == "rub":
		sum := num * EUR_TO_RUB
		fmt.Printf("Результат: %.2f\n", sum)
	case value == "rub" && value_end == "eur":
		sum := num * RUB_TO_EUR
		fmt.Printf("Результат: %.2f\n", sum)
	case value == "rub" && value_end == "usd":
		sum := num * RUB_TO_USD
		fmt.Printf("Результат: %.2f\n", sum)
	}
}
