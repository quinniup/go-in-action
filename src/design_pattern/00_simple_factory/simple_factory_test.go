package _0_simple_factory

import (
	"fmt"
	"testing"
)

/*
	由于Go语言没有构造函数，一般约定使用NewXXX命名形式进行函数创建，并进行相关初始化；
	伪代码范例：
		func NewXXX() *interface {}
	仅有API interface & NewXXX() 外部可见, 简单工厂具体实现进行了封装处理。
*/
func TestSimpleFactory(t *testing.T) {

	fapi := NewAPI(1).Print("First API.")
	if fapi != "f-api" {
		t.Fatal("f-api of testing is error.")
	}
	fmt.Println()
	tapi := NewAPI(2).Print("First API.")
	if tapi != "t-api" {
		t.Fatal("f-api of testing is error.")
	}
}
