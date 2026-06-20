package config

type Config struct {
	App      AppConfig `mapstructure:"app" validate:"required"`
}

type AppConfig struct{
	Port string `mapstructure:"port" validate:"required"`
}