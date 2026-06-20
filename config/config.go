package config

import (
	//fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	var cfg Config
	
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./files/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error read config file: %v", err)
	}
	// scope of variable `err` is if statement
	// Unmarshal maps config.yaml to `cfg` variable 
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshal config: %v", err)
	}

	// Check configuration
	// tempDebug, _ := json.Marshal(cfg)
	// fmt.Printf("%s\n", tempDebug)

	return cfg
}