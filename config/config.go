package config

import (
	"fmt"

	_ "github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Address      string `mapstructure:"address"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	EnableCreate bool   `mapstructure:"enableCreate"`
	EnableLog    bool   `mapstructure:"enableLog"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Ip   string `mapstructure:"ip"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func Init(env string) {
	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")

	vp.ReadInConfig()
	err := vp.ReadInConfig()

	if err != nil {
		fmt.Println("configInit", err)
	}

}

func GetConfig() *viper.Viper {
	return vp
}
