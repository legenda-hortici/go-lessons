package main

import "fmt"

func main() {
	fizz_buzz()
	// devider()
	// switch_structure()
	// fmt.Println()
	// if_structure()
	// fmt.Println()
	// for_structure()
}

// Оператор for даёт возможность повторять список инструкций (блок) определённое количество раз.

func for_structure() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Оператор if аналогичен оператору for в том, что он выполняет блок в зависимости от условия.

func if_structure() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

// Переключатель начинается с ключевого слова switch, за которым следует выражение (в нашем случае i) и серия возможных значений (case).

func switch_structure() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)

	switch input {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("First")
	case 2:
		fmt.Println("Second")
	default:
		fmt.Println("Unknown number")
	}
}

// Exercise 1

func devider() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizz_buzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
