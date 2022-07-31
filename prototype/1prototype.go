package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type iPhone struct {
	Name string
	Feature *feature
}

type feature struct {
	Width int
	Height int
	Volume string
}

type application struct {
	app []string
}

func Prototype() {
	iPhone50 := iPhone {
		Name: "iPhone50",
		Feature: &feature{
			Width: 4.0,
			Height: 8.0,
			Volume: "256/512/1024",
		},
	}

	iPhone100 := iPhone50
	iPhone100.Name = "iPhone100"
	iPhone100.Feature = iPhone100.Feature.deepCopy()
	iPhone100.Feature.Width = 2.0
	iPhone100.Feature.Height = 4.0
	iPhone100.Feature.Volume = "1024/1PB"

	fmt.Println(iPhone50.Name, iPhone50.Feature)
	fmt.Println(iPhone100.Name, iPhone100.Feature)

	arr1 := application{
		app: []string{"man", "woman"},
	}
	arr2 := arr1.deepCopy()
	arr2[0] = "mysterious"

	fmt.Println(arr1)
	fmt.Println(arr2)
}

// pointer
func (f *feature) deepCopy() *feature {
	return &feature{
		f.Width,
		f.Height,
		f.Volume,
	}
}

// slice
func (a *application) deepCopy() []string {
	var arr []string
	arr = append(arr, a.app...)
	return arr
}

// encode decode
func (i *iPhone) deepCopy() *iPhone {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(i)

	d := gob.NewDecoder(&b)
	new := iPhone{}
	_ = d.Decode(&new)

	return &new
}

