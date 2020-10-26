package _0_facade

import "fmt"

// 外观模式
// facade模块同时暴露A、B两个Module的NewXXX和interface，可根据需求直接调用细节部分代码

type API interface {
	Creator() string
}

type apiImpl struct {
	a ModuleA
	b ModuleB
}

func NewAPI() API {
	return &apiImpl{
		a: NewModuleA(),
		b: NewModuleB(),
	}
}

func (a *apiImpl) Creator() string {
	aRet := a.a.CreatorA()
	bRet := a.b.CreatorB()
	return fmt.Sprintf("Module A : %s\nModule B : %s", aRet, bRet)
}
