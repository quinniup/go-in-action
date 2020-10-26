package _0_facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	const preRes = "Module A : Module A string.\nModule B : Module B string."
	api := NewAPI()
	ret := api.Creator()
	if preRes != ret {
		t.Fatal("New Facade API is error.")
	}
}
