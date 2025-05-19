package main

import (
	"fmt"
	"log"
)

func main() {
	v := in()
	fmt.Println("Факториал числа ", v, "= ", fact(v))
	fmt.Println("Факториал числа ", v, " через замыкание = ", callFactClos(v))
	fmt.Println("Факториал числа ", v, " через рекурсию = ", factRec(v))

}
func in() int {
	var val int
	fmt.Print("Введите число, для вычисления факториала: ")
	_, err := fmt.Scan(&val)
	if err != nil {
		log.Fatal("Ошибка при вводе")
	}
	if val < 1 {
		log.Fatal("Введено некорректное значение")
	}
	return val
}
func fact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i
	}
	return result
}

func factClos() func() int {
	res, count := 1, 0
	return func() int {
		count++
		res *= count
		return res
	}
}

func callFactClos(v int) int {
	var out int
	f := factClos()
	for i := 1; i <= v; i++ {
		out = f()
	}
	return out
}

func factRec(v int) int {
	if v == 0 {
		return 1
	}
	return v * factRec(v-1)
}
