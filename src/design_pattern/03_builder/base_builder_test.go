package _3_builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	dellBuilder := &DellBuilder{}
	NewComputer(dellBuilder).Construct()
	dellRes := dellBuilder.Result()
	if "screen is boe; mouse is logitech; keyboard is also logitech; " != dellRes {
		t.Fatal("Dell create the computer having the problem.")
	}

	lenovoBuilder := &LenovoBuilder{}
	NewComputer(lenovoBuilder).Construct()
	lenovoRes := lenovoBuilder.Result()
	//fmt.Println(lenovoRes)
	if "screen is samsung; mouse is a4tech; keyboard is cherry; " != lenovoRes {
		t.Fatal("Dell create the computer having the problem.")
	}
}
