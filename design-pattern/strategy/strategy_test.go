package strategy

import "testing"

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 10:56
  @File     : strategy_test.go
  @Desc     :
*/

func TestNewStorageStrategy(t *testing.T) {
	strategy, _ := NewStorageStrategy("normal")
	strategy.Save("name", "data")
}
