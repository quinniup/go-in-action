package _4_prototype

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, cloneable Cloneable) {
	p.prototypes[name] = cloneable
}
