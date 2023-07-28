package main

import (
	"fmt"
	"strings"
)

func main() {
	// for int; condition; post{
	// 	init - блок инициализции переменных циклаconst
	// 	condition - условие (если верно - то тело цикло выолпняется, если нет - то цикл завершается)
	// 	post - изменение переменной цикла
	// }

	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}

	var i string
	i = "Vova"
	fmt.Println(i)

	//break - команда, прерывающая текущее выполнение тело цикла

	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	//continue - команда, прерывающая текущее выполнение тела цикла и передающая управления СЛЕДУЮЩЕЙ ИТЕРАЦИИ
	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

	//Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идеии треугольник")

	//
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
		}
	}

	// Лейблы
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу чтобы вообще все циклы (Внешние тоже остановились)
			}
		}
	}
	//Модификации цикла for.
	//1. Классчиеский цикл while do
	var loopVar int = 0
	// while (loopVar < 10){
	// 	...
	// 	loopVar++
	// }

	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}

	// 2. Классический бесконечный цикл
	var password string
outer2:
	for {
		fmt.Println("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
			continue
		} else {
			fmt.Println("Password Accepted")
			break outer2
		}
	}

}
