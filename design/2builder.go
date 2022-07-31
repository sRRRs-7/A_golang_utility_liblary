package design

import (
	"fmt"
	"strings"
)

type HtmlElement struct {
	tagName 	string
	content 	string
	elements 	[]HtmlElement
}

type HtmlBuilder struct {
	rootTagName 	string
	root 			HtmlElement
}

var indentSize = 2

func Main() {
	b := NewHtmlBuilder("div")
	b.addChildTag("button", "click").
		addChildTag("header", "header").
		addChildTag("footer", "footer")
	fmt.Println(b.String())
}

func NewHtmlBuilder(tagName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootTagName: tagName,
		root: HtmlElement{
			tagName: tagName,
			content: "",
			elements: []HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.string(0)
}

func (b *HtmlBuilder) addChildTag(tagName, content string) *HtmlBuilder{
	el := HtmlElement{tagName, content, []HtmlElement{}}
	b.root.elements = append(b.root.elements, el)
	return b
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
