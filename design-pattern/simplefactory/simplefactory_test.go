package simplefactory

import (
	"reflect"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 2023-06-03 上午 9:05
  @File     : simplefactory_test.go
  @Desc     :
*/

func TestNewRuleParser(t *testing.T) {
	parser := NewRuleParser("json")
	parser.Parse([]byte("hello"))

	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want RuleParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonRuleParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlRuleParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}

}
