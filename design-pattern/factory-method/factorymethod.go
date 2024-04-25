package factory_method

import "fmt"

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 9:03
  @File     : factorymethod.go
  @Desc     :
*/

type RuleParser interface {
	Parse(data []byte)
}

type RuleParserFactory interface {
	CreateParser() RuleParser
}

type JsonRuleParserFactory struct {
}

type JsonRuleParser struct {
}

func (j *JsonRuleParser) Parse(data []byte) {
	fmt.Println("implement json parser")
}

func (j *JsonRuleParserFactory) CreateParser() RuleParser {
	fmt.Println("implement json parser factory")
	return &JsonRuleParser{}
}

type YamlRuleParserFactory struct {
}

type YamlRuleParser struct {
}

func (j *YamlRuleParser) Parse(data []byte) {
	fmt.Println("implement yaml parser")
}

func (j *YamlRuleParserFactory) CreateParser() RuleParser {
	fmt.Println("implement yaml parser factory")
	return &YamlRuleParser{}
}

func NewRuleParserFactory(t string) RuleParserFactory {
	switch t {
	case "yaml":
		return &YamlRuleParserFactory{}
	case "json":
		return &JsonRuleParserFactory{}
	default:
		fmt.Println("NOT implement")
		return nil
	}
}
