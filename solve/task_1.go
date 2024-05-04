package solve

import (
	"fmt"
	"strconv"
)

// задание про компьютеры
func Task_1(num int) string {

	var stroka string = strconv.Itoa(num)
	var result string

	// получаем последнюю цифру из числа, и склонение зависит от неё
	var oneChar, _ = strconv.Atoi(stroka[len(stroka)-1:])
	switch oneChar {
	case 0:
		result = "компьютеров"
	case 1:
		result = "компьютер"
	case 2, 3, 4:
		result = "компьютера"
	case 5, 6, 7, 8, 9:
		result = "компьютеров"
	}

	// если число больше 10, то если последние 2 цифры в диапазоне от 10 до 20, склонение будет изменено.
	if num > 10 {
		var twoChar, _ = strconv.Atoi(stroka[len(stroka)-2:])
		if twoChar > 10 && twoChar < 20 {
			result = "компьютеров"
		}
	}

	// вывод ответа
	return fmt.Sprintf("%d %s", num, result)
}
