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

	// bob: {bob 20}  &bob: &{bob 20}
	fmt.Println("bob:", bob, " &bob:", &bob)

	// jim: &{jim 33}  &jim: 0xc00000a030
	jim := &Person{"jim", 33}
	fmt.Println("jim:", jim, " &jim:", &jim)

	// John: &{John 40}  &John: 0xc00000a038
	var John *Person = &Person{name: "John", age: 40}
	fmt.Println("John:", John, " &John:", &John)

	var c = 20
	d := &c
	// 0xc0000a6090
	fmt.Println("c的內存地址", &c)

	// 20
	fmt.Println("*取出這個地址上的值", *d)

	*d++
	// 21
	fmt.Println("*d", *d)

	// 0xc0000ca030
	fmt.Println("&d", &d)

	fmt.Println("################")
}
