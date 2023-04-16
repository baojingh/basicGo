package main

import (
	"fmt"
)

/**
&是取地址符。
*可以表示一个变量是指针类型，也可以表示指针类型变量所指向的存储单元，也就是这个地址所指向的值。
*/

type Person struct {
	name string
	age  int
}

func main() {

	//https://blog.csdn.net/weixin_41910694/article/details/107917518
	bob := Person{"bob", 20}

	fmt.Println(bob)

	fmt.Println("################")
}
