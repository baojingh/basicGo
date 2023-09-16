package chain

import "strings"

/**
  @Author:      He Bao Jing
  @Date:        6/2/2023 6:45 PM
  @Description:
*/

// PoliticalSensitiveFilter 判断是否有政治敏感词
type PoliticalSensitiveFilter struct{}

func (f *PoliticalSensitiveFilter) Filter(content string) bool {
	isContaine := strings.Contains(content, "political")
	return isContaine
}
