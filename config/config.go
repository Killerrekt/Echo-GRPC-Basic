package config

import "github.com/spf13/viper"

type Config struct {
	Port             string `mapstructure:"port"`
	PostgresPort     string `mapstructure:"postgres_port"`
	PostgresHost     string `mapstructure:"postgres_host"`
	PostgresUser     string `mapstructure:"postgres_user"`
	PostgresDB       string `mapstructure:"postgres_db"`
	PostgresPassword string `mapstructure:"postgres_password"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
