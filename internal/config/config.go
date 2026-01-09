package config

import "github.com/spf13/viper"

type Config struct {
    Server struct {
        Port string
    }
    DB struct {
        URL string
    }
}

func Load() *Config {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    cfg := &Config{}
    viper.Unmarshal(cfg)
    return cfg
}
