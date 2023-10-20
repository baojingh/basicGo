package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 6:26
  @File     : structure_test.go
  @Desc     :
*/

func TestStructDemo(t *testing.T) {
	type student struct {
		id   int
		name string
	}

	stu1 := student{id: 1, name: "bob"}
	stu2 := student{2, "vov"}
	stu3 := new(student)
	stu3.id = 4
	stu3.name = "ww"

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(*stu3)

	s1 := student{}
	fmt.Printf("%d, %s\n", s1.id, s1.name)

}
