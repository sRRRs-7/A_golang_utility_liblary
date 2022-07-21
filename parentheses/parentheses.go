package parentheses

import (
	"fmt"
	"strings"
)

type Parentheses struct {
	open []string
	close []string
}

func ParenthesesMain() {
	p := Parentheses{}
	var input string
	fmt.Print("input value: ")
	fmt.Scanf("%s", &input)

	arr := strings.Split(input, "")
	p.switchFunc(arr)

	p.pair()

	open := strings.Join(p.open,"")
	close := strings.Join(p.close,"")
	openArr := strings.Split(open, "")
	closeArr := strings.Split(close, "")
	fmt.Println("open rest: ", openArr)
	fmt.Println("close rest: ", closeArr)

	err := p.check(openArr, closeArr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("complete")
	}
}

func (p *Parentheses) switchFunc(arr []string) {
	for _, v := range arr {
		switch v {
		case "(":
			p.open = append(p.open, v)
		case "[":
			p.open = append(p.open, v)
		case "{":
			p.open = append(p.open, v)
		case ")":
			p.close = append(p.close, v)
		case "]":
			p.close = append(p.close, v)
		case "}":
			p.close = append(p.close, v)
		}
	}
}

func (p *Parentheses) pair() {
	for _, v := range p.close {
		switch v {
		case ")":
			for i, val := range p.open {
				if val != "(" {
					continue
				} else {
					p.open[i] = ""
					p.close[i]= ""
					break
				}
			}
		case "]":
			for i, val := range p.open {
				if val != "[" {
					continue
				} else {
					p.open[i] = ""
					p.close[i]= ""
					break
				}
			}
		case "}":
			for i, val := range p.open {
				if val != "{" {
					continue
				} else {
					p.open[i] = ""
					p.close[i]= ""
					break
				}
			}
		}
	}
}

func (p *Parentheses) check(open, close []string) error {
	if len(open) != 0 {
		return fmt.Errorf("open parentheses %d rest", len(open))
	}
	if len(close) != 0 {
		return fmt.Errorf("close parentheses %d rest", len(close))
	}

	return nil
}
