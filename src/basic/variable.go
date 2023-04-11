package main

import (
	"fmt"
	"unsafe"
)

func foo(arr [5]int8) {
	fmt.Printf("%d\n", arr[0])
	fmt.Printf("%d\n", len(arr))
	fmt.Printf("%d\n", unsafe.Sizeof(arr))
}

func main() {

	var arr1 = [5]int8{1, 2, 3, 4, 5}
	foo(arr1)

	fmt.Println("this is end")
}
