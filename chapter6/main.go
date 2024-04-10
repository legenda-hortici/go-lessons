package main

import "fmt"

// Функция начинается с ключевого слова func, за которым следует имя функции.
// Аргументы (входы) определяются так: имя тип, имя тип, ….
// Наша функция имеет один параметр (список оценок) под названием xs.
// За параметром следует возвращаемый тип.
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Переменное число аргументов

// Использование ... перед типом последнего аргумента означает, что функция может содержать
// ноль и более таких параметров. В нашем случае мы берем ноль и более int.
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Замыкание
// add является локальной переменной типа
// func(int, int) int (функция принимает два аргумента типа int и возвращает int).
func closures() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}

// Функцию, использующую переменные, определенные вне этой функции, называют замыканием.
// В нашем случае функция increment и переменная x образуют замыкание.

func closures_2() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// output: 2
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Рекурсия

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// В Go есть специальный оператор defer, который позволяет отложить вызов
// указанной функции до тех пор, пока не завершится текущая.
// Эта программа выводит 1st, затем 2nd. Грубо говоря defer перемещает вызов second в конец функции
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

// func main() {
// 	defer second()
// 	first()
// }

func half(x int) (int, bool) {
	r := x / 2
	return r, x%2 == 0
}

func maxElem(args ...int) int {
	max := 0
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (r uint) {
		r += i
		i += 2
		return
	}
}

func fib(n int) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(maxElem(1, 3, 2))
	oddNum := makeOddGenerator()
	fmt.Println(oddNum())
	fmt.Println(fib(10))
}
