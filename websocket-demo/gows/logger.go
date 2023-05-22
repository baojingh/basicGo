package gows

import (
	"github.com/sirupsen/logrus"
	"sync"
)

/*
*

	@Author   : bob
	@Datetime : 2023-05-22 下午 10:15
	@File     : logger.go
	@Desc     :
*/
var (
	log     *logrus.Logger
	initLog sync.Once
)

func init() {
	initLog.Do(func() {
		log = logrus.New()
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:           "2006-01-02 15:04:05",
			ForceColors:               true,
			EnvironmentOverrideColors: true,
			FullTimestamp:             true,
			DisableLevelTruncation:    true,
		})
	})
}
