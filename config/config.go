package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var c Config

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv string `mapstructure:"APP_ENV"`
	AppDebug bool `mapstructure:"APP_DEBUG"`
	AppUrl string `mapstructure:"APP_URL"`
	AppPort string `mapstructure:"APP_PORT"`
}


func NewConfig() (*Config, error){
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")

	var err error
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) GetAddress() string{
	return fmt.Sprintf("%s:%s", c.AppUrl, c.AppPort)
}