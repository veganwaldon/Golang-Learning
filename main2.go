package main

import (
	"fmt"
)

var globalVar string = "Я глобальная переменная!"

func main() {
	localVar := "а я локальная, обьявлена внутри функции main()"

	fmt.Println(globalVar)
	fmt.Println(localVar)
}