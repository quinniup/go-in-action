package _2_abstract_factory

import "testing"

func createFactory(base FactoryBase) {
	base.Receiving().DoReceiving()
	base.Pushing().DoPushing()
}

// 抽象工厂
// 生成的所有对象具有关联关系，针对*工厂方法*模式(即，抽象工厂模式为工厂方法模式的升级)
func TestAbstractFactory(t *testing.T) {
	one := &OneFactory{}
	one.Pushing().DoPushing()
	one.Receiving().DoReceiving()

	two := &TwoFactory{}
	createFactory(two)
}
