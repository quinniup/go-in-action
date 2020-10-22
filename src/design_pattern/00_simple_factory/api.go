package _0_simple_factory

type API interface {
	Print(str string) string
}

func NewAPI(val int) API {
	switch val {
	case 1:
		return &FAPI{}
	case 2:
		return &TAPI{}
	}
	return nil
}
