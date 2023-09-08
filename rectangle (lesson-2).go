package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b int

	fmt.Println("Введите длину меньшей стороны")
	fmt.Scanln(&a)

	fmt.Println("Введите длину большей стороны")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника равна " + strconv.Itoa(a*b))
}
