package main

import "fmt"

var school1 string = "houliu666"  //全局变量
var school2  = "houliu666"  //全局变量
//var school3  := "houliu666"  //全局变量，不能简写成 := 的方式

func main() {

	var name  = "愿有岁月可回首"  //局部变量
	fmt.Println(name)

	if true {
		age := 23   //局部变量
		fmt.Println(age)
		fmt.Println(name)
		fmt.Println(school1)
		fmt.Println(school2)
		//fmt.Println(school3)
	}

	//fmt.Println(age)  // 此处不能使用age，类似于 java
	fmt.Println(name)

	fmt.Println(school1)
	fmt.Println(school2)
	//fmt.Println(school3)



	//var name string = "houliu"
	//fmt.Println(name)
	//
	//var age int = 25
	//fmt.Println(age)
	//
	//var flag bool = false
	//fmt.Println(flag)
	//
	//var gender string
	//gender = "男"
	//fmt.Println(gender)

	//var name string
	//fmt.Scanf("%s",&name)  //用户输入，输入的文本赋值给name
	//fmt.Println(name)
	//
	//var age int
	//fmt.Scanf("%d",&age)
	//fmt.Println(age)

	//var (
	//	name = "猴六"
	//	age = 23
	//	hobby = "大保健"
	//	salary = 10000.00
	//	gender string  //string 只声明，不赋值，默认值为 ""
	//	length int     //int 只声明，不赋值，默认值为 0
	//	sb bool        //bool 只声明，不赋值，默认值为false
	//)
	//
	//fmt.Println(name,age,hobby,salary,gender,length,sb)



}
