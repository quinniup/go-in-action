package _2_abstract_factory

type OneFactory struct {
}

func (o OneFactory) Receiving() Receiving {
	return &OneB{}
}

func (o OneFactory) Pushing() Pushing {
	return &OneA{}
}
