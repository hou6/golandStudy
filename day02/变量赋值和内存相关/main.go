package main

import "fmt"

var number = 99

func main() {
	fmt.Println(number) //99
	number := 88
	if true{
		number = 123
		fmt.Println(number)  //123
	}
	fmt.Println(number)  //123

	//name := "猴六"
	//nickname := name  //此时的name值会被复制一份
	//
	//fmt.Println(name,&name)
	//fmt.Println(nickname,&nickname)   //此时的nickname会指向和name不同的新的内存地址
	//
	//name = "alex"  // name 的值会被覆盖，但是内存地址不变
	//fmt.Println(name,&name)
	//fmt.Println(nickname,&nickname)
	//
	//var v4,v5 = 11,12
	//fmt.Println(v4 + v5)
}
