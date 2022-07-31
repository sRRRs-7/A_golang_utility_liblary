package factory

import (
	"fmt"
)

type phone struct {
	app, number, mail string
}

const (
	one = iota
	two
	three
	four
)

func FactoryDesign2() {
	arr := []phone{}

	init := constructor2("000-000-000", "abcde@email.com")

	phone1 := init("Youtube")
	phone2 := init("TikTok")

	arr = append(arr, *phone1, *phone2)
	fmt.Println(arr)

	init2 := constructor2("111-111-111", "fghij@email.com")

	phone3 := init2("Instagram")
	phone4 := init2("Twitter")

	fmt.Println(len(phone4.app))
	fmt.Println(*phone3, *phone4)

	for _, v := range phone4.app {
		// rune -> string
		fmt.Println(string(v))
	}

	// const iota -> index
	fmt.Println(one, two, three, four)
}

func constructor2(number, mail string) func(app string) *phone {
	return func(app string) *phone {
		return &phone{app, number, mail}
	}
}