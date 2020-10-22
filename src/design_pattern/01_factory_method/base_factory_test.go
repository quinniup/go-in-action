package _1_factory_method

import (
	"fmt"
	"testing"
)

func creator(factory OperatorFactory, a int, b int) int {
	f := factory.Creator()
	f.SetObjA(a)
	f.SetObjB(b)
	return f.Result()
}

func TestFactoryMethod(t *testing.T) {
	factoryA := AddFactory{}
	resA := creator(&factoryA, 1, 2)
	fmt.Println("a + b =", resA)
	if resA != 3 {
		t.Fatal("The factoryA calculate failed.")
	}

	factoryB := SubFactory{}
	resB := creator(&factoryB, 4, 2)
	fmt.Println("a - b =", resB)
	if resB != 2 {
		t.Fatal("The factoryB calculate failed.")
	}
}
