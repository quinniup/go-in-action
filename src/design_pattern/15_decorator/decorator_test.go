package _5_decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var comm Component = &ConcreteComponent{}
	comm = WarpAddDecorator(comm, 10)
	comm = WarpAddDecorator(comm, 8)

	res := comm.Calc()
	fmt.Println(res)
}
