package person

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func (p *Person) Display() {
	fmt.Println("name: ", p.name, "\nage: ", p.age)
}

func (p *Person) Birthday() {
	p.age++
}
