package main

import "testing"

/**
  Author     : He Bao Jing
  Date       : 11/28/2023 5:58 PM
  Description:
*/

// 闭包
func TestClosure(t *testing.T) {
	/**
	得到的结果全部是1
	闭包没生效
	*/
	println(f1()())
	println(f1()())
	println(f1()())
}

func TestClosure2(t *testing.T) {
	/**
	得到的结果是
	1
	2
	3
	*/
	ff := f1()
	println(ff())
	println(ff())
	println(ff())
}

func f1() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
