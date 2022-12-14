package main

import (
	"go-redis/config"
	"go-redis/lib/logger"
	"os"
)

const configFile string = "redis.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	if fileExist(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defaultProperties
	}

}
