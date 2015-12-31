package main

import (
	"log"

	"github.com/spf13/viper"
	fsnotify "gopkg.in/fsnotify.v1"
)

func initCofig() {
	viper.SetDefault("port", ":9007")
	viper.SetDefault("logLevel", "DEBUG")
	viper.SetDefault("watch", false)

	viper.SetConfigName("octopus-conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/.octopus/")
	viper.AddConfigPath("/etc/octopus/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Panicf("Fatal error config file: %s \n", err)
	}

	if viper.GetBool("watch") {
		viper.WatchConfig()
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	log.Printf("loading config %s \n", viper.ConfigFileUsed())
}
