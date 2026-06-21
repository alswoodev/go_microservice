package config

type Config struct {
	App      AppConfig `mapstructure:"app" validate:"required"`
	Database DatabaseConfig `mapstructure:"database" validate:"required"`
	Redis    RedisConfig `mapstructure:"redis" validate:"required"`
	Keys     KeysConfig `mapstructure:"keys" validate:"required"`
}

type AppConfig struct{
	Port string `mapstructure:"port" validate:"required"`
}

type DatabaseConfig struct{
	Host     string `mapstructure:"host" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
}

type RedisConfig struct{
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

type KeysConfig struct{
	JWTKey string `mapstructure:"jwtkey" validate:"required"`
}