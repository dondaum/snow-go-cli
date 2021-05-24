package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var configData = `
server:
  port: 8080

database:
  dbname: EDW_DB_PROD
  dbuser: "dbuser"
  dbpassword: "dbpassword"

Type: "Snowflake"
`

func getUserHomet() string {
	dirName, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find User Home directory.")
	}
	return dirName
}
func TestGetUserHome(t *testing.T) {
	userHome := getUserHome()
	fmt.Println("User Home path is: ", userHome)
	if userHome == "" {
		t.Errorf("Could not find User home path!")
	}
}

func TestLoadConfig(t *testing.T) {
	configName := "config.yml"
	baseDirName := ".snowgocli"
	dirName := getUserHomet()
	targetFilePath := filepath.Join(dirName, baseDirName, configName)
	targetDirPath := filepath.Join(dirName, baseDirName)

	if _, err := os.Stat(targetDirPath); os.IsNotExist(err) {
		/* Permission Bits */
		os.Mkdir(targetDirPath, 0777)
	}

	f, err := os.Create(targetFilePath)
	if err != nil {
		fmt.Println(err)
	}
	l, err := f.WriteString(configData)
	if err != nil {
		fmt.Println(err)
		f.Close()
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}

	config := LoadConfig()

	expectedDbName := "EDW_DB_PROD"
	if config.Database.DBName != expectedDbName {
		t.Errorf("Found dbname is: %s But expected: %s", config.Database.DBName, expectedDbName)
	}
}
