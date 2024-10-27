package main

import (
	"fmt"
)

func main() {
	fmt.Println("Выберите операцию:")
	fmt.Println("1 - Сложение")
	fmt.Println("2 - Вычитание")
	fmt.Println("3 - Умножение")
	fmt.Println("4 - Деление")
	fmt.Println("5 - Выход")

	for {
		fmt.Print("Введите номер операции: ")
		var operation int
		fmt.Scanln(&operation)

		if operation == 5 {
			fmt.Println("До свидания!")
			break
		}

		fmt.Print("Введите первое число: ")
		var num1 float64
		fmt.Scanln(&num1)

		fmt.Print("Введите второе число: ")
		var num2 float64
		fmt.Scanln(&num2)

		if operation == 1 {
			fmt.Println(num1 + num2)
		} else if operation == 2 {
			fmt.Println(num1 - num2)
		} else if operation == 3 {
			fmt.Println(num1 * num2)
		} else if operation == 4 {
			if num2 == 0 {
				fmt.Println("Деление на ноль!")
			} else {
				fmt.Println(num1 / num2)
			}
		}
	}
}
