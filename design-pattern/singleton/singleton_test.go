package singleton

import (
	"testing"
)

/**
  @Author   : bob
  @Datetime : 2023-06-02 下午 10:42
  @File     : singleton_test.go
  @Desc     :
*/

func TestGetSingleton(t *testing.T) {
	for i := 0; i < 300; i++ {
		go func() {
			_ = GetSingleton()
		}()

	}

}
