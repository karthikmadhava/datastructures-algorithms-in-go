package simplefunctions

import (
	"fmt"
	"main/DataStructures"
)

func Hello() {
	fmt.Println("Hello World")
	return
}

func processArray() {
	DataStructures.AddPerson(DataStructures.Person{
		"su",
		1,
	}, DataStructures.Init())
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

	processArray()
	//binary search
	//arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	//target := 7
	//var result = DataStructures.BinarySearch(arr, target)
	//if result != -1 {
	//	fmt.Printf("Element found at index %d.\n", result)
	//} else {
	//	fmt.Println("Element not found in the array.")
	//}
	return 1
}
