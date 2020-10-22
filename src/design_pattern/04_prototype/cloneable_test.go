package _4_prototype

import "testing"

func TestCloneable(t *testing.T) {
	t1 := Manager.Get("p")
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("Get clone obj is not current")
	}

	p := Manager.Get("p").Clone()
	t0 := p.(*Pattern)
	if t0.name != "pattern" {
		t.Fatal("Get clone obj is error")
	}
}
