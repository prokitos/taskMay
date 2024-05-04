package solve

import "math"

// проверка что число простое
func primeNumber(num int) bool {

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// вывод массива простых чисел
func Task_3(start int, end int) []int {

	// проверяем каждое число из диапазона, и если оно простое то добавляем в слайс.
	result := make([]int, 0, (end-start)/3)
	for i := start; i <= end; i++ {
		if primeNumber(i) {
			result = append(result, i)
		}
	}

	return result
}
