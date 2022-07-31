package factory

import "fmt"

type person struct {
	Name string
	Age int
	Money int
}

type happyPerson struct {
	person person
}

type crazyPerson struct {
	person person
}

type Person interface {
	Greet()
}


func FactoryDesign() {
	p := constructor("michel", 20, 100)
	p.Greet()

	p = constructor("john", 30, 30)
	p.Greet()

	p = constructor("debt", 30, -100)
	p.Greet()
}

func constructor(name string, age, money int) Person {
	p := person{
		Name: name,
		Age: age,
		Money: money,
	}
	if money > 99 {
		return &happyPerson{
			person: p,
		}
	} else if money < 0 && money < 99 {
		return &p
	} else {
		return &crazyPerson{
			person: p,
		}
	}
}

func (p *person) Greet() {
	fmt.Printf("I'm %s. I have %d age. I'm normal \n", p.Name, p.Age)
}

func (p *happyPerson) Greet() {
	fmt.Printf("I'm %s. I have %d age. I'm happy \n", p.person.Name, p.person.Age)
}

func (p *crazyPerson) Greet() {
	fmt.Printf("I'm %s. I have %d age. I'm crazy \n", p.person.Name, p.person.Age)
}