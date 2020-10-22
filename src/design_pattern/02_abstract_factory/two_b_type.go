package _2_abstract_factory

import "fmt"

type TwoB struct {
}

func (o TwoB) DoPushing() {
	fmt.Println("pushing two-b:", o)
}
