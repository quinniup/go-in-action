package _3_composite

import "testing"

func TestComposite(t *testing.T) {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "child_1")
	c2 := NewComponent(CompositeNode, "child_2")

	l1 := NewComponent(LeafNode, "leave_1")
	l2 := NewComponent(LeafNode, "leave_2")

	root.AddChild(c1)
	root.AddChild(c2)

	c1.AddChild(l1)
	c2.AddChild(l2)

	root.Print("")
}
