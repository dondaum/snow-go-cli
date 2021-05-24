package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Type     string
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName     string
	DBuser     string
	DBPassword string
}

func getUserHome() string {
	dirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find User Home directory.")
	}
	return dirName
}

func LoadConfig() Configuration {
	UserHome := getUserHome()
	baseDirName := ".snowgocli"
	targetPath := filepath.Join(UserHome, baseDirName)
	viper.SetConfigName("config")
	viper.AddConfigPath(targetPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return configuration
}
