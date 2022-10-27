package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func Init(env string) {
	fmt.Println("config-init", env)

	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	vp.ReadInConfig()
	// fmt.Println(vp.Get("name"))
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		// log
	}
}

func GetConfig() *viper.Viper {
	return vp
}
