package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Println("Введите чему равна площадь круга")
	fmt.Scanln(&s)

	var L float64 = (s * 4 * math.Pi)
	var l float64 = math.Sqrt(L)

	fmt.Println("Длина окружности равна " + fmt.Sprint(l))
	fmt.Println("Диаметр окружности равен " + fmt.Sprint(l/math.Pi))
}
