package main

import (
	"fmt"
	"iter"
)

// countTo Пример простейшей функции-итератора. Возвращаем одно значение типа int (iter.Seq[int])
func countTo(s int) iter.Seq[int] {
	// функция-итератор должна ПРИ КАЖДОЙ ИТЕРАЦИИ возвращать такое значение:
	return func(yield func(int) bool) {
		// описываем свой итератор (цикл или что-то типа того)
		for i := 0; i <= s; i++ {
			// Здесь производится возврат значения итератора: yield(i).
			// Если он false, выходим из функции. (yield приобретает значение false при достижении конца перебора или вручную)
			if !yield(i) {
				return
			}
		}

	}
}

func main() {
	//nums := []int{0, 1, 2, 3, 4, 5}
	// вызов функции-итератора (countTo) в рамках range
	for v := range countTo(5) {
		fmt.Println(v)
	}
}
