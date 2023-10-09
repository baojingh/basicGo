package chain

import (
	"strings"
)

/**
  @Author:      He Bao Jing
  @Date:        6/2/2023 5:35 PM
  @Description:
*/

// AdSensitiveFilter 判断是否有广告敏感词
type AdSensitiveFilter struct{}

func (f *AdSensitiveFilter) Filter(content string) bool {
	// 执行判断逻辑
	isContaine := strings.Contains(content, "ad")
	return isContaine
}
