package main

import "fmt"

/**
  @Author   : bob
  @Datetime : 2023-05-29 下午 10:36
  @File     : mapArrayDemo.go
  @Desc     :
*/

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]string

func (v Values) Set(key, value string) {
	v[key] = []string{value}
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Get get the first value
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	val := v[key]
	if len(val) == 0 {
		return ""
	}
	return val[0]
}

func main() {
	mapArr := Values{}
	mapArr.Set("stdout", "v1")
	mapArr.Set("stderr", "v1")

	//mapArr.Set("stdout", "v2")
	//mapArr.Set("stderr", "v2")

	for k, vArr := range mapArr {
		fmt.Printf("%s, %s\n", k, vArr)
	}
}
