package common

import (
	"errors"
	"fmt"
	"giskard_api/config"
	"giskard_api/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"strings"
)

func GetClientName(c *gin.Context) (client string, err error) {

	referer, err := url.Parse(c.Request.Referer())
	var clientName string
	if err != nil {
		return "test", err
	}
	//Get the first index of dot
	dotIndex := strings.Index(referer.Host, ".")

	if c.PostForm("clientName") != "" {
		clientName = c.PostForm("clientName")
	} else if c.GetString("clientName") != "" {
		clientName = c.GetString("clientName")
	} else if dotIndex == -1 {
		return "test", errors.New(fmt.Sprintf("Malformed url or unsupported referrer => %s", referer.Host))
	} else {
		tmpString := referer.Host[:dotIndex]
		//Get the last index of - in the intermediate string
		i := strings.LastIndex(tmpString, "-")
		if i > 0 {
			clientName = referer.Host[:i]
		} else {
			//Case : host doesn't have -
			clientName = referer.Host
		}
	}
	config.InitViperConfig()
	clients := config.GetClientsConfig()
	if _, ok := clients[clientName]; ok {
		return clients[clientName], nil
	} else {
		return "test", nil
	}

}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ItemExists(arrayType []string, item string) bool {
	for _, typ := range arrayType {
		if typ == item {
			return true
		}
	}
	return false
}

func GetDataBaseConfig(client string) (parameters models.DatabaseParameters, e error) {
	viper.SetConfigName("clients")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file , %s", err)
	}

	clients := viper.GetStringMap("clients")

	for k, v := range clients {
		if k == client {
			data, _ := v.(map[string]interface{})
			return models.DatabaseParameters{
				DbUser:       data["database-user"].(string),
				DbPassword:   data["database-password"].(string),
				DbHost:       data["database-host"].(string),
				DbPort:       data["database-port"].(int),
				DatabaseName: data["database"].(string),
				DbDriver:     data["database-driver"].(string),
			}, nil

		}
	}
	return parameters, errors.New(fmt.Sprintf("Client name %s is not registered in config files", client))
}
