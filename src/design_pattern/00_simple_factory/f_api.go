package _0_simple_factory

import "fmt"

type FAPI struct {
}

func (fapi *FAPI) Print(str string) string {
	fmt.Println("This is f-api printer:", str)
	return "f-api"
}
