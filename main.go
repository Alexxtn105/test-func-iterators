package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"iter"
	"slices"
)

//region Простейший пример

// countTo Пример простейшей функции-итератора. Возвращаем одно значение типа int (iter.Seq[int])
func countTo(s int) iter.Seq[int] {
	// функция-итератор должна ПРИ КАЖДОЙ ИТЕРАЦИИ возвращать такое значение:
	return func(yield func(int) bool) {
		// описываем свой итератор (цикл или что-то типа того)
		for i := 1; i <= s; i++ {
			// Здесь производится возврат значения итератора: yield(i).
			// Если он false, выходим из функции. (yield приобретает значение false при достижении конца перебора или вручную)
			if !yield(i) {
				return
			}
		}
	}
}

//endregion

// region Пример со слайсами (с дженериками)

// Backwards Функция обратного перебора элементов слайса
func Backwards[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		// пишем собственную реализацию итератора
		for i := len(s) - 1; i >= 0; i-- {
			// возвращаем значение итератор
			if !yield(s[i]) {
				return
			}
		}
	}
}

//endregion

// region Пример для функций, принимающих итераторы
func PrintElement[V any](s iter.Seq[V]) {
	for v := range s {
		fmt.Println(v)
	}
}

//endregion

// region Примеры дженериков (comparable и constraints.Ordered)
func testComparable[s comparable](v s) s {
	return v
}

func testOrdered[s constraints.Ordered](v s) s {
	return v
}

//endregion

func main() {
	//region Вызов простейшей функции-итератора (countTo)
	fmt.Println("---------------------------------------------")
	fmt.Println("Вывод простейшей функции-итератора - countTo:")
	// вызываем функцию-итератор
	for v := range countTo(5) {
		fmt.Println(v)
	}
	//endregion

	//region Вызов функции-итератора для слайса (Backwards)
	fmt.Println("----------------------------------------------")
	fmt.Println("Вывод функции-итератора для слайса - Backwards")
	nums := []int{1, 2, 3, 4, 5}
	// 5, 4, 3, 2, 1
	// вызываем функцию-итератор для слайса
	for v := range Backwards(nums) {
		fmt.Println(v)
	}

	//endregion

	//region Вызов функции, принимающей итератор (PrintElement)
	fmt.Println("----------------------------------------------")
	fmt.Println("Вывод функции, принимающей итератор - PrintElement")
	PrintElement(Backwards(nums))

	//endregion

	//region Встроенные функции-итераторы (в пакете slices)
	fmt.Println("----------------------------------------------")
	fmt.Println("Вывод функции, принимающей итератор - PrintElement")
	for _, v := range slices.Backward(nums) {
		fmt.Println(v)
	}
	//endregion
}
