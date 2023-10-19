package main

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 7/5/2023 下午 10:41
  @File     : logrus_test.go
  @Desc     :
*/

func TestLogrus(t *testing.T) {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	log.Info("hello, world")
	// INFO[2023-05-07 22:42:56] hello, world, 10 is a number, cs is a string
	log.Infof("hello, world, %d is a number, %s is a string", 10, "cs")
}
