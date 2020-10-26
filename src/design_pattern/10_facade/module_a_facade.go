package _0_facade

type ModuleA interface {
	CreatorA() string
}

type moduleAImpl struct {
}

func NewModuleA() ModuleA {
	return &moduleAImpl{}
}

func (a *moduleAImpl) CreatorA() string {
	return "Module A string."
}
