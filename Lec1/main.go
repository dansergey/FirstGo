package main

import "fmt"

func main() {
	//Простейший ввывод на консоль. println - это вывод аргумента + '\n' - перенос на новую строчку
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")
	//фУНКЦИЯ PRINT - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
	/////
	////
	/////
	//Декларирование переменных

	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age agter initialixation:", age)

	//
	var height int = 183
	fmt.Println("My heignt is:", height)

	//Полустрогость типизации
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	//
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)

	var (
		personNamse string = "Bob"
		personAge   int    = 42
		personUid   int
	)

	fmt.Printf("Name: %s\nAge %d\nUid: %d\n", personNamse, personAge, personUid)
}
