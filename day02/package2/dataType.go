package main

import "fmt"

func main() {

	fmt.Println("猴六")
	fmt.Println("猴六","侯琦")
	fmt.Println("猴六" + "侯琦")
	fmt.Println("猴六" + "houliu")
	fmt.Println("猴六" + "666")
	fmt.Println("666")
	fmt.Println(666)
	fmt.Println("1" + "2")   //12
	fmt.Println(1 + 2)   // 3
	fmt.Println(1 > 3)   // false
	fmt.Println(1 < 3)   // true

	//超前
	if 2 > 1 {
		fmt.Println("叫爸爸")
	}else {
		fmt.Println("叫弟弟")
	}

}
