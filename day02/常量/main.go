package main

import "fmt"

/**
常量在定义的时候就必须赋值，因为一旦赋值无法修改
 */

/**
全局常量，一般常量都是定义在全局的
 */
const data = 999
const (
	pi = 3.1415926
	gender = "男"
)

func main() {

	//const age int = 23
	const age  = 23
	//age = 25  //不能被修改
	fmt.Println(age)

	//一行定义多个常量
	const a1,a2,a3 = 11,12,13
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	//常量因式分解
	const (
		v1 = "haha"
		v2 = "hehe"
	)
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(data)
	fmt.Println(pi)
	fmt.Println(gender)


}
