package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func (p Person) Great() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.Name, p.Age)
}
