package _2_abstract_factory

type Receiving interface {
	DoReceiving()
}

type Pushing interface {
	DoPushing()
}

type FactoryBase interface {
	Receiving() Receiving
	Pushing() Pushing
}
