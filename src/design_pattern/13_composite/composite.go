package _3_composite

import "fmt"

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(component Component) {
	c.parent = component
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {
}

func (c *component) Print(string) {
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) Print(str string) {
	fmt.Printf("%s-%s\n", str, l.Name())
}

type Composite struct {
	component
	child []Component
}

func NewComposite() *Composite {
	return &Composite{
		child: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.child = append(c.child, child)
}

func (c *Composite) Print(str string) {
	fmt.Printf("%s+%s\n", str, c.Name())
	str += " "
	for _, comp := range c.child {
		comp.Print(str)
	}
}
