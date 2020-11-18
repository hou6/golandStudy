package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//示例1
	//var name string
	//fmt.Println("请输入用户名：")
	//fmt.Scan(&name)  // &name 表示的是name的内存地址
	//fmt.Printf(name)

	//示例2
	//var name string
	//var age int
	//fmt.Println("请输入用户名：")
	////当使用Scan时，会提示用户输入
	////用户输入完成后，会得到两个值，count 和 err
	////count : 表示用户输入正确了几个值
	////err ： 用户输入错误则是错误信息
	//_, err := fmt.Scan(&name,&age)  //这里如果不想用count，可以把 count 用 _ 代替，这样下面可以不用count也不会报错
	//if err == nil {
	//	fmt.Println(name,age)
	//}else {
	//	fmt.Println("您输入的数据有错误，错误信息为：", err)
	//}
	//fmt.Println(name,age)

	/*
	请输入用户名：
	alex niupo
	1 expected integer   //正确了一个，错误信息为： expected integer
	alex 0
	 */

	//说明 ： fmt.Scan 要求输入几个值，就必须输入几个值，否则就会一直等待
	//说明 ： fmt.Scanln 不会一直等待，它只等待回车，只要一点回车，它就表示你输入完成了

	////示例3  fmt.Scanf  格式化
	//var name string
	//var age int
	//fmt.Print("请输入用户名：")
	////fmt.Scanf("%s",&name)
	////fmt.Scanf("我是%s",&name)
	////fmt.Scanf("我是%s 今年23岁",&name)
	////fmt.Scanf("我是%s 今年%d 岁",&name,&age)   //注意，占位符后一定要加空格
	//count, err := fmt.Scanf("我是%s 今年%d 岁", &name, &age) //注意，占位符后一定要加空格
	//if err == nil {
	//	fmt.Println(name,age)
	//}else {
	//	fmt.Println("您输入的数据有错误，错误信息为：", err, "正确的个数为：", count)
	//}

	// 示例 4
	//var message string
	//fmt.Println("请输入用户信息:")
	//fmt.Scanln(&message)
	//fmt.Println(message)

	// 示例 5
	fmt.Print("请输入您的个人信息：")
	reader := bufio.NewReader(os.Stdin)  //获取用户标准输入
	line, isPrefix, err := reader.ReadLine()  //读取用户输入
	//line : 从 Stdin 中读取一行的数据 （字节集合 ————》 转化为字符串）
	// reader 默认一次能读 4096 个字节 （4096/3）
		// 1. 一次性读完，isPrefix = false
		// 2. 先读一部分，isPrefix = true ， 再去读取 isPrefix = false
	data := string(line)  // 字节转换为字符串
	fmt.Println(data)
	fmt.Println(isPrefix)
	fmt.Println(err)
}
