package _3_builder

type BaseBuilder interface {
	Screen()
	Mouse()
	Keyboard()
}

type Computer struct {
	builder BaseBuilder
}

func NewComputer(builder BaseBuilder) *Computer {
	return &Computer{builder: builder}
}

func (c *Computer) Construct() {
	c.builder.Screen()
	c.builder.Mouse()
	c.builder.Keyboard()
}
