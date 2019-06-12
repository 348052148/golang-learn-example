package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)


type Document struct {
	root *Element
}

type Element struct {
	Name string
	Child   []*Element
	Attr	[]Attr
	Data   string
}

type Attr struct {
	Name string
	Value string
}

func NewElement(name string, attrs []xml.Attr) *Element  {
	var attrList []Attr
	for _,attr := range attrs {
		attrList = append(attrList,Attr{
			//区分Space Local
			Name:attr.Name.Local,
			Value:attr.Value,
		})
	}
	return &Element{
		Name:name,
		Attr:attrList,
	}
}

func (e *Element)AddData(value string)  {
	e.Data = value
}

func (e *Element)AddChild(node *Element)  {
	e.Child = append(e.Child, node)
}

func (e *Element)GetChildAll(name string) []*Element  {
	var elems []*Element
	for _,elem := range e.Child {
		if elem.Name == name {
			elems = append(elems, elem)
		}
	}
	return elems
}

func (e *Element)GetChild(name string) *Element  {
	for _,elem := range e.Child {
		if elem.Name == name {
			return elem
		}
	}
	return nil
}

func NewDocument() *Document {
	return &Document{
		root:nil,
	}
}

func (doc *Document)Read(reader io.Reader)  {
	dec := xml.NewDecoder(reader)
	var stack stack
	for {
		t, err := dec.RawToken()
		if err == io.EOF || err!=nil {
			break
		}
		//stack.push(doc.root)
		top := stack.peek()
		switch t := t.(type) {
		case xml.StartElement:
			e := NewElement(t.Name.Local, t.Attr)
			if (doc.root == nil) {
				doc.root = e
			}else {
				top.(*Element).AddChild(e)
			}
			stack.push(e)
		case xml.EndElement:
			stack.pop()
		case xml.CharData:
			if e,ok := top.(*Element); ok {
				e.AddData(string(t))
			}
		case xml.Comment:
			//fmt.Println(t)
		case xml.Directive:
			//fmt.Println(t)
		case xml.ProcInst:
			//fmt.Println(t)
		}
	}
}


type stack struct {
	data []interface{}
}

func (s *stack)empty() bool  {
	return len(s.data) == 0
}

func (s *stack)push(value interface{})  {
	s.data = append(s.data, value)
}

func (s *stack)pop() interface{}  {
	value := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = nil
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *stack)peek() interface{}  {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}


func main() {
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	doc := NewDocument()
	doc.Read(strings.NewReader(data))
	fmt.Println(doc.root.GetChild("City").Data)
}
