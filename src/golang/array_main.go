package main

import "fmt"

// 一维数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str0 = [5]string{0: "a", 1: "b"} // 索引号形式指定数组元素值

// 多维数组
var array0 [5][3]int
var array1 [2][2]int = [...][2]int{{1, 2}, {3, 4}}

func main() {

	// 多维数组
	fmt.Println("arr0:", arr0, ";arr1:", arr1, ";arr2:", arr2, ";str0:", str0)

	// 未初始化值为数组类型默认值
	a := [3]int{1, 2}
	// 输出结果: [1 2 0]
	fmt.Println("a:", a)

	b := [...]struct {
		name  string
		value int
	}{
		{
			name:  "a",
			value: 1,
		},
		{
			"b",
			2,
		},
	}
	fmt.Println("struct params:", b)

	// 多维数组
	fmt.Println("array0:", array0, ";array1:", array1)
	//
	array2 := [...][2]int{{11, 22}, {33, 44}}
	fmt.Println("array2:", array2)

}
