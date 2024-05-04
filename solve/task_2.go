package solve

import "sort"

// пустая структура для реализации set
type Emptys struct{}

// задание про общий делитель
func Task_2(mass []int) []int {

	if len(mass) == 0 {
		return []int{}
	}

	// создаем словарь (set). и записываем длину первого значения
	dictionary := make(map[int]Emptys)
	var firstLen int = mass[0] // если не считать само число делителем, то надо поделить на 2  !!!!!!!!

	// начальное заполнение словаря. заполняем все делители первого числа.
	for i := 2; i <= firstLen; i++ {
		if mass[0]%i == 0 {
			dictionary[i] = Emptys{}
		}
	}

	// перебираем по каждому числу кроме первого, и удаляем делители из словаря если не делится.
	for i := 1; i < len(mass); i++ {

		// проверяем на делимости до firstlen, что удалять делители которые больше этого числа
		for j := 2; j <= firstLen; j++ {
			if mass[i]%j != 0 {
				delete(dictionary, j)
			}
		}

		// delete(dictionary, mass[i]) // разкомментировать если не брать в делители себя же  !!!!!!!!
	}

	// переводим значения из словаря в слайс и выводим
	result := make([]int, 0, len(dictionary))
	for i, _ := range dictionary {
		result = append(result, i)
	}

	// словарь содержит неупорядоченный набор значений, поэтому в массиве тоже неупорядоченно
	sort.Ints(result)
	return result
}
