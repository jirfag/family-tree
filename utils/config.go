package utils

import (
	"log"
	"os"
	"strings"

	"github.com/jinzhu/configor"
)

// Config is a struct for backend config
type Config struct {
	APPName string `default:"Gin App"`

	Server struct {
		Host      string `default:"127.0.0.1"`
		Port      string `default:"9012"`
		SecretKey string `default:"SecretKey"`
	}

	Mongo struct {
		Host      string `default:"127.0.0.1"`
		Port      string `default:"27017"`
		SecretKey string `default:"SecretKey"`
		DB        string `default:"db"`
		Username  string `default:"username"`
		Password  string `default:"password"`
	}

	Dayu struct {
		AccessID  string `required:"true"`
		AccessKey string `required:"true"`
	}

	Redis struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"6379"`
		Password string `default:""`
		DB       string `default:"db"`
	}
}

// LoadConfiguration is a function to load cfg from file
func LoadConfiguration() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}

	if os.Getenv("GIN_MODE") == "release" {
		path = strings.Replace(path, "test", "", -1) + "/config.deploy.yml"
	} else {
		path = strings.Replace(path, "test", "", -1) + "/config.yml"
	}

	var config Config
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}

	configor.Load(&config, path)
	if count := len([]rune(config.Server.SecretKey)); count <= 32 {
		for i := 1; i <= 32-count; i++ {
			config.Server.SecretKey += "="
		}
	} else {
		config.Server.SecretKey = string([]byte(config.Server.SecretKey)[:32])
	}
	return config
}

// AppConfig is a struct loaded from config file
var AppConfig = LoadConfiguration()
