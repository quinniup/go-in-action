package _2_abstract_factory

import "fmt"

type TwoA struct {
}

func (o TwoA) DoReceiving() {
	fmt.Println("receiving two-a:", o)
}
