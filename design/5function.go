package design

import "fmt"

type computer struct {
	Maker, CPU, RAM, ROM, Design  string
}
type computerSet func(*computer)
type computerBuilder struct {
	acts []computerSet
}

func Function() {
	b := computerBuilder{}
	showComputer(
		b.maker("Apple").
			cpu("octal-Core").
			ram("16GB").
			rom("1TB").
			design("fantastic design that shape of information").
			builder())
}

func showComputer(com *computer) {
	fmt.Println(com)
}

func (c *computerBuilder) builder() *computer {
	com := computer{}
	for _, act := range c.acts {
		act(&com)
	}
	return &com
}

func (b *computerBuilder) maker(maker string) *computerBuilder{
	b.acts = append(b.acts, func(c *computer) {
		c.Maker = maker
	})
	return b
}

func (b *computerBuilder) cpu(cpu string) *computerBuilder{
	b.acts = append(b.acts, func(c *computer) {
		c.CPU = cpu
	})
	return b
}

func (b *computerBuilder) ram(ram string) *computerBuilder{
	b.acts = append(b.acts, func(c *computer) {
		c.RAM = ram
	})
	return b
}

func (b *computerBuilder) rom(rom string) *computerBuilder{
	b.acts = append(b.acts, func(c *computer) {
		c.ROM = rom
	})
	return b
}

func (b *computerBuilder) design(design string) *computerBuilder{
	b.acts = append(b.acts, func(c *computer) {
		c.Design = design
	})
	return b
}



