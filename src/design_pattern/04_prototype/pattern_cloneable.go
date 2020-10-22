package _4_prototype

var Manager *PrototypeManager

type Pattern struct {
	name string
}

func (p *Pattern) Clone() Cloneable {
	tc := *p
	return &tc
}

func init() {
	Manager = NewPrototypeManager()
	p := &Pattern{
		name: "pattern",
	}
	Manager.Set("p", p)
}
