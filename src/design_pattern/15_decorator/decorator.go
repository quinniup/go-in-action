package _5_decorator

// 装饰模式使用对象组合的方式通过动态改变或者增加对象的行为。
// Golang借助匿名组合和非入侵接口可以很方便实现装饰模式。使用匿名组合，在装饰器中不需要显式定义转调原对象方法。

type Component interface {
	Calc() int
}

// 初始化
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// 乘法运算
type MulDecorator struct {
	// 匿名组合
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// 加法运算
type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
