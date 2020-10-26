package _1_adapter

//适配器模式
// 	用于转换一种接口并适配并一种接口的设计模式。
// 	Adapter中包含匿名组合Adapted接口，Adapter中也拥有SpecificRequest方法
// 	由于Golang非入侵式接口特征，Adapter也适用于Adapted接口
type adapter struct {
	Adapted Adapted
}

// 进行接口兼容
func (a *adapter) Request() string {
	return a.Adapted.SpecificRequest()
}

func NewAdapter(a Adapted) Target {
	return &adapter{Adapted: a}
}
