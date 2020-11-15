package main

import (
	"fmt"
)

func main() {
	//内置函数（不推荐）
	print("好吃不过饺子")  //输出不换行
	println("好玩不过嫂子")  //输出换行
	print("好吃不过饺子")  //输出不换行

	//fmt包（推荐使用）
	fmt.Print("好吃不过饺子")  //输出不换行
	fmt.Println("好玩不过嫂子")  //输出换行
	fmt.Print("好吃不过饺子")  //输出不换行
	fmt.Println("我叫","猴六","我有","媳妇、、、")
	fmt.Println("我叫猴六我有媳妇、、、")

	//fmt包扩展：格式化输出
	//%s, 占位符   "文本"
	//%d，占位符   整数
	//%f，占位符   小数（浮点数）
	//%.2f,占位符  小数(保留小数点后两位，四舍五入)
	//百分比，当想输出一个%时，要写两个%

	// 使用占位符时，要是用 printf，
	fmt.Printf("老汉开着%s,去接%s的媳妇，多少钱一次？ %d块。嫂子给打个折吧，%.2f怎么样？小哥哥包你100%%满意")

	/*

	 */

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()



}
