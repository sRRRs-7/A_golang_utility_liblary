package design

import (
	"fmt"
	"log"
)

// compare PC and Tablet

// create sequence defined
type process interface {
	SetCPU() process
	SetOperation() process
	SetSize() process
	SetOS() process
	GetDevice() deviceProc
}

type deviceProc struct {
	CPU string
	Operation string
	Size string
	OS string
}


// director -> proc struct -> engineering -> implement -> get proc
func DirectorFunc() {
	manufacture := manageDirecter{}	// interface instance

	pcBuilder := &pcBuilder{} 	// pc struct
	manufacture.setBuilder(pcBuilder)	// pc process apply
	manufacture.constructor()	// interface implement
	pc := pcBuilder.GetDevice()	// get pc
	if pc.CPU == "" {
		log.Fatal("CPU performance is unknown")
	}
	if pc.Operation == "" {
		log.Fatal("Operation performance is unknown")
	}
	if pc.Size == "" {
		log.Fatal("Size performance is unknown")
	}
	if pc.OS == "" {
		log.Fatal("OS performance is unknown")
	}
	fmt.Println("personal computer: ", pc)

	tabBuilder := &tabletBuilder{}
	manufacture.setBuilder(tabBuilder)
	manufacture.constructor()
	tablet := tabBuilder.GetDevice()
	if tablet.CPU == "" {
		log.Fatal("CPU performance is unknown")
	}
	if tablet.Operation == "" {
		log.Fatal("Operation performance is unknown")
	}
	if tablet.Size == "" {
		log.Fatal("Size performance is unknown")
	}
	if tablet.OS == "" {
		log.Fatal("OS performance is unknown")
	}
	fmt.Println("tablet: ", tablet)
}


// create director
type manageDirecter struct {
	Builder process
}

func (m *manageDirecter) constructor() {
	m.Builder.
		SetCPU().
		SetOperation().
		SetSize().
		SetOS()
}

func (m *manageDirecter) setBuilder(p process) {
	m.Builder = p
}

// create PC
// interface all implement therefore pcBuilder can be setBuilder method args
type pcBuilder struct {
	PC deviceProc
}

func (p *pcBuilder) SetCPU() process {
	p.PC.CPU = "16GB"
	return p
}
func (p *pcBuilder) SetOperation() process {
	p.PC.Operation = "keyboard"
	return p
}
func (p *pcBuilder) SetSize() process {
	p.PC.Size = "13Inch"
	return p
}
func (p *pcBuilder) SetOS() process {
	p.PC.OS = "Linux, Mac OS or Windows"
	return p
}
func (p *pcBuilder) GetDevice() deviceProc {
	return p.PC
}

// create Tablet
// interface all implement therefore pcBuilder can be setBuilder method args
type tabletBuilder struct {
	Tablet deviceProc
}

func (p *tabletBuilder) SetCPU() process {
	p.Tablet.CPU = "4GB"
	return p
}
func (p *tabletBuilder) SetOperation() process {
	p.Tablet.Operation = "touch panel"
	return p
}
func (p *tabletBuilder) SetSize() process {
	p.Tablet.Size = "6Inch"
	return p
}
func (p *tabletBuilder) SetOS() process {
	p.Tablet.OS = "iOS, Android"
	return p
}
func (p *tabletBuilder) GetDevice() deviceProc {
	return p.Tablet
}