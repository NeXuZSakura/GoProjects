package main

import (
	"fmt"
	"os"
)

func main() {
	first := 0
	var operation = ""
	second := 0
	fmt.Print("Введите первое число: ")
	ScanInt(&first)
	fmt.Print("Выберите операцию (+, -, *, /,): ")
	ScanStr(&operation)
	fmt.Print("Введите второе число: ")
	ScanInt(&second)
	AllOperation(&first, &operation, &second)
}

func AllOperation(fst *int, op *string, sec *int) {
	res := 0
	switch *op {
	case "+":
		res = *fst + *sec
		fmt.Print("Результат ", *fst, " + ", *sec, " = ", res)
	case "-":
		res = *fst - *sec
		fmt.Print("Результат ", *fst, " - ", *sec, " = ", res)
	case "*":
		res = *fst * *sec
		fmt.Print("Результат ", *fst, " * ", *sec, " = ", res)
	case "/":
		if *sec == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
			os.Exit(4)
		}
		res = *fst / *sec
		fmt.Print("Результат ", *fst, " / ", *sec, " = ", res)
	}
}

func ScanInt(a *int) {
	_, err := fmt.Scan(&*a)
	if err != nil {
		fmt.Println("Некорректное значение. Пожалуйста, введите числовое значение.")
		os.Exit(1)
	}
}

func ScanStr(a *string) {
	_, err := fmt.Scan(&*a)
	if err != nil {
		fmt.Println("Введено некоректное значение!")
		os.Exit(2)
	}
	var b = *a == "*" || *a == "/" || *a == "+" || *a == "-"
	if !b {
		fmt.Println("Некоректная операция. Пожалуйста, используйте +, -, *, / ")
		os.Exit(3)
	}
}