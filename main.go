package main

import (
	"github.com/dondaum/snow-go-cli/cmd"
	"github.com/dondaum/snow-go-cli/config"
	"github.com/dondaum/snow-go-cli/db"
)

func main() {
	db.TestSnowConn()
	config.LoadConfig()
	cmd.Execute()
}
