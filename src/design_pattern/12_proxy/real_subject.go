package _2_proxy

type Subject interface {
	Doing() string
}

type RealSubject struct {
}

func (RealSubject) Doing() string {
	return "real"
}
