package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/**
  @Author   : bob
  @Datetime : 2023-05-23 下午 10:22
  @File     : config.go
  @Desc     :
*/

type WebConfig struct {
	WebSocketHost string `yaml:"webSocketHost"`
	WebSocketPort string `yaml:"webSocketPort"`
}

type AppConfig struct {
	WebConfigInfo WebConfig `yaml:"web"`
}

var AppConfigInfo AppConfig

func init() {
	viper.SetConfigName("websocket-demo/config/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("fail to read config file, ", err)
		return
	}
	err = viper.Unmarshal(&AppConfigInfo)
	if err != nil {
		log.Error("fail to get config object, ", err)
	}
	log.Infof("get config info success, %s", AppConfigInfo)

}
