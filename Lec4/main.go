package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	//Boolean операции
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	//Numerics. Integers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b", b, "Sum of a+b:", a+b)
	//Вывод типа через %T форматирования
	fmt.Printf("Type is %T\n", a)
	//Узнаём,  сколько байт занимает переменная типа int'
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент. При использовании короткого объявления - тип для целого числа - int платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент 2. Используйте явное приведение типов при необходимости, если уверены что не произойдёт коллиззии
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперимент 3.
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64 + int64(fourthInt))

	//Float
	//float32, float64
	floatFirst, floatSecons := 5.67, 12.54
	fmt.Printf("type of a %T and typr of %T\n", floatFirst, floatSecons)

}
