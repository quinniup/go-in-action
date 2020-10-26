package _1_adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	const preRes = "adapted method"
	adapted := NewAdapted()
	targetString := NewAdapter(adapted).Request()
	if preRes != targetString {
		fmt.Println("Adapter string is error.")
	}
}
