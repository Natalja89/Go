package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Println("Введите любое трехзначное число ")
	fmt.Scanln(&a)

	var b = (a / 100)
	var c = (a % 100) / 10
	var d = (a % 10)

	fmt.Println("В числе сотен: " + strconv.Itoa(b) + " десятков: " + strconv.Itoa(c) + " единиц: " + strconv.Itoa(d))
}
