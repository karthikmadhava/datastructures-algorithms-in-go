package DataStructures

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Init() []Person {
	var list = []Person{
		{"Adam", 28},
		{"Eve", 26},
		{"Son", 10},
		{"Daughter", 5},
	}
	return list
}

func AddPerson(member Person, list []Person) []Person {
	list = append(list, member)
	for _, person := range list {
		fmt.Println(person.Name, person.Age)
	}
	return list
}
