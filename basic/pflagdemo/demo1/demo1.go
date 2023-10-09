package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

/**
  @Author:      He Bao Jing
  @Date:        6/27/2023 3:50 PM
  @Description:
*/

func main() {
	var name *string = pflag.String("name", "bob", "your name")

	var age *int = pflag.Int("age", 0, "your age")

	var score int
	pflag.IntVar(&score, "score", 0, "your score")

	pflag.Parse()

	log.Infof("your name: %s, age: %d, score: %d", *name, *age, score)

}
