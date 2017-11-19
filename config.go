package main

import (
	"github.com/spf13/viper"
)

type workServerConfig struct {
	address                     string
	isDebug                     bool
	tokenKey                    string
	timerWorkDispatchChanCache  int
	timerWorkChanCache          int
	timerWorkCount              int
}

var (
	conf *workServerConfig
)

func readConfiguration() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf = &workServerConfig{
		address:                    viper.GetString("server.address"),
		isDebug:                    viper.GetBool("server.debug"),
		tokenKey:                   viper.GetString("middleware.jwt.tokenkey"),
		timerWorkDispatchChanCache: viper.GetInt("timerwork.dispatch.chancache"),
		timerWorkChanCache:         viper.GetInt("timerwork.work.chancache"),
		timerWorkCount:             viper.GetInt("timerwork.work.count"),
	}
}
