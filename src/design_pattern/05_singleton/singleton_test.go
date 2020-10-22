package _5_singleton

import "testing"

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("The singleton isn't equal")
	}
}
