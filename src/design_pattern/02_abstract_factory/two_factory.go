package _2_abstract_factory

type TwoFactory struct {
}

func (t TwoFactory) Receiving() Receiving {
	return TwoA{}
}

func (t TwoFactory) Pushing() Pushing {
	return TwoB{}
}
