package _3_builder

type DellBuilder struct {
	result string
}

func (l *DellBuilder) Result() string {
	return l.result
}

func (l *DellBuilder) Screen() {
	l.result += "screen is boe; "
}

func (l *DellBuilder) Mouse() {
	l.result += "mouse is logitech; "
}

func (l *DellBuilder) Keyboard() {
	l.result += "keyboard is also logitech; "
}
