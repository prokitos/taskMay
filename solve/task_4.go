package solve

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// таблица умножения
func Task_4(max int) {

	// использование tabwriter для уменьшения размера табуляции, чтобы таблица выглядела лучше
	fmt.Println()
	var tabLen int = 4
	w := tabwriter.NewWriter(os.Stdout, tabLen, 0, 0, ' ', tabwriter.AlignRight)

	// выводим строку начальных значений (1 2 3 4 5...)
	fmt.Fprint(w, " ", "\t")
	for j := 1; j <= max; j++ {
		fmt.Fprint(w, j, "\t")
	}
	fmt.Fprintln(w)

	// вывод самой таблицы
	for i := 1; i <= max; i++ {
		// выводим столбец начальных значений (1 2 3 4 5...)
		fmt.Fprint(w, i, "\t")
		for j := 1; j <= max; j++ {
			// выводим сами значения в таблице
			fmt.Fprint(w, i*j, "\t")
		}
		fmt.Fprintln(w)
	}

	// вывод
	w.Flush()

}
