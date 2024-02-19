package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&DefaultConf); err != nil {
		return err
	}
	return nil
}
