package config

/**
  @Author   : bob
  @Datetime : 2023-05-23 下午 10:22
  @File     : config.go
  @Desc     :
*/

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config/app")
	viper.AddConfigurationPath(".")
	err := viper.ReadInConfig()
	if err {
		log.Error("fail to read config file, ", err)
		return
	}
	log.Infof("app.webSocketPort is %s", viper.Get("app.webSocketPort"))

}
