package chain

import (
	"fmt"
)

/**
  @Author:      He Bao Jing
  @Date:        6/2/2023 5:35 PM
  @Description:
*/

// SensitiveWordFilter 接口。判断是否是敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain 数组。存放过滤器对象。此对象是一个接口
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
	fmt.Printf("Register filter success\n")
}

// StartFilter 执行过了逻辑
func (c *SensitiveWordFilterChain) StartFilter(content string) bool {
	for _, f := range c.filters {
		if f.Filter(content) {
			fmt.Println("Found sensitive words...")
			return true
		}
	}
	return false
}
