package simplefunctions

import "fmt"

func Hello() {

	fmt.Println("hello world")
	return
}

func Add(i, i2 int) int {

	Hello()
	println("i value is ", i)
	return Sum(100, 900)
}

func Sum(i int, j int) int {

	Hello()
	fmt.Println(i + j)
	return i + j
}
