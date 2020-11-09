package _3_composite

import "fmt"

// 组合模式统一对象和对象集，使得相同的接口使用对象和对象集
// 组合模式下，常用树状结构，用于统一叶子节点和树节点访问，并且可以用于应用某一操作到所有子节点。

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
