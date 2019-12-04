package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 88
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 44})

	fmt.Println(NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}

/**
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 44}
&{Jon 88}
Sean
50
51

*/
