package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"personal-assistant/internal/config"
	"personal-assistant/internal/repositories"
	"personal-assistant/internal/router"

	"github.com/google/logger"
	"github.com/spf13/viper"
)

// readConfigFromFile reads the configuration from a file.
func readConfigFromFile() (*config.Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.AddConfigPath("/etc/personal-assistant/")
	viper.AddConfigPath("$HOME/.sme-demo")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	var readConfig config.Config
	err = viper.Unmarshal(&readConfig)
	if err != nil {
		return nil, err
	}

	return &readConfig, nil
}

func main() {
	logger.Init("Logger", true, false, io.Discard)
	readConfig, err := readConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}

	db := repositories.NewDatabase()
	db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	if readConfig.Mode == config.StartupDevelopment {
		logger.Info("Running in development mode")
		gin.SetMode(gin.DebugMode)
	} else {
		logger.Info("Running in production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	route := router.Router(*readConfig)

	err = route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
