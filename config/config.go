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
		Secret   Secret   `yaml:"secret"`
	}

	Secret struct {
		Key string `yaml:"key"`
	}

	Database struct {
		Type string `yaml:"type"`
		Url  string `yaml:"url"`
	}
)

func (c *Config) DBType() string {
	return c.Database.Type
}

func (c *Config) DBUrl() string {
	return c.Database.Url
}

func (c *Config) SecretKey() string {
	return c.Secret.Key
}

func LoadConfig() (config *Config) {
	// secret
	secret := viper.New()
	secret.AddConfigPath("config")
	secret.SetConfigName("secret")
	secret.SetConfigType("yaml")

	config = loadConfig(config, secret)

	// config
	nonSecret := viper.New()
	nonSecret.AddConfigPath("config")
	nonSecret.SetConfigType("yaml")

	config = loadConfig(config, nonSecret)

	return
}

func loadConfig(c *Config, v *viper.Viper) *Config {
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("設定ファイルが変更されました:", e.Name)
	})

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return c
}
