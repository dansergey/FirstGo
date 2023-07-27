package main

import "fmt"

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	//Boolean операции
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
}
