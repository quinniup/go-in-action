package _2_proxy

import "testing"

func TestProxy_Doing(t *testing.T) {
	const preRes = "pre:real:after"
	var sub Subject
	sub = &Proxy{}
	res := sub.Doing()
	if res != preRes {
		t.Fatal("The proxy model is error.", res)
	}
}
