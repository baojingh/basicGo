package main

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

Get()

func main() {
	mapArr := Values{}
	mapArr.Set("sydout", "1")


}
