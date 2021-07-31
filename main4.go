package main

import (
	"fmt"
)

const pi = 3.1415

func main() {
	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Радиус круга: %d см\n", circleRadius)
	fmt.Println("Формула для расчета площади круга: S=pr2")

	fmt.Printf("Площадь круга: %.2f см. кв\n", circleArea)
}