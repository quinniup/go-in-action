package _1_adapter

// 被适配对象
type Adapted interface {
	SpecificRequest() string
}

type adapted struct {
}

func (*adapted) SpecificRequest() string {
	return "adapted method"
}

func NewAdapted() Adapted {
	return &adapted{}
}
