package package1

import "fmt"

//注意： 在文件夹中编写功能时，首字母要大写
//注意： 文件中的函数首字母是小写，表示此函数只能被当前内部函数调用。首字母是大写，意味着任何包都可以调用
func Google() {
	fmt.Println("我是 Google，，，")
}
