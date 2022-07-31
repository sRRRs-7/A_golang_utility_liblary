package design

import "fmt"

type Person struct {
	country, postalCode, city string
	company, occupation, position string
}

type PersonBuilder struct {
	person *Person
}

func Facets() {
	b := constructor()
	b.at("Australia").
		code("123-123").
		in("townsBill").
		comp("freelance").
		occup("software developer").
		pos("CEO")

	b.person.show()
}

// create instance
func constructor() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (p *PersonBuilder) at(country string) *PersonBuilder {
	p.person.country = country
	return p
}

func (p *PersonBuilder) code(postalCode string) *PersonBuilder {
	p.person.postalCode = postalCode
	return p
}

func (p *PersonBuilder) in(city string) *PersonBuilder {
	p.person.city = city
	return p
}

func (p *PersonBuilder) comp(company string) *PersonBuilder {
	p.person.company = company
	return p
}

func (p *PersonBuilder) occup(occupation string) *PersonBuilder {
	p.person.occupation = occupation
	return p
}

func (p *PersonBuilder) pos(position string) *PersonBuilder {
	p.person.position = position
	return p
}

func (p *Person) show() {
	fmt.Println(p.country)
	fmt.Println(p.postalCode)
	fmt.Println(p.city)
	fmt.Println(p.company)
	fmt.Println(p.occupation)
	fmt.Println(p.position)
}