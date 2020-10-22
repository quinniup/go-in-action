package _2_abstract_factory

import "fmt"

type OneA struct {
}

func (o OneA) DoPushing() {
	fmt.Println("pushing one-a:", o)
}
