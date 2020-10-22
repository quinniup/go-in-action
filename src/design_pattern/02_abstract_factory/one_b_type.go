package _2_abstract_factory

import "fmt"

type OneB struct {
}

func (o OneB) DoReceiving() {
	fmt.Println("receiving one-b:", o)
}
