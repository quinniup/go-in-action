package _2_proxy

//代理模式
//	用于延时处理操作、或在实际操作前后进行其他处理。
//应用场景
// 虚代理、cow代理、远程代理、保护代理、Cache代理、防火墙代理、同步代理、智能指引

type Proxy struct {
	real RealSubject
}

func (receiver Proxy) Doing() string {
	var res string
	res += "pre:"
	res += receiver.real.Doing()
	res += ":after"
	return res
}
