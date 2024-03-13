package main

import (
	"fmt"
	"testing"
)

/**
  Author     : He Bao Jing
  Date       : 3/13/2024 3:04 PM
  Description:
*/

func TestStringCut(t *testing.T) {
	mystring := "12345"

	fmt.Println(mystring)
	fmt.Println(mystring[:len(mystring)])
	fmt.Println(mystring[:len(mystring)-1])
	fmt.Println(mystring[:len(mystring)-2])
}
