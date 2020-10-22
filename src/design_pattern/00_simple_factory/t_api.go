package _0_simple_factory

import "fmt"

type TAPI struct {
}

func (fapi *TAPI) Print(str string) string {
	fmt.Println("This is t-api printer:", str)
	return "t-api"
}
