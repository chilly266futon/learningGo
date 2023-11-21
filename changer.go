package main

import "fmt"

var Global = 5

func UseGlobal() {
	defer func(checkout int) {
		Global = checkout // присваиваем переменной Global значение аргумента
	}(Global) // копируем значение Global в аргументы функции
	Global = 42 // меняем значение Global
	fmt.Println(Global)
	// Здесь будет вызвана отложенная функция
}

//func main() {
//	fmt.Println(Global)
//	UseGlobal()
//	fmt.Println(Global)
//}
