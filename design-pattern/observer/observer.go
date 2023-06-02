package observer

import "fmt"

/**
  @Author   : bob
  @Datetime : 2023-06-02 下午 11:06
  @File     : observer.go
  @Desc     :
*/

type Observer interface {
	Update(msg string)
}

type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(msg string)
}

type SubjectObj struct {
	observers []Observer
}

func (subject *SubjectObj) Register(observer Observer) {
	subject.observers = append(subject.observers, observer)
	fmt.Println("Register observer success")
}

func (subject *SubjectObj) Remove(observer Observer) {
	for i, ob := range subject.observers {
		if ob == observer {
			subject.observers = append(subject.observers[:i],
				subject.observers[i+1:]...)
			fmt.Println("Remove observer success")
		}
	}
}

// Notify 通知给所有观察者
func (subject *SubjectObj) Notify(msg string) {
	for _, o := range subject.observers {
		o.Update(msg)
	}
}

type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("observer1: %s\n", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("observer2: %s\n", msg)
}
