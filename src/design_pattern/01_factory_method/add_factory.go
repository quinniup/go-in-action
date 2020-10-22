package _1_factory_method

type AddFactory struct {
}

func (a *AddFactory) Creator() Operator {
	return &AddOperator{
		OperatorBase: &OperatorBase{},
	}
}

type AddOperator struct {
	OperatorBase *OperatorBase
}

func (a AddOperator) SetObjA(x int) {
	a.OperatorBase.SetObjA(x)
}

func (a AddOperator) SetObjB(x int) {
	a.OperatorBase.SetObjB(x)
}

func (a AddOperator) Result() int {
	return a.OperatorBase.a + a.OperatorBase.b
}
