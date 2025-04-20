package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	userChoise := choiseOperation()
	userNumbers := selectNumbers()
	x := []float64{}
	for _, value := range userNumbers {
		valueFloat, _ := strconv.ParseFloat(value, 64)
		x = append(x, valueFloat)
	}
	res := result(userChoise, x)
	if res == 0 {
		fmt.Println("Вы выбрали не существующую операцию")
	} else {
		fmt.Println(res)
	}
}

func choiseOperation() string {
	var userChoise string
	fmt.Print("Выберете операцию (AVG, SUM, MED): ")
	fmt.Scan(&userChoise)
	return userChoise
}

func selectNumbers() []string {
	var userNumbers string
	fmt.Println("Введите числа через запятую (Например: 1,4,29)")
	fmt.Scan(&userNumbers)
	numbers := strings.Split(userNumbers, ",")
	return numbers
}

func result(oper string, x []float64) float64 {
	var sum float64
	switch oper {
	case "AVG":
		for _, value := range x {
			sum += value
		}
		res := sum / float64(len(x))
		return res
	case "SUM":
		for _, value := range x {
			sum += value
		}
		return sum
	case "MED":
		var res float64
		sort.Float64s(x)
		if len(x)%2 == 0 {
			index := len(x) / 2
			res = (x[index] + x[index-1]) / 2
			return res
		} else {
			index := len(x) / 2
			res := x[index]
			return res
		}
	default:
		return 0
	}
}
