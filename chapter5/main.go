package main

import "fmt"

func main() {
	// mass()
	// mass_2()
	// slice()
	// slice_2()
	// slice_3()
	// map_1()
	map_2()
	// map_3()
	// exercise_1()
	// exercise_2()
	// map_4()
}

func mass() {
	// Массив — это нумерованная последовательность элементов одного типа с фиксированной длиной.
	var x [5]int
	x[2] = 100
	fmt.Println(x)
}

func mass_2() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0

	// value будет тем же самым что и x[i].
	// Мы использовали ключевое слово range перед переменной, по которой мы хотим пройтись циклом.

	// Одиночный символ подчеркивания _ используется, чтобы сказать компилятору,
	// что переменная нам не нужна (в данном случае нам не нужна переменная итератора).

	for _, value := range x { //
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func slice() {
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[:3]

	fmt.Println(x)
}

func slice_2() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func slice_3() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func map_1() {
	//Карта (также известна как ассоциативный массив или словарь) — это неупорядоченная коллекция пар вида ключ-значение.

	//Карта представляется в связке с ключевым словом map, следующим за ним
	// типом ключа в скобках и типом значения после скобок.
	//Читается это следующим образом: «x — это карта string-ов для int-ов».

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 11
	fmt.Println(y[1])
}

func map_2() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements["Li"])
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}

func map_3() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func map_4() {
	x := make(map[int]int)
	x[1] = 10
	x[9] = 23
	// delete(x, 1)
	fmt.Println(x[9])
}

func exercise_1() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])

	y := make([]int, 3, 9)
	fmt.Println(len(y))
}

func exercise_2() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := 1000
	for i := 0; i < len(x); i++ {
		if min > x[i] {
			min = x[i]
		}
	}
	fmt.Println(min)
}
