package _1_factory_method

type SubFactory struct {
}

func (a *SubFactory) Creator() Operator {
	return &SubOperator{
		OperatorBase: &OperatorBase{},
	}
}

type SubOperator struct {
	OperatorBase *OperatorBase
}

func (a SubOperator) SetObjA(x int) {
	a.OperatorBase.SetObjA(x)
}

func (a SubOperator) SetObjB(x int) {
	a.OperatorBase.SetObjB(x)
}

func (a SubOperator) Result() int {
	return a.OperatorBase.a - a.OperatorBase.b
}
