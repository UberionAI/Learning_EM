package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: app <name>")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Println("hello,", name)
}

//Задание 8
//Пусть будут два коммента от юзеров по начальной версии программы:
//-При отсутствии аргумента программа паникует с непонятным runtime‑stackом.
//-Нет сообщений пользователю
// Исправим данные ошибки
//старая(проблемная версия):
//func main() {
//	arg := os.Args[1]
//	_ = arg
//}
