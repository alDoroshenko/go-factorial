package main

import (
	"fmt"
	"log"
)

func main() {
	v := in()
	fmt.Println("Факториал числа ", v, "= ", fact(v))
	f := factClos()
	var out int
	for i := 1; i <= v; i++ {
		out = f()
	}
	fmt.Println("Факториал числа ", v, " через замыкание= ", out)

}
func in() int {
	var val int
	fmt.Print("Введите число, для вычисления факториала: ")
	fmt.Scan(&val)
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
