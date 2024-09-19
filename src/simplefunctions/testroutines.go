package simplefunctions

import "fmt"

func Hello() {
	fmt.Println("Hello World")
	return
}

func Add(i, y int) int {
	return Sum(100, 100)
}

func Sum(i int, j int) int {
	return i + j
}

func Calc(i int) (int, int) {
	p := i * i
	s := i + i
	return p, s
}

func Execute() int {
	Hello()
	fmt.Println("Sum  is ", Sum(1001, 1001))
	fmt.Println("Add is ", Add(
		1000,
		900,
	))
	fmt.Print("Calc is ")
	fmt.Println(Calc(10))
	return 1
}
