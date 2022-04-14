package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// if you want to use other config file, please export these 3 enviroment variable
var configKind = "LOCAL"
var configPath string
var configFileName string

func init() {
	defer func() {
		recover()
	}()

	viper.AutomaticEnv()

	switch configKind {
	case "LOCAL":
		// get current file path
		_, file, _, _ := runtime.Caller(0)
		currPath := filepath.Dir(file)

		// set config path
		configPath = currPath + "/../../../env"
		configFileName = "config"
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(configPath)

	case "TEST":
		// get current file path
		_, file, _, _ := runtime.Caller(0)
		currPath := filepath.Dir(file)

		// set config path
		configPath = currPath + "/../../../env"
		configFileName = "config_test"
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(configPath)

	case "DEPLOY":
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(configPath)
	}

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.WatchConfig()
}

// Get return key's corresponding value
func Get(key string) interface{} {
	return viper.Get(key)
}

// Set set key-value pair in config file
func Set(key string, value interface{}) {
	viper.Set(key, value)
}
