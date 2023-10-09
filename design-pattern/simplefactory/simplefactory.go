package simplefactory

import "fmt"

/**
  @Author   : bob
  @Datetime : 2023-06-02 下午 11:29
  @File     : simplefactory.go
  @Desc     :
*/

type RuleParser interface {
	Parse(data []byte)
}

type JsonRuleParser struct{}

func (j *JsonRuleParser) Parse(data []byte) {
	fmt.Println("implement json parser")
}

type YamlRuleParser struct{}

func (y *YamlRuleParser) Parse(data []byte) {
	fmt.Println("implement yaml parser")
}

func NewRuleParser(ruleType string) RuleParser {
	switch ruleType {
	case "yaml":
		return &YamlRuleParser{}
	case "json":
		return &JsonRuleParser{}
	default:
		fmt.Println("NOT implement")
		return nil
	}
}
