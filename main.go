package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вычислитель квадрата целого числа v.1.1")
	fmt.Println("Приветствую Вас! Пожалуйста, представьтесь:")

	var username string
	var x int

	fmt.Scan(&username)

	fmt.Print(username, ", прошу Вас, ниже введите целое число, квадрат которого вы хотите получить:\n")

	fmt.Scan(&x)
	squareResult := x * x

	fmt.Print(username, ", квадрат вашего числа ", x, " равен ", squareResult, ".\n")
}
