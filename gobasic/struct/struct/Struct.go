package main

import "fmt"

type PersonMethods interface {
	haveBirthDay()
}

type Person struct {
	Name       string
	Age        int
	Hobby      []string
	InLoveWith string
}

// constructor
func NewPerson(name string, age int, hobby []string, inLoveWith string) *Person {

	if name == "" {
		name = "Emma Watson"
	}

	if len(hobby) == 0 {
		hobby = []string{"acting", "taking photos"}
	}

	if inLoveWith == "" {
		inLoveWith = "Mujibul Haque Tanim"
	}
	return &Person{
		Name:       name,
		Age:        age,
		Hobby:      hobby,
		InLoveWith: inLoveWith,
	}
}

// method
func (p *Person) haveBirthDay() {
	p.Age++
}

func main() {
	p1 := NewPerson("Millie bobby brown", 20, []string{}, "")

	p1.haveBirthDay()
	//as NewPerson returns a pointer, we can use *p1 to access the value (dereferencing)
	fmt.Println(*p1)
}
