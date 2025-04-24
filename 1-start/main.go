package main

import (
	"errors"
	"fmt"
	"strings"
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
	usd, rub, eur := "usd", "rub", "eur"
	for {
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
	usd, rub, eur := "usd", "rub", "eur"
	for {
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
	value_s := map[string]float64{
		"USD_TO_EUR": 0.88,
		"USD_TO_RUB": 81.99,
		"EUR_TO_USD": 1.14,
		"EUR_TO_RUB": 93.23,
		"RUB_TO_EUR": 0.011,
		"RUB_TO_USD": 0.012,
	}
	res := strings.ToUpper(value) + "_TO_" + strings.ToUpper(value_end)
	result := value_s[res] * num
	fmt.Println(result)
}
