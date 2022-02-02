package config

import (
	"github.com/spf13/viper"
	"log"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func InitViperConfig() {

	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file , %s", err)
	}

}

func GetClientsConfig() map[string]string {
	return viper.GetStringMapString("clients.mapping")
}
