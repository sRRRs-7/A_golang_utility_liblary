package stringBuilder

import (
	"fmt"
	"strings"
)

func Builder() {
	strs := []string{"React", "Golang", "Postgres"}
	sb := strings.Builder{}

	sb.WriteString("<ul> \n")
	for _, v := range strs {
		sb.WriteString("<li> \n")
		sb.WriteString(v)
		sb.WriteString("\n")
		sb.WriteString("</li> \n")
	}
	sb.WriteString("</ul>")

	fmt.Println(sb.String())
}



// HTML logic
type HtmlElement struct {
	tagName string
	content string
	elements []HtmlElement
}

var indentSize = 2

func Main() {
	e := HtmlElement{
		tagName: "div",
	}
	e.elements = append(e.elements, HtmlElement{
		tagName: "div",
		content: "golang",
	})
	e.elements = append(e.elements, HtmlElement{
		tagName: "strong",
		content: "postgres",
	})
	str := e.String()
	fmt.Println(str)
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}

	// create indent
	i := strings.Repeat(" ", indentSize * indent)

	// create start tag
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.tagName))

	// insert content
	if len(e.content) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * indent + 1))
		sb.WriteString(e.content)
		sb.WriteString("\n")
	}

	// create child element
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	// close tag
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.tagName))

	return sb.String()
}