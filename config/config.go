package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Database Database `yaml:"database"`
	}

	Database struct {
		DBType string `yaml:"dbType"`
		DBUrl  string `yaml:"dbUrl"`
	}
)

func (c *Config) DBType() string {
	return c.Database.DBType
}

func (c *Config) DBUrl() string {
	return c.Database.DBUrl
}

func LoadConfig() (config *Config) {
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("設定ファイルが変更されました:", e.Name)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return
}
