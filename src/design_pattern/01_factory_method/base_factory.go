package _1_factory_method

// Go语言不存在继承关系，需要通过匿名组合方式进行实现
// 工厂方法模式通过使用子类的方式延迟生成对象到子类中

type OperatorBase struct {
	a, b int
}

type OperatorFactory interface {
	Creator() Operator
}

type Operator interface {
	SetObjA(x int)
	SetObjB(x int)
	Result() int
}

func (ob *OperatorBase) SetObjA(x int) {
	ob.a = x
}

func (ob *OperatorBase) SetObjB(x int) {
	ob.b = x
}
