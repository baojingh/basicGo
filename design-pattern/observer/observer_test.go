package observer

import "testing"

/**
  @Author   : bob
  @Datetime : 2023-06-02 下午 11:20
  @File     : observer_test.go
  @Desc     :
*/

func TestNotify(t *testing.T) {
	o1 := &Observer1{}
	o2 := &Observer2{}

	subObj := &SubjectObj{}
	subObj.Register(o1)
	subObj.Register(o2)
	subObj.Notify("hello")

	subObj.Remove(o1)
	subObj.Notify("go")

}
