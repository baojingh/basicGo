package chain

import "strings"

/**
  @Author:      He Bao Jing
  @Date:        6/2/2023 6:45 PM
  @Description:
*/

// ForceSensitiveFilter 判断是否有暴力敏感词
type ForceSensitiveFilter struct{}

func (f *ForceSensitiveFilter) Filter(content string) bool {
	isContaine := strings.Contains(content, "force")
	return isContaine
}
