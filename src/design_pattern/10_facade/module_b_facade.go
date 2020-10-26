package _0_facade

type ModuleB interface {
	CreatorB() string
}

type moduleBImpl struct {
}

func NewModuleB() ModuleB {
	return &moduleBImpl{}
}

func (a *moduleBImpl) CreatorB() string {
	return "Module B string."
}
