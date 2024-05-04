package main

import (
	"fmt"
	"module/solve"
	"os"
	"os/exec"
)

func main() {

	// в задании не было сказано делать удобный интерфейс для ввода значений в функции из консоли, поэтому менять в коде

	// выбор заданий
	var input string
	fmt.Println("введите '1' для задания 1")
	fmt.Println("введите '2' для задания 2")
	fmt.Println("введите '3' для задания 3")
	fmt.Println("введите '4' для задания 4")
	fmt.Scan(&input)

	// очистка консоли после выбора
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	switch input {
	case "1":
		// склонение компьютеров
		fmt.Println(solve.Task_1(1310))
	case "2":
		// вывод общих делителей у чисел
		fmt.Println(solve.Task_2([]int{42, 12, 18}))
	case "3":
		// вывод простых чисел из диапазона
		fmt.Println(solve.Task_3(11, 20))
	case "4":
		// вывод таблицы умножения
		solve.Task_4(5)
	}

	// ожидание кнопки для закрытия программы
	fmt.Scan(&input)

}
