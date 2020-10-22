package _3_builder

type LenovoBuilder struct {
	result string
}

func (l *LenovoBuilder) Result() string {
	return l.result
}

func (l *LenovoBuilder) Screen() {
	l.result += "screen is samsung; "
}

func (l *LenovoBuilder) Mouse() {
	l.result += "mouse is a4tech; "
}

func (l *LenovoBuilder) Keyboard() {
	l.result += "keyboard is cherry; "
}
