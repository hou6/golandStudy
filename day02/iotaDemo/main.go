package main

import "fmt"

/*
  const 相当于一个计数器，一般用于生成一些连续的数字
 */

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)


func main() {

	// 结果 0 1 2 3 4
	//const(
	//	v0 = 0
	//	v1 = 1
	//	v2 = 2
	//	v3 = 3
	//	v4 = 4
	//)
	//fmt.Println(v0,v1,v2,v3,v4)

	//使用 ioto 后可以写成  // 结果 0 1 2 3 4
	//const(
	//	v0 = iota
	//	v1 = iota
	//	v2 = iota
	//	v3 = iota
	//	v4 = iota
	//)
	//fmt.Println(v0,v1,v2,v3,v4)

	// 或者写成  // 结果 0 1 2 3 4
	//const(
	//	v0 = iota
	//	v1
	//	v2
	//	v3
	//	v4
	//)
	//fmt.Println(v0,v1,v2,v3,v4)


	//结果 1 2 3 4 5
	//const(
	//	v0 = iota + 1
	//	v1
	//	v2
	//	v3
	//	v4
	//)
	//fmt.Println(v0,v1,v2,v3,v4)

	//结果 2 4 5 6 7
	const(
		v0 = iota + 2
		_  //隔 1
		v1
		v2
		v3
		v4
	)
	fmt.Println(v0,v1,v2,v3,v4)


}
