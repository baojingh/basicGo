package singleton

import (
	"fmt"
	"sync"
)

/**
  @Author   : bob
  @Datetime : 2023-06-02 下午 10:36
  @File     : singleton.go
  @Desc     :
*/

var (
	lazySingle *Singleton
	single     = &sync.Once{}
)

type Singleton struct {
	name string
}

func GetSingleton() *Singleton {
	if lazySingle != nil {
		fmt.Println("Singleton exists already")
		return lazySingle
	}
	single.Do(func() {
		lazySingle = &Singleton{name: "bob"}
		fmt.Println("Singleton is created")
	})
	return lazySingle
}
