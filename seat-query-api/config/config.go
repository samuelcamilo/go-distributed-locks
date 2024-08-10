package config

import "github.com/spf13/viper"

type (
	Config struct {
		Env      string   `mapstrcture:"env"`
		Port     string   `mapstructure:"port"`
		Database Database `mapstructure:"database"`
	}
	Database struct {
		Writer string `mapstructure:"writer"`
	}
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
