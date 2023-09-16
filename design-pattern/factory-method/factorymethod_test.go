package factory_method

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 9:22
  @File     : factorymethod_test.go
  @Desc     :
*/

func TestRuleParserFactory_CreateParser(t *testing.T) {

	factory := NewRuleParserFactory("json")
	parser := factory.CreateParser()
	parser.Parse([]byte("hello"))

	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want RuleParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonRuleParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlRuleParserFactory{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRuleParserFactory(tt.args.t)
			fmt.Printf("got: %v; want: %v", got, tt.want)

		})
	}

}
